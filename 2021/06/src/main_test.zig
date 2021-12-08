const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to yield 26 fish after 18 days" {
    var input = [_]u32{
        3, 4, 3, 1, 2,
    };
    const expected: u32 = 26;

    try std.testing.expect((try app.task1(std.testing.allocator, 18, input[0..])) == expected);
}

test "expect task 1 to yield 5934 fish after 80 days" {
    var input = [_]u32{
        3, 4, 3, 1, 2,
    };
    const expected: u32 = 5934;

    try std.testing.expect((try app.task1(std.testing.allocator, 80, input[0..])) == expected);
}

test "expect task 2 to yield 26 fish after 18 days" {
    var input = [_]u32{
        3, 4, 3, 1, 2,
    };
    const expected: u32 = 26;

    try std.testing.expect(app.task2(18, input[0..]) == expected);
}

test "expect task 2 to yield 5934 fish after 80 days" {
    var input = [_]u32{
        3, 4, 3, 1, 2,
    };
    const expected: u32 = 5934;

    try std.testing.expect(app.task2(80, input[0..]) == expected);
}

test "expect task 2 to yield 26984457539 fish after 256 days" {
    var input = [_]u32{
        3, 4, 3, 1, 2,
    };
    const expected: u64 = 26984457539;

    try std.testing.expect(app.task2(256, input[0..]) == expected);
}
