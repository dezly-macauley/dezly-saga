// NOTE: If statements accept `bool` values (true or false)

// Unlike C or JavaScript, there are no values that
// implicitly coerce to bool values

const expect = @import("std").testing.expect;

test "if statement" {
    const a = true;

    if (a) {
        x += 1;
    } else {
        x += 2;
    }
    try expect (x == 1)
}
