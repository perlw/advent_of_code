const std = @import("std");
const builtin = @import("builtin");

pub const CommandType = enum {
    unknown,
    forward,
    down,
    up,
};

pub const Command = struct {
    cmd: CommandType,
    value: i32,
};

fn readInputFile(allocator: std.mem.Allocator, filename: []const u8) anyerror![]Command {
    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    var result = std.ArrayList(Command).init(allocator);

    var index: u32 = 0;
    var buffer: [16]u8 = undefined;
    const reader = file.reader();
    die: while (true) {
        const byte = reader.readByte() catch break :die;
        switch (byte) {
            ' ' => {
                var cmd: CommandType = .unknown;
                if (std.ascii.eqlIgnoreCase(buffer[0..index], "forward")) {
                    cmd = .forward;
                } else if (std.ascii.eqlIgnoreCase(buffer[0..index], "down")) {
                    cmd = .down;
                } else if (std.ascii.eqlIgnoreCase(buffer[0..index], "up")) {
                    cmd = .up;
                }
                const value = (try reader.readByte()) - 48;
                try result.append(.{ .cmd = cmd, .value = value });

                _ = reader.readByte() catch break :die;
                if (builtin.os.tag == .windows) {
                    // NOTE: Read another byte on windows due to two-byte line-endings.
                    _ = reader.readByte() catch break :die;
                }
                index = 0;
            },
            else => {
                buffer[index] = byte;
                index += 1;
            },
        }
    }

    return result.items;
}

pub fn task1(commands: []const Command) i32 {
    var horiz_pos: i32 = 0;
    var depth: i32 = 0;

    for (commands) |command| {
        switch (command.cmd) {
            .forward => {
                horiz_pos += command.value;
            },
            .down => {
                depth += command.value;
            },
            .up => {
                depth -= command.value;
            },
            else => {},
        }
    }

    return horiz_pos * depth;
}

pub fn task2(commands: []const Command) i32 {
    var horiz_pos: i32 = 0;
    var depth: i32 = 0;
    var aim: i32 = 0;

    for (commands) |command| {
        switch (command.cmd) {
            .forward => {
                horiz_pos += command.value;
                depth += aim * command.value;
            },
            .down => {
                aim += command.value;
            },
            .up => {
                aim -= command.value;
            },
            else => {},
        }
    }

    return horiz_pos * depth;
}

pub fn main() anyerror!void {
    var buffer: [65536]u8 = undefined;
    const allocator = std.heap.FixedBufferAllocator.init(&buffer).allocator();

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
