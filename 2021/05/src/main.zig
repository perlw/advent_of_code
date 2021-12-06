const std = @import("std");

pub const Point = struct {
    x: u32 = 0,
    y: u32 = 0,
};

pub const Vent = struct {
    a: Point = .{},
    b: Point = .{},
};

fn readInputFile(allocator: *std.mem.Allocator, filename: []const u8) anyerror![]Vent {
    var result = std.ArrayList(Vent).init(allocator);

    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    const reader = file.reader();
    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 32) catch break;
        defer allocator.free(line);

        var points = std.mem.tokenize(line, " -> ");
        var vent: Vent = .{};
        if (points.next()) |slice| {
            var coords = std.mem.tokenize(slice, ",");
            vent.a.x = try std.fmt.parseUnsigned(u32, coords.next().?, 10);
            vent.a.y = try std.fmt.parseUnsigned(u32, coords.next().?, 10);
        }
        if (points.next()) |slice| {
            var coords = std.mem.tokenize(slice, ",");
            vent.b.x = try std.fmt.parseUnsigned(u32, coords.next().?, 10);
            vent.b.y = try std.fmt.parseUnsigned(u32, coords.next().?, 10);
        }

        try result.append(vent);
    }

    return result.items;
}

fn drawGrid(max: Point, grid: []u32) void {
    std.log.debug("===", .{});
    var y: u32 = 0;
    while (y < max.y) : (y += 1) {
        const start: u32 = (y * max.x);
        const stop: u32 = (y * max.x) + max.x;
        std.log.debug("{any}", .{grid[start..stop]});
    }
}

fn putDownVents(diagonals: bool, max: Point, vents: []const Vent, grid: []u32) void {
    for (vents) |vent| {
        if (vent.a.x == vent.b.x) {
            const a = if (vent.a.y < vent.b.y) vent.a else vent.b;
            const b = if (vent.a.y < vent.b.y) vent.b else vent.a;

            var x = a.x;
            var y = a.y;
            while (y <= b.y) : (y += 1) {
                grid[(y * max.x) + x] += 1;
            }
        } else if (vent.a.y == vent.b.y) {
            const a = if (vent.a.x < vent.b.x) vent.a else vent.b;
            const b = if (vent.a.x < vent.b.x) vent.b else vent.a;

            var x = a.x;
            var y = a.y;
            while (x <= b.x) : (x += 1) {
                grid[(y * max.x) + x] += 1;
            }
        } else if (diagonals) {
            const dir_x = @as(i32, if (vent.a.x < vent.b.x) 1 else -1);
            const dir_y = @as(i32, if (vent.a.y < vent.b.y) 1 else -1);

            var x = vent.a.x;
            var y = vent.a.y;
            while (x != vent.b.x and y != vent.b.y) {
                grid[(y * max.x) + x] += 1;

                x = if (dir_x < 0) x - 1 else x + 1;
                y = if (dir_y < 0) y - 1 else y + 1;
            }
            grid[(vent.b.y * max.x) + vent.b.x] += 1; // NOTE: Make sure to add final point.
        }
    }
}

pub fn task1(allocator: *std.mem.Allocator, max: Point, vents: []const Vent) !u32 {
    var result: u32 = 0;

    var grid = try allocator.alloc(u32, max.x * max.y);
    defer allocator.free(grid);
    std.mem.set(u32, grid, 0);

    putDownVents(false, max, vents, grid);

    // drawGrid(max, grid);
    for (grid) |cell| {
        if (cell > 1) {
            result += 1;
        }
    }

    return result;
}

pub fn task2(allocator: *std.mem.Allocator, max: Point, vents: []Vent) !u32 {
    var result: u32 = 0;

    var grid = try allocator.alloc(u32, max.x * max.y);
    defer allocator.free(grid);
    std.mem.set(u32, grid, 0);

    putDownVents(true, max, vents, grid);

    // drawGrid(max, grid);
    for (grid) |cell| {
        if (cell > 1) {
            result += 1;
        }
    }

    return result;
}

pub fn main() anyerror!void {
    var buffer: [4000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    var allocator = &fixed_buffer.allocator;

    const input = try readInputFile(allocator, "input.txt");
    var max: Point = .{ .x = 0, .y = 0 };
    for (input) |vent| {
        max.x = if (vent.a.x > max.x) vent.a.x else max.x;
        max.y = if (vent.a.y > max.y) vent.a.y else max.y;
        max.x = if (vent.b.x > max.x) vent.b.x else max.x;
        max.y = if (vent.b.y > max.y) vent.b.y else max.y;
    }
    max.x += 1;
    max.y += 1;

    const task_1_result = try task1(allocator, max, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = try task2(allocator, max, input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
