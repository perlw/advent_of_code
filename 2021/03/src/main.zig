const std = @import("std");

fn readInputFile(allocator: *std.mem.Allocator, filename: []const u8) anyerror![]u32 {
    var result = std.ArrayList(u32).init(allocator);

    const file = try std.fs.cwd().openFile(filename, .{ .read = true });
    defer file.close();

    const reader = file.reader();
    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 16) catch break;
        defer allocator.free(line);

        var message: u32 = 0;
        for (line) |char| {
            message |= @as(u32, if (char == '1') 0x1 else 0x0);
            message <<= 1;
        }
        message >>= 1;
        try result.append(message);
    }

    return result.items;
}

pub fn task1(comptime T: type, diagnostic: []const u32) u32 {
    var gamma_rate: u32 = 0;
    var epsilon_rate: u32 = 0;

    var bit: u5 = 1;
    while (bit <= @bitSizeOf(T)) : (bit += 1) {
        var on: u32 = 0;
        var off: u32 = 0;

        var message: u32 = 0;
        while (message < diagnostic.len) : (message += 1) {
            const bit_val = (diagnostic[message] >> (@bitSizeOf(T) - bit)) & 0x1;
            if (bit_val == 1) {
                on += 1;
            } else {
                off += 1;
            }
        }

        gamma_rate |= @as(u32, if (on > off) 1 else 0);
        gamma_rate <<= 1;
        epsilon_rate |= @as(u32, if (on < off) 1 else 0);
        epsilon_rate <<= 1;
    }
    gamma_rate >>= 1;
    epsilon_rate >>= 1;

    std.log.debug("gamma_rate {}", .{gamma_rate});
    std.log.debug("epsilon_rate {}", .{epsilon_rate});

    return gamma_rate * epsilon_rate;
}

pub fn task2(diagnostic: []const []const u8) u32 {
    return 0;
}

pub fn main() anyerror!void {
    var buffer: [65536]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    var allocator = &fixed_buffer.allocator;

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(u12, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    //const task_2_result = task2(input);
    //std.log.info("Task 2 result: {}", .{task_2_result});
}
