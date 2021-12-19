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

fn printPairs(pairs: std.StringHashMap(u64)) void {
    var count: u32 = 0;
    var it = pairs.iterator();
    while (it.next()) |entry| {
        std.log.debug("{s}=>{}", .{ entry.key_ptr.*, entry.value_ptr.* });
        count += 1;
    }
    std.log.debug("num pairs {}", .{count});
}

pub fn task(allocator: *std.mem.Allocator, steps: u32, input: Input) !u64 {
    var pairs = std.StringHashMap(u64).init(allocator);
    defer {
        var it = pairs.iterator();
        while (it.next()) |entry| {
            allocator.free(entry.key_ptr.*);
        }
        pairs.deinit();
    }

    std.log.info("initial {s}", .{input.template});
    var i: u32 = 0;
    while (i < input.template.len - 1) : (i += 1) {
        const key = input.template[i .. i + 2];
        if (pairs.contains(key)) {
            try pairs.put(key, pairs.get(key).? + 1);
        } else {
            try pairs.put(try allocator.dupe(u8, key), 1);
        }
    }
    printPairs(pairs);

    i = 0;
    while (i < steps) : (i += 1) {
        var new_pairs = std.StringHashMap(u64).init(allocator);

        var it = pairs.iterator();
        while (it.next()) |entry| {
            const key = entry.key_ptr.*;
            const value = entry.value_ptr.*;

            const insert = input.rules.get(key).?;
            const new_pair_1 = [2]u8{ key[0], insert };
            const new_pair_2 = [2]u8{ insert, key[1] };
            const new_value_1 = (new_pairs.get(new_pair_1[0..]) orelse 0) + value;
            const new_value_2 = (new_pairs.get(new_pair_2[0..]) orelse 0) + value;

            if (new_pairs.contains(new_pair_1[0..])) {
                try new_pairs.put(new_pair_1[0..], new_value_1);
            } else {
                try new_pairs.put(try allocator.dupe(u8, new_pair_1[0..]), new_value_1);
            }
            if (new_pairs.contains(new_pair_2[0..])) {
                try new_pairs.put(new_pair_2[0..], new_value_2);
            } else {
                try new_pairs.put(try allocator.dupe(u8, new_pair_2[0..]), new_value_2);
            }
            //std.log.debug("new pairs from {}x{s}=>{s},{s}", .{ value, key, new_pair_1, new_pair_2 });
            //std.log.debug("new values from {},{}", .{ new_pairs.get(new_pair_1[0..]), new_pairs.get(new_pair_2[0..]) });
        }
        it = pairs.iterator();
        while (it.next()) |entry| {
            allocator.free(entry.key_ptr.*);
        }
        pairs.deinit();

        pairs = new_pairs;
        //printPairs(pairs);
    }

    printPairs(pairs);
    var counts = [_]u64{0} ** ('Z' - 'A');
    var it = pairs.iterator();
    while (it.next()) |entry| {
        const key = entry.key_ptr.*;
        const value = entry.value_ptr.*;
        counts[key[0] - 'A'] += value;
        counts[key[1] - 'A'] += value;
    }
    for (counts) |*c| {
        c.* = (c.* + 1) / 2;
    }
    var min: u64 = std.math.maxInt(u64);
    var max: u64 = 0;
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

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    var allocator = &fixed_buffer.allocator;

    const file = try std.fs.cwd().openFile("input.txt", .{ .read = true });
    defer file.close();

    var input = try readInput(allocator, file.reader());
    defer input.deinit();

    const task_1_result = try task(allocator, 10, input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = try task(allocator, 40, input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
