const std = @import("std");

fn readInputFile(allocator: std.mem.Allocator, filename: []const u8) anyerror![]u32 {
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

fn pickDiagnostics(comptime T: type, allocator: std.mem.Allocator, bit: u5, invert: bool, diagnostic: []const u32) []const u32 {
    var result = std.ArrayList(u32).init(allocator);

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

    message = 0;
    while (message < diagnostic.len) : (message += 1) {
        var bit_val = (diagnostic[message] >> (@bitSizeOf(T) - bit)) & 0x1;
        bit_val = if (invert) (~bit_val) & 0x1 else bit_val;
        if ((on >= off and bit_val == 1) or (on < off and bit_val == 0)) {
            result.append(diagnostic[message]) catch unreachable;
        }
    }

    return result.items;
}

pub fn task2(comptime T: type, diagnostic: []const u32) u32 {
    var buffer: [65536]u8 = undefined;
    const allocator = std.heap.FixedBufferAllocator.init(&buffer).allocator();

    var oxygen_rating: u32 = 0;
    var co2_rating: u32 = 0;

    var filtered_oxygen: []const u32 = allocator.dupe(u32, diagnostic) catch unreachable;
    var filtered_co2: []const u32 = allocator.dupe(u32, diagnostic) catch unreachable;
    var bit: u5 = 1;
    while (bit <= @bitSizeOf(T)) : (bit += 1) {
        if (filtered_oxygen.len > 1) {
            var prev = filtered_oxygen;
            filtered_oxygen = pickDiagnostics(T, allocator, bit, false, prev);
            allocator.free(prev);
        }
        if (filtered_co2.len > 1) {
            var prev = filtered_co2;
            filtered_co2 = pickDiagnostics(T, allocator, bit, true, prev);
            allocator.free(prev);
        }
    }
    oxygen_rating = filtered_oxygen[0];
    co2_rating = filtered_co2[0];

    std.log.debug("oxygen_rating {}", .{oxygen_rating});
    std.log.debug("co2_rating {}", .{co2_rating});

    return oxygen_rating * co2_rating;
}

pub fn main() anyerror!void {
    var buffer: [65536]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const input = try readInputFile(allocator, "input.txt");

    const task_1_result = task1(u12, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(u12, input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
