const print = @import("std").debug.print;

pub fn main() void {

    //_________________________________________________________________________
    // SECTION: Basic Syntax
    // const (or var) name_of_varible: the_type_of_the_variable = value;

    //_________________________________________________________________________
    // SECTION: Integers (Negative and positive numbers allowed)

    const player_score: i32 = -15;
    print("Your score is {}\n", .{player_score});
    // Your score is -15

    //_________________________________________________________________________

    // NOTE: How to print the variable type of a variable in Zig

    const type_of_player_score = @TypeOf(player_score);
    print("player_score is of the type: {}\n", .{type_of_player_score});
    // player_score is of the type: i32

    //_________________________________________________________________________
    // SECTION: Unsigned Integers (Only positive numbers allowed)

    const number_of_ninjas: u32 = 5;
    print("There are {} ninjas\n", .{number_of_ninjas});
    // There are 5 ninjas

    //_________________________________________________________________________
    // SECTION: booleans

    const has_magic: bool = true;
    print("The player can use magic: {}\n", .{has_magic});
    // The player can use magic: true

    const is_team_leader: bool = false;
    print("The player is the team leader: {}\n", .{is_team_leader});
    // The player is the team leader: false

    //_________________________________________________________________________
    // SECTION: String Slice (Immutable, dynamic length)

    // String slices are immutable,
    // meaning you cannot change the contents of the slice.

    // NOTE: Breakdown of []const u8:

    // []: This indicates a slice. It means the array can be of variable length.
    // const: This makes the contents immutable.
    // You can't modify the elements of the slice.
    // u8: This is the type of the elements in the slice.
    // u8 represents an 8-bit unsigned integer,
    // which is the fundamental type used for bytes.

    const player_name: []const u8 = "Dezly Macauley";
    print("Hello {s}\n", .{player_name});

    //_________________________________________________________________________
    // SECTION: Byte Arry (Mutable, fixed sized)

    // const my_byte_arry: [3]u8 = [3]u8{ 81, 21, 4 };
    // print(my_byte_arry[0]);

    //_________________________________________________________________________
    // SECTION: Unused variables

    // By default Zig will give a warning about unused variables
    // You can avoid the compiler warning by using an underscore.

    const ninja_stars: u32 = 71;

    // NOTE: This tells the compiler:
    // "I know that I am not using the ninja_stars variable I created
    // in this program. Don't warn me about it"
    _ = ninja_stars;

    //_________________________________________________________________________
}
