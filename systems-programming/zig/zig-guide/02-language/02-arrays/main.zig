const std = @import("std");

pub fn main() void {

    // In Zig characters like `h` and `e` are actually stored as integers that
    // represent their ASCII or Unicode values.

    // 'h' → 104 (ASCII value)
    // 'e' → 101 (ASCII value)
    // 'l' → 108 (ASCII value)
    // 'o' → 111 (ASCII value)

    // This is an array that holds five 8-bit unsigned integers
    const array_a = [5]u8{ 'h', 'e', 'l', 'l', 'o' };

    // There is also a convient `_` syntax that will make Zig automatically
    // set the length of an array when it is declared
    const array_b = [_]u8{ 'w', 'o', 'r', 'l', 'd' };

    // NOTE: How to get the length of an array
    const array_a_length = array_a.len;
    std.debug.print("The length of the array_a is: {}\n", .{array_a_length});

    const array_b_length = array_b.len;
    std.debug.print("The length of the array_b is: {}\n", .{array_b_length});
}
