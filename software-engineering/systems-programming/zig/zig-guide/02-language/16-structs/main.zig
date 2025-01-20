const print = @import("std").debug.print;

pub fn main() void {

    //_________________________________________________________________________
    // NOTE: Structs

    // A struct is used to group variables together,
    // to create a custom variable type
    const Player = struct {
        wins: u32,
        losses: u32,
    };

    // Creating an instance of the struct
    const red_player: Player = Player{ .wins = 20, .losses = 5 };

    print("The red player has {} wins and {} losses\n", .{ red_player.wins, red_player.losses });

    //_________________________________________________________________________
    // SECTION: Creating a struct with default values

    const Sword = struct {
        attack_damage: u32 = 10,
        armour_break: u32 = 15,
    };

    const cloud_katana: Sword = Sword{};

    print("attack_damage: {}, armour_break: {}\n", .{ cloud_katana.attack_damage, cloud_katana.armour_break });

    //_________________________________________________________________________

}
