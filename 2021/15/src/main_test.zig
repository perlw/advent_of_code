const std = @import("std");

const app = @import("./main.zig");

test "expect neighbors to index 0 to have two valid indices" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\012
        \\345
        \\678
        \\
    ;

    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();

    const expected = [4]u32{
        app.invalid_index,
        1,
        app.invalid_index,
        3,
    };

    try std.testing.expect(std.mem.eql(u32, &input.neighbors(0), &expected));
}

test "expect neighbors to index 8 to have two valid indices" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\012
        \\345
        \\678
        \\
    ;

    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();

    const expected = [4]u32{
        7,
        app.invalid_index,
        5,
        app.invalid_index,
    };

    try std.testing.expect(std.mem.eql(u32, &input.neighbors(8), &expected));
}

test "expect neighbors to index 4 to have four valid indices" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\012
        \\345
        \\678
        \\
    ;

    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();

    const expected = [4]u32{
        3,
        5,
        1,
        7,
    };

    try std.testing.expect(std.mem.eql(u32, &input.neighbors(4), &expected));
}

test "expect task 1 to result in 40" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\1163751742
        \\1381373672
        \\2136511328
        \\3694931569
        \\7463417111
        \\1319128137
        \\1359912421
        \\3125421639
        \\1293138521
        \\2311944581
        \\
    ;

    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();

    const expected: u32 = 40;

    try std.testing.expect((try app.task1(input)) == expected);
}
