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
    // SECTION: String Slice

    const player_name: []const u8 = "Dezly Macauley";
    print("Hello {s}\n", .{player_name});

    //_________________________________________________________________________

}
