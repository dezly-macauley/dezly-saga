const std = @import("std");

pub fn main() void {
    std.debug.print("Hello, {s}!\n", .{"World"});
}

// NOTE: To run this file:
// zig run name_of_file.zig
