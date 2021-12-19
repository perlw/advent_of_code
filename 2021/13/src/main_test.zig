const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to result in 17 points" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\6,10
        \\0,14
        \\9,10
        \\0,3
        \\10,4
        \\4,11
        \\6,0
        \\6,12
        \\4,1
        \\0,13
        \\10,12
        \\3,4
        \\3,0
        \\8,4
        \\1,10
        \\2,14
        \\8,10
        \\9,0
        \\
        \\fold along y=7
        \\fold along x=5
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    const input = try app.readInput(std.testing.allocator, stream.reader());
    defer std.testing.allocator.free(input.points);
    defer std.testing.allocator.free(input.folds);
    const expected: u32 = 17;

    try std.testing.expect((try app.task1(std.testing.allocator, input)) == expected);
}

test "expect task 2 to result in a square" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\6,10
        \\0,14
        \\9,10
        \\0,3
        \\10,4
        \\4,11
        \\6,0
        \\6,12
        \\4,1
        \\0,13
        \\10,12
        \\3,4
        \\3,0
        \\8,4
        \\1,10
        \\2,14
        \\8,10
        \\9,0
        \\
        \\fold along y=7
        \\fold along x=5
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    const input = try app.readInput(std.testing.allocator, stream.reader());
    defer std.testing.allocator.free(input.points);
    defer std.testing.allocator.free(input.folds);

    _ = try app.task2(std.testing.allocator, input);
}
