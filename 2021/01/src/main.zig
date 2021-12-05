const std = @import("std");

fn readInputFile(allocator: *std.mem.Allocator, filename: []const u8) anyerror![]u32 {
    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    var result = std.ArrayList(u32).init(allocator);

    var index: u32 = 0;
    var buffer: [8]u8 = undefined;
    const reader = file.reader();
    while (true) {
        var byte = reader.readByte() catch break;
        buffer[index] = byte;
        if (byte == '\n') {
            const value = try std.fmt.parseUnsigned(u32, buffer[0..index], 10);
            try result.append(value);
            index = 0;
        } else {
            index += 1;
        }
    }

    return result.items;
}

pub fn task1(depth_values: []const u32) u32 {
    var count: u32 = 0;
    var prev = depth_values[0];
    for (depth_values) |depth| {
        if (depth > prev) {
            count += 1;
        }
        prev = depth;
    }
    return count;
}

pub fn main() anyerror!void {
    var buffer: [65536]u8 = undefined;
    const allocator = &std.heap.FixedBufferAllocator.init(&buffer).allocator;

    const input = try readInputFile(allocator, "input.txt");
    std.log.debug("input {any}", .{input});

    const task_1_result = task1(input);
    std.log.info("Task 1 result: {}", .{task_1_result});
}
