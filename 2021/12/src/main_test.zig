const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 test 1 to result in 10" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\start-A
        \\start-b
        \\A-c
        \\A-b
        \\b-d
        \\A-end
        \\b-end
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();
    const expected: u32 = 10;

    try std.testing.expect((try app.task1(std.testing.allocator, &input)) == expected);
}

test "expect task 1 test 2 to result in 19" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\dc-end
        \\HN-start
        \\start-kj
        \\dc-start
        \\dc-HN
        \\LN-dc
        \\HN-end
        \\kj-sa
        \\kj-HN
        \\kj-dc
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();
    const expected: u32 = 19;

    try std.testing.expect((try app.task1(std.testing.allocator, &input)) == expected);
}

test "expect task 1 test 3 to result in 226" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\fs-end
        \\he-DX
        \\fs-he
        \\start-DX
        \\pj-DX
        \\end-zg
        \\zg-sl
        \\zg-pj
        \\pj-he
        \\RW-he
        \\fs-DX
        \\pj-RW
        \\zg-RW
        \\start-pj
        \\he-WI
        \\zg-he
        \\pj-fs
        \\start-RW
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();
    const expected: u32 = 226;

    try std.testing.expect((try app.task1(std.testing.allocator, &input)) == expected);
}

test "expect task 2 test 1 to result in 36" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\start-A
        \\start-b
        \\A-c
        \\A-b
        \\b-d
        \\A-end
        \\b-end
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();
    const expected: u32 = 36;

    try std.testing.expect((try app.task2(std.testing.allocator, &input)) == expected);
}

test "expect task 2 test 2 to result in 103" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\dc-end
        \\HN-start
        \\start-kj
        \\dc-start
        \\dc-HN
        \\LN-dc
        \\HN-end
        \\kj-sa
        \\kj-HN
        \\kj-dc
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();
    const expected: u32 = 103;

    try std.testing.expect((try app.task2(std.testing.allocator, &input)) == expected);
}

test "expect task 2 test 3 to result in 3509" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\fs-end
        \\he-DX
        \\fs-he
        \\start-DX
        \\pj-DX
        \\end-zg
        \\zg-sl
        \\zg-pj
        \\pj-he
        \\RW-he
        \\fs-DX
        \\pj-RW
        \\zg-RW
        \\start-pj
        \\he-WI
        \\zg-he
        \\pj-fs
        \\start-RW
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();
    const expected: u32 = 3509;

    try std.testing.expect((try app.task2(std.testing.allocator, &input)) == expected);
}
