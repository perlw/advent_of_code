const std = @import("std");

const app = @import("./main.zig");

test "expect task 1 to result in 26 instances" {
    var input = [_]app.Display{
        .{
            .definitions = &.{
                "be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb",
            },
            .digits = &.{
                "fdgacbe", "cefdb", "cefbgd", "gcbe",
            },
        },
        .{
            .definitions = &.{
                "edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec",
            },
            .digits = &.{
                "fcgedb", "cgb", "dgebacf", "gc",
            },
        },
        .{
            .definitions = &.{
                "fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef",
            },
            .digits = &.{
                "cg", "cg", "fdcagb", "cbg",
            },
        },
        .{
            .definitions = &.{
                "fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega",
            },
            .digits = &.{
                "efabcd", "cedba", "gadfec", "cb",
            },
        },
        .{
            .definitions = &.{
                "aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga",
            },
            .digits = &.{
                "gecf", "egdcabf", "bgf", "bfgea",
            },
        },
        .{
            .definitions = &.{
                "fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf",
            },
            .digits = &.{
                "gebdcfa", "ecba", "ca", "fadegcb",
            },
        },
        .{
            .definitions = &.{
                "dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf",
            },
            .digits = &.{
                "cefg", "dcbef", "fcge", "gbcadfe",
            },
        },
        .{
            .definitions = &.{
                "bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd",
            },
            .digits = &.{
                "ed", "bcgafe", "cdgba", "cbgef",
            },
        },
        .{
            .definitions = &.{
                "egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg",
            },
            .digits = &.{
                "gbdfcae", "bgc", "cg", "cgb",
            },
        },
        .{
            .definitions = &.{
                "gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc",
            },
            .digits = &.{
                "fgae", "cfgab", "fg", "bagce",
            },
        },
    };
    const expected: u32 = 26;

    std.testing.log_level = .debug;

    try std.testing.expect(app.task1(&input) == expected);
}

test "expect task 2 to result in 5353" {
    var input = [_]app.Display{
        .{
            .definitions = &.{
                "acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab",
            },
            .digits = &.{
                "cdfeb", "fcadb", "cdfeb", "cdbaf",
            },
        },
    };
    const expected: u32 = 5353;

    std.testing.log_level = .debug;

    var result = try app.task2(std.testing.allocator, &input);
    try std.testing.expect(result[0] == expected);
    std.testing.allocator.free(result);
}

test "expect task 2 to result in correct solutions" {
    var input = [_]app.Display{
        .{
            .definitions = &.{
                "be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb",
            },
            .digits = &.{
                "fdgacbe", "cefdb", "cefbgd", "gcbe",
            },
        },
        .{
            .definitions = &.{
                "edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec",
            },
            .digits = &.{
                "fcgedb", "cgb", "dgebacf", "gc",
            },
        },
        .{
            .definitions = &.{
                "fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef",
            },
            .digits = &.{
                "cg", "cg", "fdcagb", "cbg",
            },
        },
        .{
            .definitions = &.{
                "fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega",
            },
            .digits = &.{
                "efabcd", "cedba", "gadfec", "cb",
            },
        },
        .{
            .definitions = &.{
                "aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga",
            },
            .digits = &.{
                "gecf", "egdcabf", "bgf", "bfgea",
            },
        },
        .{
            .definitions = &.{
                "fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf",
            },
            .digits = &.{
                "gebdcfa", "ecba", "ca", "fadegcb",
            },
        },
        .{
            .definitions = &.{
                "dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf",
            },
            .digits = &.{
                "cefg", "dcbef", "fcge", "gbcadfe",
            },
        },
        .{
            .definitions = &.{
                "bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd",
            },
            .digits = &.{
                "ed", "bcgafe", "cdgba", "cbgef",
            },
        },
        .{
            .definitions = &.{
                "egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg",
            },
            .digits = &.{
                "gbdfcae", "bgc", "cg", "cgb",
            },
        },
        .{
            .definitions = &.{
                "gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc",
            },
            .digits = &.{
                "fgae", "cfgab", "fg", "bagce",
            },
        },
    };
    const expected = [_]u32{
        8394,
        9781,
        1197,
        9361,
        4873,
        8418,
        4548,
        1625,
        8717,
        4315,
    };

    std.testing.log_level = .debug;

    var result = try app.task2(std.testing.allocator, &input);
    for (expected) |_, i| {
        try std.testing.expect(result[i] == expected[i]);
    }
    std.testing.allocator.free(result);
}

test "expect task 2 to sum to 61229" {
    var input = [_]app.Display{
        .{
            .definitions = &.{
                "be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb",
            },
            .digits = &.{
                "fdgacbe", "cefdb", "cefbgd", "gcbe",
            },
        },
        .{
            .definitions = &.{
                "edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec",
            },
            .digits = &.{
                "fcgedb", "cgb", "dgebacf", "gc",
            },
        },
        .{
            .definitions = &.{
                "fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef",
            },
            .digits = &.{
                "cg", "cg", "fdcagb", "cbg",
            },
        },
        .{
            .definitions = &.{
                "fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega",
            },
            .digits = &.{
                "efabcd", "cedba", "gadfec", "cb",
            },
        },
        .{
            .definitions = &.{
                "aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga",
            },
            .digits = &.{
                "gecf", "egdcabf", "bgf", "bfgea",
            },
        },
        .{
            .definitions = &.{
                "fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf",
            },
            .digits = &.{
                "gebdcfa", "ecba", "ca", "fadegcb",
            },
        },
        .{
            .definitions = &.{
                "dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf",
            },
            .digits = &.{
                "cefg", "dcbef", "fcge", "gbcadfe",
            },
        },
        .{
            .definitions = &.{
                "bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd",
            },
            .digits = &.{
                "ed", "bcgafe", "cdgba", "cbgef",
            },
        },
        .{
            .definitions = &.{
                "egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg",
            },
            .digits = &.{
                "gbdfcae", "bgc", "cg", "cgb",
            },
        },
        .{
            .definitions = &.{
                "gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc",
            },
            .digits = &.{
                "fgae", "cfgab", "fg", "bagce",
            },
        },
    };
    const expected: u32 = 61229;

    std.testing.log_level = .debug;

    var count: u32 = 0;
    var result = try app.task2(std.testing.allocator, &input);
    for (result) |r| {
        count += r;
    }
    std.testing.allocator.free(result);

    try std.testing.expect(count == expected);
}
