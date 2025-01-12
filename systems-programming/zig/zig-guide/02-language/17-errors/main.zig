// Importing the print function from the standard library
const print = @import("std").debug.print;

pub fn main() void {

    //_________________________________________________________________________
    // SECTION: Zig's Approach to Error Handling
    
    // Zig uses `Error Sets` for handling errors in the program
    // An `Error Set` is basically a way to create a group of 
    // possible reasons why an action failed in a program.
    // E.g. Let's say a user tries to use the program to open a file 
    // and this action fails.
    // According to this `Error Set` called `FileOpenError`,
    // there are three possible reasons:

    // const FileOpenError = error{
    //     AccessDenied,
    //     OutOfMemory,
    //     FileNotFound,
    // };

    //_________________________________________________________________________
    // SECTION: `!` The Error Union Type

    // `!` can be used to create variable that is of the 
    // type `error union`.
    // A variable of this type can return two possible values.
    
    const my_string = "Hello, Zig!";
    std.debug.print("Type of my_int: {}\n", .{@TypeOf(my_int)});

    // Basically, i
    // This is another Error set
    const AllocationError = error{
        OutOfMememory
    }
    
    //_________________________________________________________________________

}


