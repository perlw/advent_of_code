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

test "expect charToBinary to produce expected result" {
    const tests = .{
        .{ '0', 0b0000 },
        .{ '1', 0b0001 },
        .{ '8', 0b1000 },
        .{ '9', 0b1001 },
        .{ 'A', 0b1010 },
        .{ 'B', 0b1011 },
        .{ 'E', 0b1110 },
        .{ 'F', 0b1111 },
    };

    std.debug.print("\n", .{});
    var ok = true;
    inline for (tests) |t| {
        std.debug.print("Expect: {}->{}..", .{ t[0], t[1] });
        checkPrint((try app.charToBinary(t[0])) == t[1]);
    }

    try std.testing.expect(ok);
}

test "expect hexToBinaryArrayList to produce expected result" {
    const tests = .{
        .{ "0123", @as(u16, 0b0000000100100011) },
        .{ "ABCD", @as(u16, 0b1010101111001101) },
    };

    std.debug.print("\n", .{});
    var ok = true;
    inline for (tests) |t| {
        var result = try app.hexToBinaryArrayList(std.testing.allocator, t[0][0..]);
        defer result.deinit();

        std.debug.print("Expect: {s}->{}..", .{ t[0], t[1] });
        var verified = true;
        for (result.items) |b, i| {
            if (b != (t[1] >> 15 - @intCast(u4, i)) & 0x1) {
                verified = false;
                ok = false;
                break;
            }
        }
        checkPrint(verified);
    }

    try std.testing.expect(ok);
}

test "expect readLiteralValue to produce expected result" {
    var input = std.ArrayList(u1).init(std.testing.allocator);
    defer input.deinit();
    try input.appendSlice(&[_]u1{
        1,
        1,
        0,
        1,
        0,
        0,
        1,
        0,
        1,
        1,
        1,
        1,
        1,
        1,
        1,
        0,
        0,
        0,
        1,
        0,
        1,
        0,
        0,
        0,
    });

    const expected: u32 = 2021;

    var version: u32 = 0;
    var result: u16 = 0;
    _ = try app.readLiteralValue(&input, &version, &result);
    try std.testing.expect(result == expected);
}

test "expect readOperatorPacket to produce expected result" {
    const tests = .{
        .{ "38006F45291200", 0 },
        .{ "EE00D40C823060", 0 },
    };

    std.debug.print("\n", .{});
    var ok = true;
    inline for (tests) |t| {
        var input = try app.hexToBinaryArrayList(std.testing.allocator, t[0]);
        defer input.deinit();

        var version: u32 = 0;
        var result: u16 = 0;
        std.debug.print("Expect: {s}->{}..", .{ t[0], t[1] });
        _ = try app.readOperatorPacket(&input, &version, &result);
        checkPrint(true);
    }
    try std.testing.expect(ok);
}

test "expect runPacketDecode to produce expected result" {
    const tests = .{
        .{ "8A004A801A8002F478", 16 },
        .{ "620080001611562C8802118E34", 12 },
        .{ "C0015000016115A2E0802F182340", 23 },
        .{ "A0016C880162017C3686B18A3D4780", 31 },
    };

    std.debug.print("\n", .{});
    var ok = true;
    inline for (tests) |t| {
        var input = try app.hexToBinaryArrayList(std.testing.allocator, t[0]);
        defer input.deinit();

        std.debug.print("Expect: {s} -> v{}", .{ t[0], t[1] });
        var version: u32 = 0;
        var result: u16 = 0;
        try app.runPacketDecode(&input, &version, &result);
        checkPrint(version == t[1]);
    }
    try std.testing.expect(ok);
}

test "expect task 1 to result in 31" {
    std.testing.log_level = .debug;

    const raw_input: []const u8 =
        \\A0016C880162017C3686B18A3D4780
        \\
    ;

    var stream = std.io.fixedBufferStream(raw_input);
    var input = try app.readInput(std.testing.allocator, stream.reader());
    defer input.deinit();

    const expected: u32 = 31;

    try std.testing.expect((try app.task1(&input)) == expected);
}
