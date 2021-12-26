const std = @import("std");
const builtin = @import("builtin");

// NOTE: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
const segment_counts = [_]u32{ 6, 2, 5, 5, 4, 5, 6, 3, 7, 6 };

pub const Display = struct {
    definitions: []const []const u8,
    digits: []const []const u8,
};

fn readInputFile(allocator: std.mem.Allocator, filename: []const u8) ![]Display {
    var result = std.ArrayList(Display).init(allocator);

    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    const reader = file.reader();

    while (true) {
        var line: []u8 = undefined;
        if (builtin.os.tag == .windows) {
            // NOTE: Read another byte on windows due to two-byte eol.
            line = reader.readUntilDelimiterAlloc(allocator, '\r', 512) catch break;
            _ = try reader.readByte();
        } else {
            line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        }
        defer allocator.free(line);

        var display: Display = undefined;

        var parts = std.mem.split(u8, line, " | ");

        var definitions = std.ArrayList([]u8).init(allocator);
        var it = std.mem.split(u8, parts.next().?, " ");
        while (it.next()) |slice| {
            var def = try allocator.dupe(u8, slice);
            try definitions.append(def);
        }
        display.definitions = definitions.items;

        var digits = std.ArrayList([]u8).init(allocator);
        it = std.mem.split(u8, parts.next().?, " ");
        while (it.next()) |slice| {
            var dig = try allocator.dupe(u8, slice);
            try digits.append(dig);
        }
        display.digits = digits.items;

        try result.append(display);
    }

    return result.items;
}

pub fn task1(input: []const Display) u32 {
    var result: u32 = 0;

    for (input) |display| {
        for (display.digits) |digit| {
            if (digit.len == segment_counts[1] or
                digit.len == segment_counts[4] or
                digit.len == segment_counts[7] or
                digit.len == segment_counts[8])
            {
                result += 1;
            }
        }
    }

    return result;
}

fn findMatch(segment_count: u32, base_on: usize, match_count: u32, partial: []usize, display: *const Display) usize {
    skip: for (display.definitions) |digit, index| {
        for (partial) |p| {
            if (p == index) {
                continue :skip;
            }
        }

        if (digit.len == segment_count) {
            var found: u32 = 0;
            for (display.definitions[base_on]) |def| {
                for (digit) |dig| {
                    if (def == dig) {
                        found += 1;
                    }
                }
            }

            if (found == match_count) {
                return index;
            }
        }
    }

    unreachable;
}

pub fn task2(allocator: std.mem.Allocator, input: []const Display) ![]u32 {
    var result = std.ArrayList(u32).init(allocator);

    for (input) |display| {
        var solution = [_]usize{99} ** 10;
        for (display.definitions) |digit, index| {
            if (digit.len == segment_counts[1]) {
                solution[1] = index;
                std.log.debug("1 {s}", .{digit});
            } else if (digit.len == segment_counts[4]) {
                solution[4] = index;
                std.log.debug("4 {s}", .{digit});
            } else if (digit.len == segment_counts[7]) {
                solution[7] = index;
                std.log.debug("7 {s}", .{digit});
            } else if (digit.len == segment_counts[8]) {
                solution[8] = index;
                std.log.debug("8 {s}", .{digit});
            }
        }

        solution[6] = findMatch(segment_counts[6], solution[1], 1, &solution, &display);
        std.log.debug("6 {s}", .{display.definitions[solution[6]]});
        solution[0] = findMatch(segment_counts[0], solution[4], 3, &solution, &display);
        std.log.debug("0 {s}", .{display.definitions[solution[0]]});
        solution[9] = findMatch(segment_counts[9], solution[0], 5, &solution, &display);
        std.log.debug("9 {s}", .{display.definitions[solution[9]]});
        solution[5] = findMatch(segment_counts[5], solution[6], 5, &solution, &display);
        std.log.debug("5 {s}", .{display.definitions[solution[5]]});
        solution[3] = findMatch(segment_counts[3], solution[9], 5, &solution, &display);
        std.log.debug("3 {s}", .{display.definitions[solution[3]]});
        solution[2] = findMatch(segment_counts[2], solution[8], 5, &solution, &display);
        std.log.debug("2 {s}", .{display.definitions[solution[2]]});

        std.log.debug("solution {any}", .{solution});

        var solve_result: u32 = 0;
        solve: for (display.digits) |dig| {
            std.log.debug("looking for {s}", .{dig});

            for (display.definitions) |def, def_index| {
                if (dig.len == def.len) {
                    std.log.debug("potential {s}#{}", .{ def, def.len });

                    var count: u32 = 0;
                    for (dig) |dig_char| {
                        for (def) |def_char| {
                            if (dig_char == def_char) {
                                count += 1;
                            }
                        }
                    }
                    if (count == def.len) {
                        std.log.debug("found! {s}#{}", .{ def, def.len });

                        for (solution) |sol, sol_index| {
                            if (sol == def_index) {
                                std.log.debug("={}", .{sol_index});
                                solve_result = (solve_result * 10) + @intCast(u32, sol_index);
                            }
                        }

                        continue :solve;
                    }
                }
            }
        }

        try result.append(solve_result);
    }

    return result.items;
}

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = try task2(allocator, input);
    var count: u32 = 0;
    for (task_2_result) |r| {
        count += r;
    }
    std.log.info("Task 2 result: {}", .{count});
}
