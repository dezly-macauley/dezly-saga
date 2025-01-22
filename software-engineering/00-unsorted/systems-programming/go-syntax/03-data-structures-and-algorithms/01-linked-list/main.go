package main

import "fmt"

// NOTE: What is a Linked List?

/*
    A linked list is a data structure that is simliar to a slice (or an array
    in other programming languages).

    The difference is that unlike an array where each value is stored in a 
    contiguous place in memory (each value is stored next to each other)...

    In a Linked List, the values are stored in random places in memory

*/

// NOTE: What does a Linked List contain
// A linked list is a data structure that is made of up a chain of Nodes
// So to create a Linked List you first need to create a Node

// NOTE: What is a Node?
// A Node is a simply data structure that has two values.
// 1. A value that it holds
// 2. And a pointer, that points to the next Node in the chain

//=============================================================================

// NOTE: Step 1: Create a Node

// In the example below I have chosen to call this LinkedListNode rather than
// just "Node" because "Node" is a reserved word in some programming languages,
// and also I want to make it clear that this is a Node for a Linked List
// Nodes for other data structures will be different

type LinkedListNode struct {
    value int
    next_node *LinkedListNode
}


// NOTE: Step 2: Create a function that will make it convinient to create a 
// new nodes

// This will create a new with a value, 
// and this node does not point to anything
func CreateLinkedListNode(_value int) *LinkedListNode {

    // This creates a new_node, sets its value, and returns the node
    new_node := &LinkedListNode {
        value: _value,
        next_node: nil,
    }

    return new_node 

}


// NOTE: Step 3: Create methods for the node


//=============================================================================

// NOTE: Step 4: Create the Linked Linked List
type LinkedList struct {

    // The first LinkedList Node on the chain
    head *LinkedListNode

    // The last LinkedList Node on the chain
    tail *LinkedListNode

    // A variable to keep track of the number of items in the Linked List
    number_of_nodes int

}

// NOTE: Step 5: Create a function that will create a New Linked List
// In Go the concept of a constructor function does not exist so you have to
// create a regular function to create it

// This function will create a new empty linked list
func CreateLinkedList() *LinkedList {

    // This creates a new linked list, set the head and tail pointers to nil
    // set the number of nodes to 0, and return it

    new_linked_list := &LinkedList {
        head: nil,
        tail: nil,
        number_of_nodes: 0,
    }

    return new_linked_list

}

// NOTE: Step 6: Create methods for the Linked List


// This method will add a new linked list node to the linked list 
// NOTE: Remember its (linked_list_instance *LinkedList)
// not (linked_list_instance LinkedList)
// If you are creating a method that modifies a value of an instance then you
// use the pointer *LinkedList. Using LinkedList simply creates a copy
 
func (linked_list_instance *LinkedList) Push(_value int) {

    // Step 1: Create a new node with a value 
    new_node := CreateLinkedListNode(_value)
    
    // However there is one more property that has to be set, next_node
    // Every LinkedListNode must point to the next node.
  
    // But... there is something else to consider...
    // When calling this method, the Linked List could be empty or it could
    // contain values
   
    // NOTE: Scenario 1: Adding a new node to an empty linked list
    // The new node becomes the head and tail 
    if linked_list_instance.head == nil && linked_list_instance.tail == nil {
        linked_list_instance.head = new_node 
        linked_list_instance.tail = new_node 
        // set the number_of_nodes to one
        linked_list_instance.number_of_nodes = 1
    } else {
        // NOTE: Scenario 2: Adding a new node to a linked list 
        // that already has values

        // First update the pointer of the tail node of the linked list
        // to point to the new node
        linked_list_instance.tail.next_node = new_node

        // And then make tail the new_node
        linked_list_instance.tail = new_node
       
        // Lastly, increase the number of nodes by one
        linked_list_instance.number_of_nodes++
        // This is the same as writing:
    }

}

//=============================================================================

func main() {

    // NOTE: Create an instance of the linked list
    linked_list_1 := CreateLinkedList()  
    linked_list_1.Push(7)
    linked_list_1.Push(4)

    fmt.Println(linked_list_1.head.value)
    // 7

    fmt.Println(linked_list_1.tail.value)
    // 4

    fmt.Println(linked_list_1.number_of_nodes)
    // 2

}
