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
    var result: u64 = 0;
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
        var result: u64 = 0;
        std.debug.print("Expect: {s}->{}..", .{ t[0], t[1] });
        _ = try app.readOperatorPacket(std.testing.allocator, &input, &version, &result);
        checkPrint(true);
    }
    try std.testing.expect(ok);
}

test "expect runPacketDecode to produce expected versions" {
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
        var result: u64 = 0;
        try app.runPacketDecode(std.testing.allocator, &input, &version, &result);
        checkPrint(version == t[1]);
        if (version != t[1]) {
            ok = false;
        }
    }
    try std.testing.expect(ok);
}

test "expect runPacketDecode to produce expected results" {
    const tests = .{
        .{ "C200B40A82", 3 },
        .{ "04005AC33890", 54 },
        .{ "880086C3E88112", 7 },
        .{ "CE00C43D881120", 9 },
        .{ "D8005AC2A8F0", 1 },
        .{ "F600BC2D8F", 0 },
        .{ "9C005AC2F8F0", 0 },
        .{ "9C0141080250320F1802104A08", 1 },
    };

    std.debug.print("\n", .{});
    var ok = true;
    inline for (tests) |t| {
        var input = try app.hexToBinaryArrayList(std.testing.allocator, t[0]);
        defer input.deinit();

        std.debug.print("Expect: {s} -> v{}", .{ t[0], t[1] });
        var version: u32 = 0;
        var result: u64 = 0;
        try app.runPacketDecode(std.testing.allocator, &input, &version, &result);
        checkPrint(result == t[1]);
        if (result != t[1]) {
            ok = false;
        }
    }
    try std.testing.expect(ok);
}
