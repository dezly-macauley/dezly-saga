const std = @import("std");
const expect = std.testing.expect;

test "always succeeds" {
    try expect(true);
}

// NOTE: How to run a test:

// zig test test_pass.zig
