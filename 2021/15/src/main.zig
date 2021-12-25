const std = @import("std");

pub const invalid_index = std.math.inf_u32;
const Cell = struct {
    distance: u32,
    index: u32,
};

fn pop_next(front: *std.ArrayList(Cell)) Cell {
    var min_distance: u32 = std.math.inf_u32;
    var pop: usize = std.math.inf_u32;
    for (front.items) |cell, array_index| {
        if (cell.distance < min_distance) {
            min_distance = cell.distance;
            pop = array_index;
        }
    }
    return front.orderedRemove(pop);
}

pub const Map = struct {
    allocator: std.mem.Allocator,
    risk_levels: []u32,
    width: u32,
    height: u32,

    pub fn init(allocator: std.mem.Allocator, risk_levels: []u32, width: u32, height: u32) !Map {
        return Map{
            .allocator = allocator,
            .risk_levels = try allocator.dupe(u32, risk_levels),
            .width = width,
            .height = height,
        };
    }

    pub fn deinit(self: *Map) void {
        self.allocator.free(self.risk_levels);
    }

    pub inline fn neighbors(self: Map, index: u32) [4]u32 {
        const x: u32 = index % self.width;
        const y: u32 = index / self.width;
        return [4]u32{
            if (x > 0) index - 1 else invalid_index,
            if (x < self.width - 1) index + 1 else invalid_index,
            if (y > 0) index - self.width else invalid_index,
            if (y < self.height - 1) index + self.width else invalid_index,
        };
    }

    pub fn find_path(self: Map) !u32 {
        var front = std.ArrayList(Cell).init(self.allocator);
        defer front.deinit();
        var visited = std.ArrayList(u32).init(self.allocator);
        defer visited.deinit();
        var distances = try self.allocator.alloc(u32, self.risk_levels.len);
        defer self.allocator.free(distances);
        std.mem.set(u32, distances, std.math.inf_u32);

        try front.append(Cell{
            .distance = 0,
            .index = 0,
        });

        iterate: while (front.items.len > 0) {
            var cell = pop_next(&front);
            for (visited.items) |vis| {
                if (vis == cell.index) {
                    continue :iterate;
                }
            }
            try visited.append(cell.index);

            var nbors = self.neighbors(cell.index);
            for (nbors) |nbor| {
                if (nbor == invalid_index) {
                    continue;
                }

                var distance: u32 = cell.distance + self.risk_levels[nbor];
                if (distance < distances[nbor]) {
                    distances[nbor] = distance;
                }

                try front.append(Cell{
                    .distance = distance,
                    .index = nbor,
                });
            }
        }

        var target: u32 = (self.width * self.height) - 1;
        return distances[target];
    }
};

pub fn readInput(allocator: std.mem.Allocator, reader: anytype) !Map {
    var risk_levels = std.ArrayList(u32).init(allocator);
    defer risk_levels.deinit();

    var width: u32 = 0;
    var height: u32 = 0;
    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        defer allocator.free(line);

        if (line.len > width) {
            width = @intCast(u32, line.len);
        }
        height += 1;

        for (line) |char| {
            const risk: u32 = char - '0';
            try risk_levels.append(risk);
        }
    }

    return try Map.init(allocator, risk_levels.items, width, height);
}

pub fn task1(input: Map) !u32 {
    return input.find_path();
}

pub fn task2(input: Map) !u32 {
    _ = input;
    return 0;
}

pub fn main() !void {
    var buffer: [2000000]u8 = undefined;
    var fixed_buffer = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fixed_buffer.allocator();

    const file = try std.fs.cwd().openFile("input.txt", .{ .read = true });
    defer file.close();

    var input = try readInput(allocator, file.reader());
    defer input.deinit();

    const task_1_result = try task1(input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = try task2(input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
