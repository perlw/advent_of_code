const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to sum to 15" {
    var input = [_]u32{
        2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
        3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
        9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
        8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
        9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
    };
    const expected: u32 = 15;

    std.testing.log_level = .debug;

    try std.testing.expect(app.task1(10, 5, &input) == expected);
}

test "expect task 2 to sum to 1134" {
    var input = [_]u32{
        2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
        3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
        9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
        8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
        9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
    };
    const expected: u32 = 1134;

    std.testing.log_level = .debug;

    try std.testing.expect((try app.task2(std.testing.allocator, 10, 5, &input)) == expected);
}
