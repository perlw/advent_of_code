const std = @import("std");
const builtin = @import("builtin");

pub const Point = struct {
    x: i32,
    y: i32,
};

pub const TargetArea = struct {
    top_left: Point,
    bottom_right: Point,
};

pub fn readInput(allocator: std.mem.Allocator, reader: anytype) !TargetArea {
    var result: TargetArea = undefined;

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

    var it = std.mem.split(u8, line[13..], ", ");
    if (it.next()) |part| {
        var it2 = std.mem.split(u8, part[2..], "..");
        if (it2.next()) |sub_part| {
            result.top_left.x = try std.fmt.parseInt(i32, sub_part, 10);
        }
        if (it2.next()) |sub_part| {
            result.bottom_right.x = try std.fmt.parseInt(i32, sub_part, 10);
        }
    }
    if (it.next()) |part| {
        var it2 = std.mem.split(u8, part[2..], "..");
        if (it2.next()) |sub_part| {
            result.bottom_right.y = try std.fmt.parseInt(i32, sub_part, 10);
        }
        if (it2.next()) |sub_part| {
            result.top_left.y = try std.fmt.parseInt(i32, sub_part, 10);
        }
    }

    return result;
}

fn fireProbe(x: i32, y: i32, top: *i32, input: TargetArea) bool {
    var pos = Point{
        .x = 0,
        .y = 0,
    };
    var initial_vel = Point{
        .x = x,
        .y = y,
    };
    var vel = initial_vel;

    while (pos.x < input.bottom_right.x and pos.y > input.bottom_right.y) {
        pos.x += vel.x;
        pos.y += vel.y;
        if (pos.y > top.*) {
            top.* = pos.y;
        }

        if (pos.x >= input.top_left.x and pos.x <= input.bottom_right.x and pos.y >= input.bottom_right.y and pos.y <= input.top_left.y) {
            return true;
        }

        vel.x += @as(i32, if (vel.x > 0) -1 else @as(i32, if (vel.x < 0) 1 else 0));
        vel.y -= 1;
    }

    return false;
}

pub fn task1(input: TargetArea) i32 {
    var result: i32 = 0;

    var y: i32 = 0;
    while (y < 1000) : (y += 1) {
        var x: i32 = 1;
        while (x < 1000) : (x += 1) {
            var max_y: i32 = 0;
            if (fireProbe(x, y, &max_y, input)) {
                if (max_y > result) {
                    result = max_y;
                }
            }
        }
    }

    return result;
}

pub fn task2(input: TargetArea) u32 {
    var result: u32 = 0;

    var y: i32 = -1000;
    while (y < 1000) : (y += 1) {
        var x: i32 = 1;
        while (x < 1000) : (x += 1) {
            var max_y: i32 = 0;
            if (fireProbe(x, y, &max_y, input)) {
                result += 1;
            }
        }
    }

    return result;
}

pub fn main() !void {
    var buffer: [8000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const file = try std.fs.cwd().openFile("input.txt", .{ .read = true });
    defer file.close();

    var input = try readInput(allocator, file.reader());

    const task_1_result = task1(input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
