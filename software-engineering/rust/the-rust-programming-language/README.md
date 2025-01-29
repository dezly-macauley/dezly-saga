# The Rust Programming Language
_______________________________________________________________________________
## If you have just `git cloned` this repo

Run the following command to build the project dependencies:
```
cargo build
```

_______________________________________________________________________________
## How to run your code

To run the `main`binary
```
cargo run --quiet --bin the-rust-programming-language
```
_______________________________________________________________________________
To run a specific binary:
```
cargo run --quiet --bin name-of-binary
```
_______________________________________________________________________________
## How to run check to see if your code compiles without running it

`cargo check` does not produce a binary executable, 
nor does it display the output of the code, 
so it is faster than `cargo build` and `cargo run`.

To check the `main`binary
```
cargo check --quiet --bin the-rust-programming-language
```
_______________________________________________________________________________
To check a specific binary:
```
cargo check --quiet --bin name-of-binary
```
_______________________________________________________________________________
## How to compile a ready for release (production ready) binary of your code

Fo the `main`binary
```
cargo run --release --quiet --bin the-rust-programming-language
```
_______________________________________________________________________________
For a specific binary:
```
cargo run --release --quiet --bin name-of-binary
```

_______________________________________________________________________________
