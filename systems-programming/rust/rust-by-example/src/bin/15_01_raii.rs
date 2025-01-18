fn main() {

    create_integer_on_the_heap();
    
}

fn create_integer_on_the_heap() {

    // This is a 32 bit integer that is stored on the heap
    // for the duration of the function.
    let _heap_integer: Box<i32> = Box::new(42); 

    // NOTE: The underscore before the name has nothing to do with this
    // concept. This is just a way to get rid of the 
    // `unused variable in Rust`

    // After this function is called heap_integer will be discarded after
    // the function completes, and the memory will be freed.

}

//_____________________________________________________________________________
