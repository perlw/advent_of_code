const std = @import("std");

const Point = struct {
    x: u32,
    y: u32,
};

const Input = struct {
    points: []Point,
    folds: []Point,
};

pub fn readInput(allocator: std.mem.Allocator, reader: anytype) !Input {
    var points = std.ArrayList(Point).init(allocator);
    var folds = std.ArrayList(Point).init(allocator);

    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        defer allocator.free(line);

        if (line.len == 0) {
            break;
        }

        var it = std.mem.split(u8, line, ",");
        var x = try std.fmt.parseInt(u32, it.next().?, 10);
        var y = try std.fmt.parseInt(u32, it.next().?, 10);
        try points.append(Point{
            .x = x,
            .y = y,
        });
    }

    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        defer allocator.free(line);

        var it = std.mem.split(u8, line, "=");
        var p1 = it.next().?;
        var dir = p1[p1.len - 1];
        var value = try std.fmt.parseInt(u32, it.next().?, 10);
        var fold = Point{
            .x = 0,
            .y = 0,
        };
        if (dir == 'x') {
            fold.x = value;
        } else {
            fold.y = value;
        }
        try folds.append(fold);
    }

    return Input{
        .points = points.items,
        .folds = folds.items,
    };
}

fn drawPoints(points: []Point) void {
    var result: u32 = 0;

    var min = Point{
        .x = 99999,
        .y = 99999,
    };
    var max = Point{
        .x = 0,
        .y = 0,
    };

    for (points) |point| {
        if (point.x < min.x) {
            min.x = point.x;
        }
        if (point.y < min.y) {
            min.y = point.y;
        }
        if (point.x > max.x) {
            max.x = point.x;
        }
        if (point.y > max.y) {
            max.y = point.y;
        }
    }

    var line: [2048]u8 = undefined;
    std.log.info("\npoints", .{});
    var y: u32 = min.y;
    while (y <= max.y) : (y += 1) {
        var x: u32 = min.x;
        while (x <= max.x) : (x += 1) {
            var found = false;
            for (points) |point| {
                if (point.x == x and point.y == y) {
                    found = true;
                    result += 1;
                    break;
                }
            }
            line[x] = if (found) '#' else '.';
        }
        std.log.info("{s}", .{line[0..x]});
    }
}

pub fn countPoints(points: []Point) u32 {
    var result: u32 = 0;

    var min = Point{
        .x = 99999,
        .y = 99999,
    };
    var max = Point{
        .x = 0,
        .y = 0,
    };

    for (points) |point| {
        if (point.x < min.x) {
            min.x = point.x;
        }
        if (point.y < min.y) {
            min.y = point.y;
        }
        if (point.x > max.x) {
            max.x = point.x;
        }
        if (point.y > max.y) {
            max.y = point.y;
        }
    }

    var y: u32 = min.y;
    while (y <= max.y) : (y += 1) {
        var x: u32 = min.x;
        while (x <= max.x) : (x += 1) {
            for (points) |point| {
                if (point.x == x and point.y == y) {
                    result += 1;
                    break;
                }
            }
        }
    }

    return result;
}

pub fn foldPoints(fold: Point, points: []Point) void {
    for (points) |point, i| {
        if (fold.x > fold.y) {
            if (point.x > fold.x) {
                points[i].x = fold.x - (point.x - fold.x);
            }
        } else {
            if (point.y > fold.y) {
                points[i].y = fold.y - (point.y - fold.y);
            }
        }
    }
}

pub fn task1(allocator: std.mem.Allocator, input: Input) !u32 {
    var result: u32 = 0;

    var points = try allocator.dupe(Point, input.points);
    defer allocator.free(points);
    var folds = try allocator.dupe(Point, input.folds);
    defer allocator.free(folds);

    foldPoints(folds[0], points);
    // drawPoints(points);
    result = countPoints(points);

    return result;
}

pub fn task2(allocator: std.mem.Allocator, input: Input) !u32 {
    var result: u32 = 0;

    var points = try allocator.dupe(Point, input.points);
    defer allocator.free(points);
    var folds = try allocator.dupe(Point, input.folds);
    defer allocator.free(folds);

    for (folds) |fold| {
        foldPoints(fold, points);
    }
    drawPoints(points);

    return result;
}

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const file = try std.fs.cwd().openFile("input.txt", .{ .read = true });
    defer file.close();

    const input = try readInput(allocator, file.reader());

    const task_1_result = try task1(allocator, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = try task2(allocator, input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
