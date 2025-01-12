const print = @import("std").debug.print;

pub fn main() void {
    const player_attack_strength: u32 = 70;

    const enemy_message: []const u8 = switch (player_attack_strength) {
        // This includes 0 and 49
        0...49 => "Haha! You are so weak",

        // exactly 50
        50 => "You are just as strong as I am",

        // 60 or 70
        60, 70 => "That was a cheap shot",

        // Anything else
        else => "I surrender!!!",
    };

    print("The enemy_message is: {s}\n", .{enemy_message});
}
