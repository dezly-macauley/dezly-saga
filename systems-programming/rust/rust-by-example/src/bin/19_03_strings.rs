// ABOUT: Strings
// 1. The two most used string types in Rust are `String` and `&str`.
// 2. This is how you store text in Rust

// NOTE: String 
// ------------------------------------------------------------------------
// `u8` 
// 1. A`u8` is an 8 bit unsigned integer. 
// 2. It can only store positive numbers from 0 to 255 (including 255)
// In Rust this range is represented like this 0..=255
// ------------------------------------------------------------------------
// `byte`
// 1. A byte is a unit of memory that holds 8 bits.
// 2. a u8 is a convenient way to represent these bytes as integers.
// 3. Text encoding (like UTF-8) uses bytes to represent characters, 
// where each character maps to a value between 0-255
// ------------------------------------------------------------------------
// `What kind of characters fall under UTF-8?`
// 1. ASCII (American Standard Code for Information Interchange)
// 2. Unicode 
// ------------------------------------------------------------------------
// ASCII:
// - You can select from 128 different characters.
// - Limited to latin characters. Like A to Z, numbers,
// punctuation marks, and control characters like `\n` which means newline.
// ------------------------------------------------------------------------
// Unicode:
// - Includes the 128 ASCII characters
// - You can select from over 143,000 characters.
// - Includes languages that don't use the Latin script like Japanese 
// and Russian.
// - Includes Emojis
// ------------------------------------------------------------------------
// Unicode encoding schemes:
// UTF-8, UTF-16, and UTF-32 are the main encoding schemes for Unicode

// ------------------------------------------------------------------------
// `Vector`
// 1. A Vector is basically a list of items that can grow 
// to keep storing more items. In the this specific case the 
// 2. This a data structure that is stored on the Heap because it does
// not have a fixed size at compile time.
// ------------------------------------------------------------------------
// A String a Vector of bytes `Vec<u8>`, 
// but guaranteed to always be a valid UTF-8 sequence. 
// ------------------------------------------------------------------------
// A string is not null terminated. 
// In simple terms that means you have to add a newline character `\n`
// at the end of the string to tell Rust where the text ends.
// ------------------------------------------------------------------------

fn main() {


}
