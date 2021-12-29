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

fn peekPacketType(packets: *std.ArrayList(u1)) PacketType {
    const peek_packet_type: []u1 = packets.items[3..6];
    return @intToEnum(PacketType, (@as(u3, peek_packet_type[0]) << 2) + (@as(u3, peek_packet_type[1]) << 1) + @as(u3, peek_packet_type[2]));
}

fn readHeader(packets: *std.ArrayList(u1), output_version: *u32, output_packet_type: *PacketType) u15 {
    output_version.* = (@as(u3, packets.orderedRemove(0)) << 2) + (@as(u3, packets.orderedRemove(0)) << 1) + @as(u3, packets.orderedRemove(0));
    output_packet_type.* = @intToEnum(PacketType, (@as(u3, packets.orderedRemove(0)) << 2) + (@as(u3, packets.orderedRemove(0)) << 1) + @as(u3, packets.orderedRemove(0)));
    return 6;
}

const PacketType = enum(u3) {
    Sum = 0,
    Product = 1,
    Minimum = 2,
    Maximum = 3,
    Literal = 4,
    GreaterThan = 5,
    LessThan = 6,
    EqualTo = 7,
};

pub fn readLiteralValue(packets: *std.ArrayList(u1), output_version: *u32, output_value: *u64) !u15 {
    var version: u32 = 0;
    var packet_type: PacketType = undefined;
    var bits_read: u15 = readHeader(packets, &version, &packet_type);
    output_version.* += version;

    if (packet_type != .Literal) {
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

pub fn readOperatorPacket(allocator: std.mem.Allocator, packets: *std.ArrayList(u1), output_version: *u32, output_value: *u64) !u15 {
    var version: u32 = 0;
    var packet_type: PacketType = undefined;
    var bits_read: u15 = readHeader(packets, &version, &packet_type);
    output_version.* += version;

    if (packet_type == .Literal) {
        return DecodingError.UnexpectedPacketType;
    }

    var values = std.ArrayList(u64).init(allocator);
    defer values.deinit();

    const length_type = @as(u1, packets.orderedRemove(0));
    if (length_type == 0) {
        var length: u15 = 0;
        var i: u4 = 0;
        while (i < 15) : (i += 1) {
            length += @as(u15, packets.orderedRemove(0)) << (14 - i);
        }
        bits_read += 15;

        var read_bits: u15 = 0;
        while (read_bits < length and length - read_bits > 6 and packets.items.len > 6) {
            var packet_version: u32 = 0;
            var packet_value: u64 = 0;
            const peeked_packet_type = peekPacketType(packets);
            if (peeked_packet_type == .Literal) {
                read_bits += readLiteralValue(packets, &packet_version, &packet_value) catch unreachable;
            } else {
                read_bits += readOperatorPacket(allocator, packets, &packet_version, &packet_value) catch unreachable;
            }
            output_version.* += packet_version;
            try values.append(packet_value);
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
            var packet_version: u32 = 0;
            var packet_value: u64 = 0;
            const peeked_packet_type = peekPacketType(packets);
            if (peeked_packet_type == .Literal) {
                bits_read += readLiteralValue(packets, &packet_version, &packet_value) catch unreachable;
            } else {
                bits_read += readOperatorPacket(allocator, packets, &packet_version, &packet_value) catch unreachable;
            }
            output_version.* += packet_version;
            try values.append(packet_value);
        }
    }

    switch (packet_type) {
        .Sum => {
            for (values.items) |item| {
                output_value.* += item;
            }
        },
        .Product => {
            var product: u64 = 1;
            for (values.items) |item| {
                product *= item;
            }
            output_value.* = product;
        },
        .Minimum => {
            var min: u64 = std.math.inf_u64;
            for (values.items) |item| {
                if (item < min) {
                    min = item;
                }
            }
            output_value.* = min;
        },
        .Maximum => {
            var max: u64 = 0;
            for (values.items) |item| {
                if (item > max) {
                    max = item;
                }
            }
            output_value.* = max;
        },
        .Literal => unreachable,
        .GreaterThan => {
            if (values.items.len != 2) unreachable;

            output_value.* = @boolToInt(values.items[0] > values.items[1]);
        },
        .LessThan => {
            if (values.items.len != 2) unreachable;

            output_value.* = @boolToInt(values.items[0] < values.items[1]);
        },
        .EqualTo => {
            if (values.items.len != 2) unreachable;

            output_value.* = @boolToInt(values.items[0] == values.items[1]);
        },
    }

    return bits_read;
}

pub fn runPacketDecode(allocator: std.mem.Allocator, packets: *std.ArrayList(u1), version: *u32, output_value: *u64) !void {
    version.* = 0;
    output_value.* = 0;

    const packet_type = peekPacketType(packets);
    if (packet_type == .Literal) {
        _ = try readLiteralValue(packets, version, output_value);
    } else {
        _ = try readOperatorPacket(allocator, packets, version, output_value);
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

pub fn main() !void {
    var buffer: [8000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const file = try std.fs.cwd().openFile("input.txt", .{ .read = true });
    defer file.close();

    var input = try readInput(allocator, file.reader());
    defer input.deinit();

    var version: u32 = 0;
    var result: u64 = 0;
    try runPacketDecode(allocator, &input, &version, &result);
    std.log.info("Task 1 result: {}", .{version});
    std.log.info("Task 2 result: {}", .{result});
}
