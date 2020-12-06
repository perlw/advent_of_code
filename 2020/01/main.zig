const std = @import("std");
const fs = std.fs;
const fmt = std.fmt;

fn asc(context: void, a: i32, b: i32) bool {
  return a > b;
}

pub fn task1(input: []i32) i32 {
  std.sort.sort(i32, input, {}, asc);

  var i: usize = 0;
  while (i < input.len - 1) : (i+=1) {
    var j: usize = i + 1;
    while (j < input.len - 1) : (j+=1) {
      if (input[i]+input[j] == 2020) {
        return input[i] * input[j];
      }
    }
  }

  return -1;
}

pub fn task2(input: []i32) i32 {
  std.sort.sort(i32, input, {}, asc);

  var i: usize = 0;
  while (i < input.len - 1) : (i+=1) {
    var j: usize = i + 1;
    while (j < input.len - 1) : (j+=1) {
      var k: usize = j + 1;
      while (k < input.len - 1) : (k+=1) {
        if (input[i]+input[j]+input[k] == 2020) {
          return input[i] * input[j] * input[k];
        }
      }
    }
  }

  return -1;
}

pub fn main() !void {
  const stdout = std.io.getStdOut().writer();
  var buffer: [1024]u8 = undefined;
  const allocator = &std.heap.FixedBufferAllocator.init(&buffer).allocator;

  const file = try fs.cwd().openFile("input.txt", fs.File.OpenFlags{
    .read = true,
  });
  defer file.close();

  var input = std.ArrayList(i32).init(allocator);
  const reader = file.reader();
  while (true) {
    const data = reader
      .readUntilDelimiterAlloc(std.heap.page_allocator, '\n', 32) catch break;

    try input.append(try fmt.parseInt(i32, data, 10));
  }

  var result = task1(input.items);
  try stdout.print("Task 1: {}\n", .{result});

  result = task2(input.items);
  try stdout.print("Task 2: {}\n", .{result});
}
