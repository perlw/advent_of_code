const std = @import("std");

const Input = struct {
    draw_numbers: []u32 = undefined,
    boards: [][25]u32 = undefined,
};

fn readInputFile(allocator: *std.mem.Allocator, filename: []const u8) anyerror!Input {
    var result = Input{};
    var draw_numbers = std.ArrayList(u32).init(allocator);
    var boards = std.ArrayList([25]u32).init(allocator);

    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    const reader = file.reader();

    var line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch unreachable;

    {
        var token = std.mem.split(line, ",");
        while (token.next()) |slice| {
            const number = try std.fmt.parseUnsigned(u32, slice, 10);
            try draw_numbers.append(number);
        }
        allocator.free(line);
    }
    result.draw_numbers = draw_numbers.items;

    while (true) {
        // NOTE: Skip line.
        _ = reader.readByte() catch break;

        var i: u32 = 0;
        var board = std.mem.zeroes([25]u32);
        var board_index: u32 = 0;
        while (i < 5) : (i += 1) {
            line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch unreachable;
            var token = std.mem.tokenize(line, " ");
            while (token.next()) |slice| {
                board[board_index] = try std.fmt.parseUnsigned(u32, slice, 10);
                board_index += 1;
            }
            allocator.free(line);
        }
        try boards.append(board);
    }
    result.boards = boards.items;

    return result;
}

const Marker: u32 = 999;

inline fn boardTickNumber(number: u32, board_w: u32, board_h: u32, board: []u32) void {
    for (board) |value, i| {
        if (value == number) {
            board[i] = Marker;
        }
    }
}

inline fn checkBoard(board_w: u32, board_h: u32, board: []u32) bool {
    var x: u32 = 0;
    var y: u32 = 0;

    while (y < board_h) : (y += 1) {
        var foundNumber = false;
        x = 0;
        while (x < board_w) : (x += 1) {
            if (board[(y * board_w) + x] != Marker) {
                foundNumber = true;
            }
        }
        if (!foundNumber) {
            return true;
        }
    }

    while (x < board_w) : (x += 1) {
        var foundNumber = false;
        y = 0;
        while (y < board_h) : (y += 1) {
            if (board[(y * board_w) + x] != Marker) {
                foundNumber = true;
            }
        }
        if (!foundNumber) {
            return true;
        }
    }

    return false;
}

inline fn getBoardValue(board: []u32) u32 {
    var result: u32 = 0;
    for (board) |value| {
        if (value != Marker) {
            result += value;
        }
    }
    return result;
}

pub fn task1(draw_numbers: []const u32, board_w: u32, board_h: u32, boards: [][25]u32) u32 {
    var result: u32 = 0;

    done: for (draw_numbers) |number| {
        std.log.debug("drew {}", .{number});

        for (boards) |*b, i| {
            var board = b[0..];

            boardTickNumber(number, board_w, board_h, board);
            if (checkBoard(board_w, board_h, board)) {
                result = getBoardValue(board) * number;
                break :done;
            }

            std.log.debug("board {}:{any}", .{ i, board.* });
        }
    }

    return result;
}

pub fn main() anyerror!void {
    var buffer: [65536]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    var allocator = &fixed_buffer.allocator;

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(input.draw_numbers, 5, 5, input.boards);
    std.log.info("Task 1 result: {}", .{task_1_result});

    // const task_2_result = task2(...);
    // std.log.info("Task 2 result: {}", .{task_2_result});
}
