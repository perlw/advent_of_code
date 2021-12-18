const std = @import("std");

const Cave = struct {
    links: std.ArrayList([]const u8),
};

const CaveSystem = struct {
    system: std.StringHashMap(Cave),
    allocator: *std.mem.Allocator,

    pub fn init(allocator: *std.mem.Allocator) !CaveSystem {
        return CaveSystem{
            .system = std.StringHashMap(Cave).init(allocator),
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *CaveSystem) void {
        var it = self.system.iterator();
        while (it.next()) |entry| {
            var value = entry.value_ptr.*;
            for (value.links.items) |item| {
                self.allocator.free(item);
            }
            value.links.deinit();

            self.allocator.free(entry.key_ptr.*);
        }
        self.system.deinit();
    }
};

pub fn readInput(allocator: *std.mem.Allocator, reader: anytype) !CaveSystem {
    var result = try CaveSystem.init(allocator);

    while (true) {
        const line = reader.readUntilDelimiterAlloc(allocator, '\n', 512) catch break;
        defer allocator.free(line);

        var it = std.mem.split(line, "-");

        var key = try allocator.dupe(u8, it.next().?);
        var free_the_key = false;
        var value = try allocator.dupe(u8, it.next().?);

        var res = try result.system.getOrPut(key);
        if (!res.found_existing) {
            var list = std.ArrayList([]const u8).init(allocator);
            res.value_ptr.* = Cave{
                .links = list,
            };
        } else {
            free_the_key = true;
        }
        try res.value_ptr.*.links.append(try allocator.dupe(u8, value));

        res = try result.system.getOrPut(value);
        if (!res.found_existing) {
            var list = std.ArrayList([]const u8).init(allocator);
            res.value_ptr.* = Cave{
                .links = list,
            };
        } else {
            allocator.free(value);
        }
        try res.value_ptr.*.links.append(try allocator.dupe(u8, key));
        if (free_the_key) {
            allocator.free(key);
        }
    }

    return result;
}

fn walkCaves(current_cave: []const u8, cave_system: *CaveSystem, visited_small_caves: *std.StringHashMap(u32), check_twice: bool, did_visit_twice: bool) u32 {
    if (std.mem.eql(u8, current_cave, "end")) {
        return 1;
    }

    var result: u32 = 0;

    var is_lower = true;
    for (current_cave) |c| {
        if (!std.ascii.isLower(c)) {
            is_lower = false;
            break;
        }
    }
    var checked_twice = did_visit_twice;
    if (is_lower) {
        if (visited_small_caves.contains(current_cave)) {
            checked_twice = true;
        }
        visited_small_caves.put(current_cave, 1) catch unreachable;
    }

    if (cave_system.system.get(current_cave)) |this_cave| {
        // std.log.debug("at: {s}", .{current_cave});
        for (this_cave.links.items) |cave| {
            if (std.mem.eql(u8, cave, "start")) {
                continue;
            }

            // std.log.debug("visiting: {s}", .{cave});
            if (check_twice) {
                if (visited_small_caves.contains(cave) and checked_twice) {
                    continue;
                }
            } else {
                if (visited_small_caves.contains(cave)) {
                    continue;
                }
            }

            var visited_clone = visited_small_caves.clone() catch unreachable;
            defer visited_clone.deinit();
            result += walkCaves(cave, cave_system, &visited_clone, check_twice, checked_twice);
        }
    }

    return result;
}

pub fn task1(allocator: *std.mem.Allocator, input: *CaveSystem) !u32 {
    var visited = std.StringHashMap(u32).init(allocator);
    defer visited.deinit();

    return walkCaves("start", input, &visited, false, false);
}

pub fn task2(allocator: *std.mem.Allocator, input: *CaveSystem) !u32 {
    var visited = std.StringHashMap(u32).init(allocator);
    defer visited.deinit();

    return walkCaves("start", input, &visited, true, false);
}

pub fn main() !void {
    var allocator = std.heap.page_allocator;

    const file = try std.fs.cwd().openFile("input.txt", .{ .read = true });
    defer file.close();

    var input = try readInput(allocator, file.reader());
    defer input.deinit();

    const task_1_result = task1(allocator, &input);
    std.log.info("Task 1 result: {}", .{task_1_result});

    const task_2_result = task2(allocator, &input);
    std.log.info("Task 2 result: {}", .{task_2_result});
}
