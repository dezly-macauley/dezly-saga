// NOTE: panic vs abort
// 1. By default, when a panic occurs the program starts unwinding, 
// which means Rust walks back up the stack and cleans up the data 
// from each function it encounters. 
// 2. abort: This ends the program immediately withour cleaning up

// NOTE: How to change this setting:

/*
    Open up the `cargo.toml` file in your project and add this line:

    [profile.release]
    panic = 'abort'
*/

fn main() {
    // How to explicitly call `panic!`
    panic!("crash and burn");
    // thread 'main' panicked at 
    // src/bin/09_01_unrecoverable_errors_with_panic.rs:18:5:
    // crash and burn
    // note: run with `RUST_BACKTRACE=1` 
    // environment variable to display a backtrace
}
