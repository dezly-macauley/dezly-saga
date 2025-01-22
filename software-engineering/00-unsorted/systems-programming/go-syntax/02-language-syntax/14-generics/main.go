package main

// NOTE: What are generics?
// Basically its a function that can work with a very flexible range of 
// arguements

import "fmt"

// NOTE: The [Key comparable, Value any] after the struct name 
// is how you create a generic

type CustomMap[Key comparable, Value any] struct {
    data map[Key]Value
}

func NewCustomMap[Key comparable, Value any]() *CustomMap[Key, Value] {
    return &CustomMap[Key,Value] {

        // NOTE: This is how to create a map
        data: make(map[Key]Value),
    }
}

func (instanceOfCustomMap CustomMap[Key, Value]) Insert(key Key, value Value) error {
    instanceOfCustomMap.data[key] = value
    return nil
}

//=============================================================================

func foo[T any](val T) {
    fmt.Println(val)
}

//=============================================================================

func main() {

    map1 := NewCustomMap[string, int]()
    map1.Insert("foo", 1)
    
    map2 := NewCustomMap[int, float64]()
    map2.Insert(7, 92.22)

    foo("ka-boom")

}
