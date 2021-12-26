const std = @import("std");
const builtin = @import("builtin");

fn readInputFile(allocator: std.mem.Allocator, filename: []const u8) anyerror![]u32 {
    var result = std.ArrayList(u32).init(allocator);

    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    const reader = file.reader();
    var line: []u8 = undefined;
    if (builtin.os.tag == .windows) {
        // NOTE: Read another byte on windows due to two-byte eol.
        line = reader.readUntilDelimiterAlloc(allocator, '\r', 1024) catch unreachable;
        _ = try reader.readByte();
    } else {
        line = reader.readUntilDelimiterAlloc(allocator, '\n', 1024) catch unreachable;
    }
    defer allocator.free(line);

    var it = std.mem.split(u8, line, ",");
    while (it.next()) |slice| {
        const value = try std.fmt.parseUnsigned(u32, slice, 10);
        try result.append(value);
    }

    return result.items;
}

pub fn task1(allocator: std.mem.Allocator, iterations: u32, initial_fish: []u32) !u32 {
    var result: u32 = 0;

    var lantern_fish = std.ArrayList(u32).init(allocator);
    defer lantern_fish.deinit();
    try lantern_fish.appendSlice(initial_fish);

    std.log.info("initial: {any}", .{lantern_fish.items});

    var i: u32 = 0;
    while (i < iterations) : (i += 1) {
        var new_fish: u32 = 0;
        for (lantern_fish.items) |*fish| {
            if (fish.* == 0) {
                fish.* = 6;
                new_fish += 1;
            } else {
                fish.* -= 1;
            }
        }
        try lantern_fish.appendNTimes(8, new_fish);

        std.log.info("{}: {}", .{ i, lantern_fish.items.len });
    }
    result = @intCast(u32, lantern_fish.items.len);

    return result;
}

pub fn task2(iterations: u32, initial_fish: []u32) u64 {
    var result: u64 = 0;

    var buckets: [9]u64 = [_]u64{0} ** 9;
    for (initial_fish) |fish| {
        buckets[fish] += 1;
    }

    std.log.info("buckets: {any}", .{buckets});

    var i: u32 = 0;
    while (i < iterations) : (i += 1) {
        const spawn = buckets[0];
        std.mem.rotate(u64, &buckets, 1);
        buckets[6] += spawn;

        var num_fish: u64 = 0;
        for (buckets) |bucket| {
            num_fish += bucket;
        }
        std.log.info("{}: {}", .{ i, num_fish });
    }

    for (buckets) |bucket| {
        result += bucket;
    }

    return result;
}

pub fn main() anyerror!void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = try task1(allocator, 80, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(256, input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
