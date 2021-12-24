const std = @import("std");

fn readInputFile(allocator: std.mem.Allocator, filename: []const u8) ![]u32 {
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

pub fn drawGrid(width: u32, height: u32, grid: []u32) void {
    var y: u32 = 0;
    while (y < height) : (y += 1) {
        std.log.debug("{any}", .{grid[(y * width)..((y + 1) * width)]});
    }
}

fn flash(x: u32, y: u32, width: u32, height: u32, squid: []u32) void {
    var yy: i32 = -1;
    while (yy < 2) : (yy += 1) {
        if (yy + @intCast(i32, y) < 0 or yy + @intCast(i32, y) >= @intCast(i32, height)) {
            continue;
        }

        var i = @intCast(u32, yy + @intCast(i32, y)) * width;
        var xx: i32 = -1;
        while (xx < 2) : (xx += 1) {
            if (xx + @intCast(i32, x) < 0 or xx + @intCast(i32, x) >= @intCast(i32, width)) {
                continue;
            }

            if (!(xx == 0 and yy == 0)) {
                squid[i + @intCast(u32, xx + @intCast(i32, x))] += 1;
            }
        }
    }
}

pub fn iterateEnergy(width: u32, height: u32, squid: []u32) u32 {
    var result: u32 = 0;

    var y: u32 = 0;
    while (y < height) : (y += 1) {
        var i = y * width;
        var x: u32 = 0;
        while (x < width) : (x += 1) {
            squid[i + x] += 1;
        }
    }

    while (true) {
        var did_flash: bool = false;
        y = 0;
        while (y < height) : (y += 1) {
            var i = y * width;
            var x: u32 = 0;
            while (x < width) : (x += 1) {
                if (squid[i + x] > 9 and squid[i + x] < 99) {
                    squid[i + x] = 99;
                    flash(x, y, width, height, squid);
                    did_flash = true;
                }
            }
        }
        if (!did_flash) {
            break;
        }
    }

    y = 0;
    while (y < height) : (y += 1) {
        var i = y * width;
        var x: u32 = 0;
        while (x < width) : (x += 1) {
            if (squid[i + x] > 9) {
                squid[i + x] = 0;
                result += 1;
            }
        }
    }

    return result;
}

pub fn task1(width: u32, height: u32, input: []u32) u32 {
    var result: u32 = 0;

    //std.log.debug("\ninitial", .{});
    //drawGrid(width, height, input);

    var i: u32 = 0;
    while (i < 100) : (i += 1) {
        result += iterateEnergy(width, height, input);
        //std.log.debug("\nafter {} step", .{i + 1});
        //drawGrid(width, height, input);
    }

    return result;
}

pub fn task2(width: u32, height: u32, input: []u32) u32 {
    var result: u32 = 1;

    while (true) : (result += 1) {
        if (iterateEnergy(width, height, input) == width * height) {
            //std.log.debug("\nafter {} steps", .{result});
            //drawGrid(width, height, input);
            break;
        }
    }

    return result;
}

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const input = try readInputFile(allocator, "input.txt");

    var input_copy = std.ArrayList(u32).init(allocator);
    try input_copy.appendSlice(input);
    const task_1_result = task1(10, 10, input_copy.items);
    std.log.info("Task 1 result: {}", .{task_1_result});
    input_copy.clearAndFree();

    try input_copy.appendSlice(input);
    const task_2_result = task2(10, 10, input_copy.items);
    std.log.info("Task 2 result: {}", .{task_2_result});
    input_copy.clearAndFree();
}
