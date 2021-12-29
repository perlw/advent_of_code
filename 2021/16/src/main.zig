const std = @import("std");
const builtin = @import("builtin");

pub const DecodingError = error{
    InvalidInput,
    UnexpectedPacketType,
};

pub inline fn charToBinary(char: u8) !u4 {
    if (char >= '0' and char <= '9') {
        return @intCast(u4, char - '0');
    } else if (char >= 'A' and char <= 'F') {
        return @intCast(u4, 10 + (char - 'A'));
    }
    return DecodingError.InvalidInput;
}

pub fn hexToBinaryArrayList(allocator: std.mem.Allocator, hex: []const u8) !std.ArrayList(u1) {
    var result = std.ArrayList(u1).init(allocator);

    for (hex) |char| {
        const bin = try charToBinary(char);
        try result.appendSlice(&.{
            @intCast(u1, (bin >> 3) & 0x1),
            @intCast(u1, (bin >> 2) & 0x1),
            @intCast(u1, (bin >> 1) & 0x1),
            @intCast(u1, (bin >> 0) & 0x1),
        });
    }

    return result;
}

fn printRawBits(value: u16) void {
    std.debug.print("\n⠿", .{});
    var i: u32 = 0;
    while (i < 16) : (i += 1) {
        std.debug.print("{}", .{value >> (@intCast(u4, 15 - i)) & 0x1});
    }
    std.debug.print("\n", .{});
}

fn peekPacketType(packets: *std.ArrayList(u1)) u3 {
    const peek_packet_type: []u1 = packets.items[3..6];
    return (@as(u3, peek_packet_type[0]) << 2) + (@as(u3, peek_packet_type[1]) << 1) + @as(u3, peek_packet_type[2]);
}

fn readHeader(packets: *std.ArrayList(u1), output_version: *u32, output_packet_type: *u3) u15 {
    output_version.* = (@as(u3, packets.orderedRemove(0)) << 2) + (@as(u3, packets.orderedRemove(0)) << 1) + @as(u3, packets.orderedRemove(0));
    output_packet_type.* = (@as(u3, packets.orderedRemove(0)) << 2) + (@as(u3, packets.orderedRemove(0)) << 1) + @as(u3, packets.orderedRemove(0));
    return 6;
}

pub fn readLiteralValue(packets: *std.ArrayList(u1), output_version: *u32, output_value: *u16) !u15 {
    if (packets.items.len < 7) {
        return 0;
    }

    var version: u32 = 0;
    var packet_type: u3 = 0;
    var bits_read: u15 = readHeader(packets, &version, &packet_type);
    output_version.* += version;

    if (packet_type != 4) {
        return DecodingError.UnexpectedPacketType;
    }

    var should_continue = true;
    while (should_continue) {
        should_continue = (@as(u1, packets.orderedRemove(0)) == 1);

        output_value.* <<= 4;
        output_value.* +=
            ((@as(u16, packets.orderedRemove(0)) << 3) +
            (@as(u16, packets.orderedRemove(0)) << 2) +
            (@as(u16, packets.orderedRemove(0)) << 1) +
            @as(u16, packets.orderedRemove(0)));

        bits_read += 5;
        // printRawBits(output_value.*);
    }

    return bits_read;
}

pub fn readOperatorPacket(packets: *std.ArrayList(u1), output_version: *u32, output_value: *u16) !u15 {
    if (packets.items.len < 7) {
        return 0;
    }

    _ = output_value;
    var version: u32 = 0;
    var packet_type: u3 = 0;
    var bits_read: u15 = readHeader(packets, &version, &packet_type);
    output_version.* += version;

    if (packet_type == 4) {
        return DecodingError.UnexpectedPacketType;
    }

    const length_type = @as(u1, packets.orderedRemove(0));
    if (length_type == 0) {
        var length: u15 = 0;
        var i: u4 = 0;
        while (i < 15) : (i += 1) {
            length += @as(u15, packets.orderedRemove(0)) << (14 - i);
        }
        bits_read += 15;

        var read_bits: u15 = 0;
        while (read_bits < length and packets.items.len > 7) {
            var literal_version: u32 = 0;
            var val: u16 = 0;
            const peeked_packet_type = peekPacketType(packets);
            if (peeked_packet_type == 4) {
                read_bits += readLiteralValue(packets, &literal_version, &val) catch unreachable;
            } else {
                read_bits += readOperatorPacket(packets, &literal_version, &val) catch unreachable;
            }
            output_version.* += literal_version;
        }
        bits_read += read_bits;
    } else {
        var num_packets: u11 = 0;
        var i: u4 = 0;
        while (i < 11) : (i += 1) {
            num_packets += @as(u11, packets.orderedRemove(0)) << (10 - i);
        }
        bits_read += 11;

        var read_packets: u11 = 0;
        while (read_packets < num_packets) : (read_packets += 1) {
            var literal_version: u32 = 0;
            var val: u16 = 0;
            const peeked_packet_type = peekPacketType(packets);
            if (peeked_packet_type == 4) {
                bits_read += readLiteralValue(packets, &literal_version, &val) catch unreachable;
            } else {
                bits_read += readOperatorPacket(packets, &literal_version, &val) catch unreachable;
            }
            output_version.* += literal_version;
        }
    }

    return bits_read;
}

pub fn runPacketDecode(packets: *std.ArrayList(u1), version: *u32, output_value: *u16) !void {
    version.* = 0;
    output_value.* = 0;

    const packet_type = peekPacketType(packets);
    if (packet_type == 4) {
        _ = try readLiteralValue(packets, version, output_value);
    } else {
        _ = try readOperatorPacket(packets, version, output_value);
    }
}

pub fn readInput(allocator: std.mem.Allocator, reader: anytype) !std.ArrayList(u1) {
    var line: []u8 = undefined;
    if (builtin.os.tag == .windows and !builtin.is_test) {
        // NOTE: Read another byte on windows due to two-byte eol.
        // NOTE: Check if in testing since tests only add single-byte eol in multiline strings.
        line = reader.readUntilDelimiterAlloc(allocator, '\r', 2048) catch unreachable;
        _ = try reader.readByte();
    } else {
        line = reader.readUntilDelimiterAlloc(allocator, '\n', 2048) catch unreachable;
    }
    defer allocator.free(line);

    return hexToBinaryArrayList(allocator, line);
}

pub fn task1(input: *std.ArrayList(u1)) !u32 {
    var version: u32 = 0;
    var result: u16 = 0;
    try runPacketDecode(input, &version, &result);
    return version;
}

pub fn task2(input: *std.ArrayList(u1)) !u32 {
    _ = input;
    return 0;
}

pub fn main() !void {
    var buffer: [8000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const file = try std.fs.cwd().openFile("input.txt", .{ .read = true });
    defer file.close();

    var input = try readInput(allocator, file.reader());
    defer input.deinit();

    const task_1_result = try task1(&input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    // const task_2_result = try task2(&input);
    // std.log.info("Task 2 result: {}", .{task_2_result});
}
