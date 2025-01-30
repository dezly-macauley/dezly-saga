 In many cases, 
 Rust requires you to acknowledge the possibility of an error 
 and take some action before your code will compile.
 This requirement makes your program more robust by ensuring 
 that you’ll discover errors and handle them appropriately 
 before you’ve deployed your code to production!

1. Recoverable Errors:
- Handled with the `Result<T, E>` enum
- E.g. A `file not found error`, 
we most likely just want to report the problem to the user 
and retry the operation.

2. Unrecoverable Errors:
- Handled with the `panic!` macro
- Always symptoms of bugs, 
such as trying to access a location beyond the end of an array, 
and so we want to immediately stop the program.
_______________________________________________________________________________
