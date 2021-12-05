const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to yield correct result" {
    const input = [_]u32{
        199,
        200,
        208,
        210,
        200,
        207,
        240,
        269,
        260,
        263,
    };
    const expected: u32 = 7;

    try std.testing.expect(app.task1(input[0..]) == expected);
}

test "expect task 2 to yield correct result" {
    const input = [_]u32{
        199,
        200,
        208,
        210,
        200,
        207,
        240,
        269,
        260,
        263,
    };
    const expected: u32 = 5;
    try std.testing.expect(app.task2(input[0..]) == expected);
}
