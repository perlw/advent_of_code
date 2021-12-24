const std = @import("std");

fn readInputFile(allocator: std.mem.Allocator, filename: []const u8) anyerror![]u32 {
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
    var result: u32 = 0;
    var prev = depth_values[0];
    for (depth_values) |depth| {
        if (depth > prev) {
            result += 1;
        }
        prev = depth;
    }
    return result;
}

pub fn task2(depth_values: []const u32) u32 {
    var result: u32 = 0;
    var prev_batch: u32 = 0;
    var batch_index: u32 = 0;
    var batch_value: u32 = 0;

    var count: u32 = 0;
    while (batch_index < depth_values.len) {
        batch_value += depth_values[batch_index];
        batch_index += 1;
        count += 1;

        if (count == 3) {
            if (prev_batch > 0 and batch_value > prev_batch) {
                result += 1;
            }

            batch_index -= 2;
            prev_batch = batch_value;
            count = 0;
            batch_value = 0;
        }
    }

    return result;
}

pub fn main() anyerror!void {
    var buffer: [65536]u8 = undefined;
    const allocator = std.heap.FixedBufferAllocator.init(&buffer).allocator();

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
