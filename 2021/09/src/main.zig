const std = @import("std");

fn readInputFile(allocator: *std.mem.Allocator, filename: []const u8) ![]u32 {
    var result = std.ArrayList(u32).init(allocator);

    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    const reader = file.reader();

    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        defer allocator.free(line);

        for (line) |char| {
            try result.append(@as(u32, char - 48));
        }
    }

    return result.items;
}

pub fn task1(width: u32, height: u32, input: []u32) u32 {
    var result: u32 = 0;

    var y: u32 = 0;
    while (y < height) : (y += 1) {
        var i = y * width;
        var x: u32 = 0;
        while (x < width) : (x += 1) {
            var j = i + x;
            var depth = input[j];

            var is_low_point: bool = true;
            if (y > 0 and input[j - width] <= depth) {
                is_low_point = false;
            }
            if (y < height - 1 and input[j + width] <= depth) {
                is_low_point = false;
            }
            if (x > 0 and input[j - 1] <= depth) {
                is_low_point = false;
            }
            if (x < width - 1 and input[j + 1] <= depth) {
                is_low_point = false;
            }

            if (is_low_point) {
                std.log.debug("low point->{}", .{depth});
                result += depth + 1;
            }
        }
    }

    return result;
}

pub fn task2(input: []u32) u32 {
    var result: u32 = 0;

    return result;
}

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    var allocator = &fixed_buffer.allocator;

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(100, 100, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(input);
    std.log.info("Task 1 result: {}", .{task_2_result});
}
