## How to create new libray project

```
cargo new --lib name-of-project
```

This will create the following structure
```
.
├── name-of-project
│   ├── Cargo.toml
│   └── src
│       └── lib.rs
```
_______________________________________________________________________________
## How to create new binary project

```
cargo new name-of-project
```

This will create the following structure
```
.
├── name-of-project
│   ├── Cargo.toml
│   └── src
│       └── main.rs
```

_______________________________________________________________________________

### The rest of this guide will focus on the binary project

`main.rs` is the root source file of your project.


E.g. `cargo new example-project`

```
.
├── example-project
│   ├── Cargo.toml
│   └── src
│       └── main.rs
```

_______________________________________________________________________________

Cargo.toml
```
[package]
name = "example-project"
version = "0.1.0"
authors = "Dezly Macauley"
edition = "2021"

[dependencies]
```
The name field under [package] determines the name of the project. 
This is used by crates.io 
if you publish the crate (more later). 
It is also the name of the output binary when you compile.

The version field is a crate version number using Semantic Versioning.

The authors field is a list of authors used when publishing the crate.

_______________________________________________________________________________
### How to build the project?

Run the command `cargo build` from any directory.

This will compile your Rust code and create an executable binary in the
directory `target/debug`

If you want to compile and run your program in one step:
```
cargo run --bin name-of-project
```

NOTE: If `main` is the only binary in your project then you can just
use the command:
```
cargo run
```

_______________________________________________________________________________
