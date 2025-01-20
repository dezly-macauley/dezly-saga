If you want to create a Rust project that has multiple binaries,
E.g. two `main.rs` files, then all you need to do 
is create a directory called `src/bin` and place the other binaries there.

```
name-of-project
.
├── Cargo.lock
├── Cargo.toml
└── src
    ├── bin
    │   ├── second_main.rs
    │   └── third_main.rs
    └── main.rs

```
### When using this structure there are a few changes:
1. `cargo build` will work the same way. Except it will compile main.rs
and all of the binaries inside the `bin` directory

2. If you want to run `main.rs` then the commmand is:
`cargo run --bin name-of-project`

Using just `cargo run` will now give you an error because you have not
told Rust which one of the binaries in your project to run.

3. If you want to run one of the binaries inside `bin` then the command is:
`cargo run --bin second_main`
`cargo run --bin third_main`

NOTE: Cargo run is for binaries, not the file (second_main.rs)



