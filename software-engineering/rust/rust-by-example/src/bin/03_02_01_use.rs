#[allow(dead_code)]
enum GameDifficulty {
    Normal,
    Expert
}

// NOTE: The options inside the enum are refered to as `enum variants`

#[allow(dead_code)]
enum GameStage {
    RoseField,
    DragonLair
}

fn main() {

    //_________________________________________________________________________
    // NOTE: How you would normally declare an enum 

    // let current_stage: GameStage = GameStage::DragonLair;
    // let current_difficulty: GameDifficulty = GameDifficulty::Easy;
    
    //_________________________________________________________________________
    // NOTE: How to simplify the declaration using the keyword `use` 
    // This allows you to use the enum variants directly

    // You can import specific enum variants
    use crate::GameStage::{ RoseField, DragonLair };
   
    // Or you can import all of them.
    use crate::GameDifficulty::*;

    let current_stage: GameStage = DragonLair;
    let current_difficulty: GameDifficulty = Expert;

    match current_stage {
       RoseField => println!("Let the Petals fly"),
       DragonLair => println!("Watch out for the Dragon")
    }

    match current_difficulty {
       Normal => println!("Normal Mode Selected: Have fun!"),
       Expert => println!("Expert Mode Selected: Be careful!")
    }

    //_________________________________________________________________________

}
