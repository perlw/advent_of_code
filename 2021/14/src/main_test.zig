const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to result in 1588" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\NNCB
        \\
        \\CH -> B
        \\HH -> N
        \\CB -> H
        \\NH -> C
        \\HB -> C
        \\HC -> B
        \\HN -> C
        \\NN -> C
        \\BH -> H
        \\NC -> B
        \\NB -> B
        \\BN -> B
        \\BB -> N
        \\BC -> B
        \\CC -> N
        \\CN -> C
        \\
    ;
    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();
    const expected: u32 = 1588;

    try std.testing.expect((try app.task1(std.testing.allocator, input)) == expected);
}
