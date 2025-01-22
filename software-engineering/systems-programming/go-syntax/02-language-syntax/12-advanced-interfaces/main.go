package main

type Putter interface {
    Put(id int, value any) (error)
    fmt.Stringer
}

type Storage interface {
  
    // NOTE: An interface can implement another interface
    Putter

    // This function accepts an id which is an integer, 
    // and it returns two values: 
    // any, 
    // and an error status (If it was success then error = nil,
    // If if is failure then an error message is returned)
    Get(id int) (any, error)
}

//=============================================================================

type simplePutter struct {}

func (instanceOfSimplePutter *simplePutter) Put(id int, val any) error {
    return nil
}

func (instanceOfSimplePutter *simplePutter) String() string {
    return ""
}

//=============================================================================

type FooStorage struct {}

// FooStorage because this is a function that does not modify the instance
// It will simply read from the instance
func (instanceOfFooStorage FooStorage) Get(id int) (any, error) {
    return nil, nil
}

// *FooStorage because it this is a function that will modify the instance
func (instanceOfFooStorage *FooStorage) Put(id int, value any) (error) {
    return nil
}

//=============================================================================

type BarStorage struct {}

// BarStorage because this is a function that does not modify the instance
// It will simply read from the instance
func (instanceOfBarStorage BarStorage) Get(id int) (any, error) {
    return nil, nil
}

// *BarStorage because it this is a function that will modify the instance
func (instanceOfBarStorage *BarStorage) Put(id int, value any) (error) {
    return nil
}


//=============================================================================

type Server struct {
    store Storage
}

func updateValue(id int, val any, putter Putter) error {
    return putter.Put(id, val)
}

func main() {
    sputter := &simplePutter{}
    updateValue(1, "foo", sputter)

}
