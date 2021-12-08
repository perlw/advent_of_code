const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to result in 37 fuel" {
    var input = [_]u32{
        16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
    };
    const expected: u32 = 37;

    try std.testing.expect(app.task1(&input) == expected);
}

test "expect task 2 to result in 168 fuel" {
    var input = [_]u32{
        16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
    };
    const expected: u32 = 168;

    try std.testing.expect(app.task2(&input) == expected);
}
