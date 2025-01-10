const std = @import("std");

pub fn main() void {
    // Value assignment has the following syntax: (const|var) identifier[: type] = value.

    // 32 bit integer
    const my_constant: i32 = 5;
    std.debug.print("my_constant: {}\n", .{my_constant});

    // 32 bit unsigned integer
    var my_variable: u32 = 5000;
    std.debug.print("my_variable: {}\n", .{my_variable});

    my_variable = 700;
    std.debug.print("my_variable: {}\n", .{my_variable});
}
