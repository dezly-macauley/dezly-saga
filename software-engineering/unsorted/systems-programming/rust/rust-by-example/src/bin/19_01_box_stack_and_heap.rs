/*
    NOTE: Box, stack and heap

    Values can be boxed (allocated on the heap) by creating a Box<T>.
    A box is a smart pointer to a heap allocated value of type T.
    When a box goes out of scope, its destructor is called,
    the inner object is destroyed, and the memory on the heap is freed.

    Boxed values can be dereferenced using the * operator;
    this removes one layer of indirection.

*/

fn main() {

    // ABOUT: Memory Allocation
    // 1. The process of reserving a portion of memory for programs and
    // processes to run
    // -----------------------------------------------------------------
    // A value in memory can be stored on the stack or the heap
    // -----------------------------------------------------------------
    // [ The Stack ] 
    // 1. Variables that have a known size at compile time will be 
    // stored here
    // 2. The Stack follows the principle LIFO (Last In, First Out)
    // Think of it as a stack a plates, the plate at the top of the 
    // stack is removed first
    // -----------------------------------------------------------------
    // [ The Heap ] 
    // 1. Variables that have an unknown size at compile time will be
    // stored here
    // -----------------------------------------------------------------

    // NOTE: I'm just creating a struct to show how to declare this as a
    // heap variable

    #[derive(Debug)]
    struct UserInfo{
        id: i32,
        num_messages: i32,
    }

    //_________________________________________________________________________
    // SECTION: Stack Allocation
    // 1. All values in Rust are stack allocated by default. 

    let stack_stored_number: i32 = 42;
    println!("stack_stored_number: {}", stack_stored_number);
    // stack_stored_number: 42

    let stack_stored_user: UserInfo = UserInfo {
        id: 29,
        num_messages: 282,
    };

    println!("stack_stored_user: {:?}", stack_stored_user);
    // stack_stored_user: UserInfo { id: 29, num_messages: 282 }

    println!("id is {}", stack_stored_user.id);
    // id is 29

    println!("num_messages is {}", stack_stored_user.num_messages);
    // num_messages is 282

    //_________________________________________________________________________
    // SECTION: Heap Allocation
        
    // Values can be boxed (allocated on the heap) by creating a Box<T>. 
    // A box is a smart pointer to a heap allocated value of type T. 
    // When a box goes out of scope, 
    // its destructor is called, the inner object is destroyed,
    // and the memory on the heap is freed.

    // NOTE: Whe you store a value to the heap, 
    // what you are actually storing in the variable `heap_stored_number` is
    // a memory address (aka a pointer) that points to the location of where
    // that value is stored in memory.
    // You can use {:p} to see what the memory address is.
    // This address is only fixed at runtime. If you restart your computer 
    // again you may be given a completely different memory address.
    
    let heap_stored_number: Box<i32> = Box::new(20);

    println!("heap_stored_number is at the address {:p}", heap_stored_number);
    // heap_stored_number is at the address 0x605627155b50
    
    println!("The value of heap_stored_number is: {}", *heap_stored_number);
    // The value of heap_stored_number is: 20

    let heap_stored_user: Box<UserInfo> = Box::new(
        UserInfo {
            id: 29,
            num_messages: 282,
        }
    );

    println!("heap_stored_user contains: {:?}", *heap_stored_user);
    // heap_stored_user contains: UserInfo { id: 29, num_messages: 282 }

    // NOTE: The parenthesis are needed.
    // Dereference the entire struct, and then access the whole field

    println!("The user has {} messages", (*heap_stored_user).num_messages);
    // The user has 282 messages

    //_________________________________________________________________________

}
