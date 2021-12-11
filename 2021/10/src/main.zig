const std = @import("std");

fn readInputFile(allocator: *std.mem.Allocator, filename: []const u8) ![][]u8 {
    var result = std.ArrayList([]u8).init(allocator);

    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    const reader = file.reader();

    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        try result.append(line);
    }

    return result.items;
}

pub fn isLineCorrupt(allocator: *std.mem.Allocator, line: []const u8) !bool {
    return (try corruptLineValue(allocator, line)) != 0;
}

fn corruptLineValue(allocator: *std.mem.Allocator, line: []const u8) !u32 {
    var open_chars = std.ArrayList(u8).init(allocator);
    defer open_chars.deinit();

    for (line) |char| {
        switch (char) {
            '(', '[', '{', '<' => {
                try open_chars.append(char);
            },
            else => {
                const popped_char = open_chars.popOrNull();

                if (popped_char) |last_open_char| {
                    var ok: bool = true;
                    if (last_open_char == '(' and char != ')') {
                        ok = false;
                    } else if (last_open_char == '[' and char != ']') {
                        ok = false;
                    } else if (last_open_char == '{' and char != '}') {
                        ok = false;
                    } else if (last_open_char == '<' and char != '>') {
                        ok = false;
                    }

                    if (!ok) {
                        return switch (char) {
                            ')' => 3,
                            ']' => 57,
                            '}' => 1197,
                            '>' => 25137,
                            else => 0,
                        };
                    }
                } else {
                    return 0;
                }
            },
        }
    }

    return 0;
}

fn getCompletedLineScore(allocator: *std.mem.Allocator, line: []const u8) !u64 {
    var result: u64 = 0;

    var open_chars = std.ArrayList(u8).init(allocator);
    defer open_chars.deinit();

    for (line) |char| {
        switch (char) {
            '(', '[', '{', '<' => {
                try open_chars.append(char);
            },
            else => {
                _ = open_chars.pop();
            },
        }
    }

    while (open_chars.popOrNull()) |char| {
        result = (result * 5) + @as(u64, switch (char) {
            '(' => 1,
            '[' => 2,
            '{' => 3,
            '<' => 4,
            else => 0,
        });
    }

    return result;
}

pub fn task1(allocator: *std.mem.Allocator, input: [][]const u8) !u32 {
    var result: u32 = 0;

    for (input) |line| {
        const score = try corruptLineValue(allocator, line);
        std.log.debug("Corrupted line: {s}={}", .{ line, score });
        result += score;
    }

    return result;
}

const asc_u64 = std.sort.asc(u64);

pub fn task2(allocator: *std.mem.Allocator, input: [][]const u8) !u64 {
    var scores = std.ArrayList(u64).init(allocator);
    defer scores.deinit();

    for (input) |line| {
        if (!(try isLineCorrupt(allocator, line))) {
            const score = try getCompletedLineScore(allocator, line);
            std.log.debug("Incomplete line: {s}={}", .{ line, score });
            if (score > 0) {
                try scores.append(score);
            }
        }
    }

    std.sort.sort(u64, scores.items, {}, asc_u64);
    return scores.items[scores.items.len / 2];
}

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    var allocator = &fixed_buffer.allocator;

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(allocator, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(allocator, input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
