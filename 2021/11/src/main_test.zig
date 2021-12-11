const std = @import("std");

const app = @import("./main.zig");

test "expect correct 2 step flash" {
    var input = [_]u32{
        1, 1, 1, 1, 1,
        1, 9, 9, 9, 1,
        1, 9, 1, 9, 1,
        1, 9, 9, 9, 1,
        1, 1, 1, 1, 1,
    };
    const expected = [_]u32{
        4, 5, 6, 5, 4,
        5, 1, 1, 1, 5,
        6, 1, 1, 1, 6,
        5, 1, 1, 1, 5,
        4, 5, 6, 5, 4,
    };

    std.testing.log_level = .debug;

    std.log.debug("\ninitial", .{});
    app.drawGrid(5, 5, &input);
    _ = app.iterateEnergy(5, 5, &input);
    std.log.debug("\nafter 1 step", .{});
    app.drawGrid(5, 5, &input);
    _ = app.iterateEnergy(5, 5, &input);
    std.log.debug("\nafter 2 step", .{});
    app.drawGrid(5, 5, &input);

    try std.testing.expect(std.mem.eql(u32, &input, &expected));
}

test "expect correct 1 step flash" {
    var input = [_]u32{
        6, 5, 9, 4, 2, 5, 4, 3, 3, 4,
        3, 8, 5, 6, 9, 6, 5, 8, 2, 2,
        6, 3, 7, 5, 6, 6, 7, 2, 8, 4,
        7, 2, 5, 2, 4, 4, 7, 2, 5, 7,
        7, 4, 6, 8, 4, 9, 6, 5, 8, 9,
        5, 2, 7, 8, 6, 3, 5, 7, 5, 6,
        3, 2, 8, 7, 9, 5, 2, 8, 3, 2,
        7, 9, 9, 3, 9, 9, 2, 2, 4, 5,
        5, 9, 5, 7, 9, 5, 9, 6, 6, 5,
        6, 3, 9, 4, 8, 6, 2, 6, 3, 7,
    };
    const expected = [_]u32{
        8, 8, 0, 7, 4, 7, 6, 5, 5, 5,
        5, 0, 8, 9, 0, 8, 7, 0, 5, 4,
        8, 5, 9, 7, 8, 8, 9, 6, 0, 8,
        8, 4, 8, 5, 7, 6, 9, 6, 0, 0,
        8, 7, 0, 0, 9, 0, 8, 8, 0, 0,
        6, 6, 0, 0, 0, 8, 8, 9, 8, 9,
        6, 8, 0, 0, 0, 0, 5, 9, 4, 3,
        0, 0, 0, 0, 0, 0, 7, 4, 5, 6,
        9, 0, 0, 0, 0, 0, 0, 8, 7, 6,
        8, 7, 0, 0, 0, 0, 6, 8, 4, 8,
    };

    std.testing.log_level = .debug;

    std.log.debug("\ninitial", .{});
    app.drawGrid(10, 10, &input);
    _ = app.iterateEnergy(10, 10, &input);
    std.log.debug("\nafter 1 step", .{});
    app.drawGrid(10, 10, &input);

    try std.testing.expect(std.mem.eql(u32, &input, &expected));
}

test "expect task 1 to result in 1656 flashes" {
    var input = [_]u32{
        5, 4, 8, 3, 1, 4, 3, 2, 2, 3,
        2, 7, 4, 5, 8, 5, 4, 7, 1, 1,
        5, 2, 6, 4, 5, 5, 6, 1, 7, 3,
        6, 1, 4, 1, 3, 3, 6, 1, 4, 6,
        6, 3, 5, 7, 3, 8, 5, 4, 7, 8,
        4, 1, 6, 7, 5, 2, 4, 6, 4, 5,
        2, 1, 7, 6, 8, 4, 1, 7, 2, 1,
        6, 8, 8, 2, 8, 8, 1, 1, 3, 4,
        4, 8, 4, 6, 8, 4, 8, 5, 5, 4,
        5, 2, 8, 3, 7, 5, 1, 5, 2, 6,
    };
    const expected: u32 = 1656;

    std.testing.log_level = .debug;

    try std.testing.expect(app.task1(10, 10, &input) == expected);
}

test "expect simultaneous flash after 2 steps" {
    var input = [_]u32{
        5, 8, 7, 7, 7, 7, 7, 7, 7, 7,
        8, 8, 7, 7, 7, 7, 7, 7, 7, 7,
        7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
        7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
        7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
        7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
        7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
        7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
        7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
        7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
    };
    const expected = [_]u32{0} ** 100;

    std.testing.log_level = .debug;

    std.log.debug("\ninitial", .{});
    app.drawGrid(10, 10, &input);
    _ = app.iterateEnergy(10, 10, &input);
    std.log.debug("\nafter 1 step", .{});
    app.drawGrid(10, 10, &input);
    _ = app.iterateEnergy(10, 10, &input);
    std.log.debug("\nafter 2 step", .{});
    app.drawGrid(10, 10, &input);

    std.testing.log_level = .debug;

    try std.testing.expect(std.mem.eql(u32, &input, &expected));
}

test "expect task 2 to result in 195 steps" {
    var input = [_]u32{
        5, 4, 8, 3, 1, 4, 3, 2, 2, 3,
        2, 7, 4, 5, 8, 5, 4, 7, 1, 1,
        5, 2, 6, 4, 5, 5, 6, 1, 7, 3,
        6, 1, 4, 1, 3, 3, 6, 1, 4, 6,
        6, 3, 5, 7, 3, 8, 5, 4, 7, 8,
        4, 1, 6, 7, 5, 2, 4, 6, 4, 5,
        2, 1, 7, 6, 8, 4, 1, 7, 2, 1,
        6, 8, 8, 2, 8, 8, 1, 1, 3, 4,
        4, 8, 4, 6, 8, 4, 8, 5, 5, 4,
        5, 2, 8, 3, 7, 5, 1, 5, 2, 6,
    };
    const expected: u32 = 195;

    std.testing.log_level = .debug;

    try std.testing.expect(app.task2(10, 10, &input) == expected);
}
