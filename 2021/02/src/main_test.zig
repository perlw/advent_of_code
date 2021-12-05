const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to yield correct result" {
    const input = [_]app.Command{
        .{ .cmd = .forward, .value = 5 },
        .{ .cmd = .down, .value = 5 },
        .{ .cmd = .forward, .value = 8 },
        .{ .cmd = .up, .value = 3 },
        .{ .cmd = .down, .value = 8 },
        .{ .cmd = .forward, .value = 2 },
    };
    const expected: i32 = 150;

    try std.testing.expect(app.task1(input[0..]) == expected);
}

test "expect task 2 to yield correct result" {
    const input = [_]app.Command{
        .{ .cmd = .forward, .value = 5 },
        .{ .cmd = .down, .value = 5 },
        .{ .cmd = .forward, .value = 8 },
        .{ .cmd = .up, .value = 3 },
        .{ .cmd = .down, .value = 8 },
        .{ .cmd = .forward, .value = 2 },
    };
    const expected: i32 = 900;

    try std.testing.expect(app.task2(input[0..]) == expected);
}
