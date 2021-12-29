const std = @import("std");
const builtin = @import("builtin");

pub const DecodingError = error{
    InvalidInput,
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

pub fn readInput(allocator: std.mem.Allocator, reader: anytype) !u32 {
    var risk_levels = std.ArrayList(u32).init(allocator);
    defer risk_levels.deinit();

    var width: u32 = 0;
    var height: u32 = 0;
    while (true) {
        var line: []u8 = undefined;
        if (builtin.os.tag == .windows and !builtin.is_test) {
            // NOTE: Read another byte on windows due to two-byte eol.
            // NOTE: Check if in testing since tests only add single-byte eol in multiline strings.
            line = reader.readUntilDelimiterAlloc(allocator, '\r', 512) catch break;
            _ = try reader.readByte();
        } else {
            line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        }
        defer allocator.free(line);

        if (line.len > width) {
            width = @intCast(u32, line.len);
        }
        height += 1;

        for (line) |char| {
            const risk: u32 = char - '0';
            try risk_levels.append(risk);
        }
    }

    return 0;
}

pub fn task1(input: u32) !u32 {
    return input;
}

pub fn task2(input: u32) !u32 {
    return input;
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

    const task_2_result = try task2(&input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
