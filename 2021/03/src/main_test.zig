const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to yield correct result" {
    const input = [_]u32{
        0b00100,
        0b11110,
        0b10110,
        0b10111,
        0b10101,
        0b01111,
        0b00111,
        0b11100,
        0b10000,
        0b11001,
        0b00010,
        0b01010,
    };
    const expected: u32 = 198;

    std.testing.log_level = .debug;

    try std.testing.expect(app.task1(u5, input[0..]) == expected);
}

test "expect task 2 to yield correct result" {}
