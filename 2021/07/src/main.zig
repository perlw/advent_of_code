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
        line = reader.readUntilDelimiterAlloc(allocator, '\r', 4096) catch unreachable;
        _ = try reader.readByte();
    } else {
        line = reader.readUntilDelimiterAlloc(allocator, '\n', 4096) catch unreachable;
    }
    defer allocator.free(line);

    var it = std.mem.split(u8, line, ",");
    while (it.next()) |slice| {
        const value = try std.fmt.parseUnsigned(u32, slice, 10);
        try result.append(value);
    }

    return result.items;
}

const asc_u32 = std.sort.asc(u32);

pub fn task1(crabs: []u32) u32 {
    var result: u32 = 999999;

    const min = std.sort.min(u32, crabs, {}, asc_u32).?;
    const max = std.sort.max(u32, crabs, {}, asc_u32).?;

    var i = min;
    while (i <= max) : (i += 1) {
        var fuel_required: u32 = 0;

        for (crabs) |crab_pos| {
            fuel_required += if (i <= crab_pos) crab_pos - i else i - crab_pos;
        }

        std.log.debug("{}: {} required", .{ i, fuel_required });
        if (fuel_required < result) {
            result = fuel_required;
        }
    }

    return result;
}

fn calculateFuel(target_position: u32, crabs: []u32) u32 {
    var result: u32 = 0;

    for (crabs) |crab_pos| {
        const distance = if (target_position <= crab_pos) crab_pos - target_position else target_position - crab_pos;
        var cost: u32 = 0;
        var i: u32 = 0;
        while (i <= distance) : (i += 1) {
            cost += i;
        }
        result += cost;
    }

    return result;
}

pub fn task2(crabs: []u32) u64 {
    var result: u64 = 9999999999;

    const min = std.sort.min(u32, crabs, {}, asc_u32).?;
    const max = std.sort.max(u32, crabs, {}, asc_u32).?;

    var i = min;
    while (i <= max) : (i += 1) {
        const fuel_required = calculateFuel(i, crabs);
        std.log.debug("{}: {} required", .{ i, fuel_required });
        if (fuel_required < result) {
            result = fuel_required;
        }
    }

    return result;
}

pub fn main() anyerror!void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
