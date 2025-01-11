// This is import the print function from the Zig standard library
const print = @import("std").debug.print;

pub fn main() void {
    print("This is message 1\n", .{});

    defer print("This is message 2\n", .{});

    print("This is message 3\n", .{});

    // This is message 1
    // This is message 3
    // This is message 2
}
