const std = @import("std");
const builtin = @import("builtin");

fn readInputFile(allocator: std.mem.Allocator, filename: []const u8) ![]u32 {
    var result = std.ArrayList(u32).init(allocator);

    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    const reader = file.reader();

    while (true) {
        var line: []u8 = undefined;
        if (builtin.os.tag == .windows) {
            // NOTE: Read another byte on windows due to two-byte eol.
            line = reader.readUntilDelimiterAlloc(allocator, '\r', 512) catch break;
            _ = try reader.readByte();
        } else {
            line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        }
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
                // std.log.debug("low point->{}", .{depth});
                result += depth + 1;
            }
        }
    }

    return result;
}

fn drawGrid(width: u32, height: u32, grid: []u32) void {
    var y: u32 = 0;
    std.log.debug("grid", .{});
    while (y < height) : (y += 1) {
        std.log.debug("{any}", .{grid[(y * width)..((y + 1) * width)]});
    }
}

fn flood(x: u32, y: u32, width: u32, height: u32, grid: []u32) u32 {
    var result: u32 = 1;
    const j = (y * width) + x;
    grid[j] = 9;
    if (y > 0 and grid[j - width] != 9) {
        result += flood(x, y - 1, width, height, grid);
    }
    if (y < height - 1 and grid[j + width] != 9) {
        result += flood(x, y + 1, width, height, grid);
    }
    if (x > 0 and grid[j - 1] != 9) {
        result += flood(x - 1, y, width, height, grid);
    }
    if (x < width - 1 and grid[j + 1] != 9) {
        result += flood(x + 1, y, width, height, grid);
    }
    return result;
}

const desc_u32 = std.sort.desc(u32);

pub fn task2(allocator: std.mem.Allocator, width: u32, height: u32, input: []u32) !u32 {
    var result: u32 = 1;

    var basins = std.ArrayList(u32).init(allocator);
    defer basins.clearAndFree();

    var y: u32 = 0;
    while (y < height) : (y += 1) {
        var i = y * width;
        var x: u32 = 0;
        while (x < width) : (x += 1) {
            if (input[i + x] != 9) {
                const basin = flood(x, y, width, height, input);
                // drawGrid(width, height, input);
                try basins.append(basin);
            }
        }
    }

    std.sort.sort(u32, basins.items, {}, desc_u32);
    for (basins.items[0..3]) |i| {
        result *= i;
    }

    return result;
}

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(100, 100, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(allocator, 100, 100, input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
