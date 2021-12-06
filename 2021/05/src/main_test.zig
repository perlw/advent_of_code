const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to yield correct result" {
    var input = [_]app.Vent{
        .{ .a = .{ .x = 0, .y = 9 }, .b = .{ .x = 5, .y = 9 } },
        .{ .a = .{ .x = 8, .y = 0 }, .b = .{ .x = 0, .y = 8 } },
        .{ .a = .{ .x = 9, .y = 4 }, .b = .{ .x = 3, .y = 4 } },
        .{ .a = .{ .x = 2, .y = 2 }, .b = .{ .x = 2, .y = 1 } },
        .{ .a = .{ .x = 7, .y = 0 }, .b = .{ .x = 7, .y = 4 } },
        .{ .a = .{ .x = 6, .y = 4 }, .b = .{ .x = 2, .y = 0 } },
        .{ .a = .{ .x = 0, .y = 9 }, .b = .{ .x = 2, .y = 9 } },
        .{ .a = .{ .x = 3, .y = 4 }, .b = .{ .x = 1, .y = 4 } },
        .{ .a = .{ .x = 0, .y = 0 }, .b = .{ .x = 8, .y = 8 } },
        .{ .a = .{ .x = 5, .y = 5 }, .b = .{ .x = 8, .y = 2 } },
    };
    const expected: u32 = 5;

    try std.testing.expect((try app.task1(std.testing.allocator, .{ .x = 10, .y = 10 }, input[0..])) == expected);
}

test "expect task 2 to yield correct result" {
    var input = [_]app.Vent{
        .{ .a = .{ .x = 0, .y = 9 }, .b = .{ .x = 5, .y = 9 } },
        .{ .a = .{ .x = 8, .y = 0 }, .b = .{ .x = 0, .y = 8 } },
        .{ .a = .{ .x = 9, .y = 4 }, .b = .{ .x = 3, .y = 4 } },
        .{ .a = .{ .x = 2, .y = 2 }, .b = .{ .x = 2, .y = 1 } },
        .{ .a = .{ .x = 7, .y = 0 }, .b = .{ .x = 7, .y = 4 } },
        .{ .a = .{ .x = 6, .y = 4 }, .b = .{ .x = 2, .y = 0 } },
        .{ .a = .{ .x = 0, .y = 9 }, .b = .{ .x = 2, .y = 9 } },
        .{ .a = .{ .x = 3, .y = 4 }, .b = .{ .x = 1, .y = 4 } },
        .{ .a = .{ .x = 0, .y = 0 }, .b = .{ .x = 8, .y = 8 } },
        .{ .a = .{ .x = 5, .y = 5 }, .b = .{ .x = 8, .y = 2 } },
    };
    const expected: u32 = 12;

    std.testing.log_level = .debug;

    try std.testing.expect((try app.task2(std.testing.allocator, .{ .x = 10, .y = 10 }, input[0..])) == expected);
}
