const std = @import("std");
const builtin = @import("builtin");

const app = @import("./main.zig");

inline fn checkPrint(ok: bool) void {
    if (ok) {
        std.debug.print(" ✅\n", .{});
    } else {
        std.debug.print(" ❌\n", .{});
    }
}

test "setup" {
    if (builtin.os.tag == .windows) {
        // NOTE: Make windows use UTF-8 for output.
        _ = std.os.windows.kernel32.SetConsoleOutputCP(65001);
    }
}

test "expect task 1 to result in 45" {
    const tests = .{
        .{
            \\target area: x=20..30, y=-10..-5
            \\
            ,
            45,
        },
    };

    std.debug.print("\n", .{});
    var ok = true;
    inline for (tests) |t| {
        var stream = std.io.fixedBufferStream(t[0]);
        var input = try app.readInput(std.testing.allocator, stream.reader());

        const result = app.task1(input);
        std.debug.print("Expect: {}, got {}", .{ t[1], result });
        checkPrint(result == t[1]);
        if (result != t[1]) {
            ok = false;
        }
    }
    try std.testing.expect(ok);
}

test "expect task 2 to result in 112" {
    const tests = .{
        .{
            \\target area: x=20..30, y=-10..-5
            \\
            ,
            112,
        },
    };

    std.debug.print("\n", .{});
    var ok = true;
    inline for (tests) |t| {
        var stream = std.io.fixedBufferStream(t[0]);
        var input = try app.readInput(std.testing.allocator, stream.reader());

        const result = app.task2(input);
        std.debug.print("Expect: {}, got {}", .{ t[1], result });
        checkPrint(result == t[1]);
        if (result != t[1]) {
            ok = false;
        }
    }
    try std.testing.expect(ok);
}
