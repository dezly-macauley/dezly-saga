const print = @import("std").debug.print;

pub fn main() void {
    const value: i32 = 5;

    if (value < 5) {
        print("Value is less than 5\n", .{});
    } else if (value == 5) {
        print("Value is exactly 5\n", .{});
    } else {
        print("The value is greater than 5\n", .{});
    }
}
