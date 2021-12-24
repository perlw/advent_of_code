const std = @import("std");

const app = @import("./main.zig");

test "expect lines to be corrupt" {
    var input = [_][]const u8{
        "(]",
        "{()()()>",
        "(((()))}",
        "<([]){()}[{}])",
    };

    std.testing.log_level = .debug;

    var buffer: [2000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();
    for (input) |line| {
        try std.testing.expect(try app.isLineCorrupt(allocator, line));
    }
}

test "expect lines to be okay" {
    var input = [_][]const u8{
        "{",
        "[]",
        "{}",
        "()",
        "<>",
        "([])",
        "{()()()}",
        "<([{}])>",
        "[<>({}){}[([])<>]]",
        "(((((((((())))))))))",
    };

    std.testing.log_level = .debug;

    var buffer: [2000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();
    for (input) |line| {
        try std.testing.expect(!(try app.isLineCorrupt(allocator, line)));
    }
}

test "expect task 1 to sum to 26397" {
    var input = [_][]const u8{
        "[({(<(())[]>[[{[]{<()<>>",
        "[(()[<>])]({[<{<<[]>>(",
        "{([(<{}[<>[]}>{[]{[(<()>",
        "(((({<>}<{<{<>}{[]{[]{}",
        "[[<[([]))<([[{}[[()]]]",
        "[{[{({}]{}}([{[{{{}}([]",
        "{<[[]]>}<{[{[{[]{()[[[]",
        "[<(<(<(<{}))><([]([]()",
        "<{([([[(<>()){}]>(<<{{",
        "<{([{{}}[<[[[<>{}]]]>[]]",
    };
    const expected: u32 = 26397;

    std.testing.log_level = .debug;

    var buffer: [2000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();
    try std.testing.expect((try app.task1(allocator, &input)) == expected);
}

test "expect task 2 to result in 288957" {
    var input = [_][]const u8{
        "[({(<(())[]>[[{[]{<()<>>",
        "[(()[<>])]({[<{<<[]>>(",
        "{([(<{}[<>[]}>{[]{[(<()>",
        "(((({<>}<{<{<>}{[]{[]{}",
        "[[<[([]))<([[{}[[()]]]",
        "[{[{({}]{}}([{[{{{}}([]",
        "{<[[]]>}<{[{[{[]{()[[[]",
        "[<(<(<(<{}))><([]([]()",
        "<{([([[(<>()){}]>(<<{{",
        "<{([{{}}[<[[[<>{}]]]>[]]",
    };
    const expected: u64 = 288957;

    std.testing.log_level = .debug;

    var buffer: [2000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();
    try std.testing.expect((try app.task2(allocator, &input)) == expected);
}
