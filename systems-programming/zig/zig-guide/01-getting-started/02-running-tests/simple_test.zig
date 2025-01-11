// This will import the Zig standard library
const std = @import("std");

// This will make it convenient to use the `expect` function
// from the std library.
// So instead of having to write this:
// try std.testing.expect(total == 14);
// You can simplify that to:
// try expect(total == 14);
const expect = std.testing.expect;

test "Check if 4 + 10 equals 14, and num_two is 4" {
    const num_one: i32 = 10;
    const num_two: i32 = 4;

    const total: i32 = num_one + num_two;

    // The expect function is used to test if a variable or function is
    // working correctly. Inside the brackets all you need to put is the
    // variable name and what you expect its value to be.
    // E.g. total should be 14.
    try expect(total == 14);

    try expect(num_two == 4);
}

// NOTE: To run this file use the following command:
// zig simple_test.zig
