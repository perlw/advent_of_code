const std = @import("std");

const Input = struct {
    allocator: *std.mem.Allocator,

    template: []const u8,
    rules: std.StringHashMap(u8),

    pub fn init(allocator: *std.mem.Allocator) !Input {
        return Input{
            .allocator = allocator,
            .template = undefined,
            .rules = std.StringHashMap(u8).init(allocator),
        };
    }

    pub fn deinit(self: *Input) void {
        self.allocator.free(self.template);
        var it = self.rules.iterator();
        while (it.next()) |entry| {
            self.allocator.free(entry.key_ptr.*);
        }
        self.rules.deinit();
    }
};

pub fn readInput(allocator: *std.mem.Allocator, reader: anytype) !Input {
    var result = try Input.init(allocator);

    result.template = try reader.readUntilDelimiterAlloc(allocator, '\n', 512);
    _ = try reader.readByte();

    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        defer allocator.free(line);

        var it = std.mem.split(line, " -> ");
        var key = try allocator.dupe(u8, it.next().?);
        try result.rules.put(key, (it.next().?)[0]);
    }

    return result;
}

pub fn task1(allocator: *std.mem.Allocator, input: Input) !u32 {
    var current = try allocator.dupe(u8, input.template);
    defer allocator.free(current);
    var chain = std.ArrayList(u8).init(allocator);
    defer chain.deinit();

    std.log.info("initial {s}", .{current});
    var i: u32 = 0;
    while (i < 10) : (i += 1) {
        var j: u32 = 0;
        while (j < current.len - 1) : (j += 1) {
            const key = current[j .. j + 2];
            const insert = input.rules.get(key).?;
            if (chain.items.len == 0) {
                try chain.append(key[0]);
            }
            try chain.append(insert);
            try chain.append(key[1]);
        }
        allocator.free(current);
        current = try allocator.dupe(u8, chain.items);
        chain.clearAndFree();

        // std.log.debug("#{}: {s}", .{ i + 1, current });
    }

    std.log.info("final {s}", .{current});
    var counts = [_]u32{0} ** ('Z' - 'A');
    for (current) |c| {
        counts[c - 'A'] += 1;
    }
    var min: u32 = 999999;
    var max: u32 = 0;
    for (counts) |c| {
        if (c > 0 and c < min) {
            min = c;
        }
        if (c > max) {
            max = c;
        }
    }

    return (max - min);
}

pub fn task2(allocator: *std.mem.Allocator, input: Input) !u32 {
    var result: u32 = 0;

    return result;
}

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    var allocator = &fixed_buffer.allocator;

    const file = try std.fs.cwd().openFile("input.txt", .{ .read = true });
    defer file.close();

    var input = try readInput(allocator, file.reader());
    defer input.deinit();

    const task_1_result = try task1(allocator, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = try task2(allocator, input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
