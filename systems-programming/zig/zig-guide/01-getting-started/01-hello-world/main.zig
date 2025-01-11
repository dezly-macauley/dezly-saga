const std = @import("std");
const print = std.debug.print;

// NOTE: In Zig `main` is a special function that is the entry point into
// your program. It must be marked as public with the `pub` keyword.

pub fn main() void {
    print("Hello, {s}!\n", .{"World"});
}

// NOTE: To run this file:
// zig run name_of_file.zig
