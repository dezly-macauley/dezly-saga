const expect = @import("std").testing.expect;

// NOTE: In Zig All function arguments are immutable
// - if a copy is desired the user must explicitly make one.

// NOTE: In Zig variables use snake_case but functions used camelCase

fn addFive(x: u32) u32 {
    return x + 5;
}

// NOTE: The test block is used to test if a function is working correctly
test "desription of test" {
    const y = addFive(0);
    try expect(@TypeOf(y) == u32);
    try expect(y == 5);
}
