package main

import "fmt"

type Position struct {
    x int
    y int
}

type Entity struct {
    name string
    id string
    version string

    // Entity will now have the fields of the "Postion" struct
    Position
}

type SpecialEntity struct {

    // NOTE: Struct embedding
    // The struct SpecialEntity will have the same fields as the Entity struct
    // This will save you from having to type everything out

    Entity
    // name string
    // id string
    // version string
    // posx int
    // posy int
    
    // So this struct has all the fields of Entity, plus its own unique fields
    specialField float64
}

func main() {

    // NOTE: How to declare an instance of the SpecialEntity struct
    my_entity := SpecialEntity {

        Entity: Entity {
            name: "my special entity",
            version: "1.1",
            id: "my special id",

            Position: Position {
                x: 100,
                y: 200,
            },

        },

        specialField: 88.88,
    }

    fmt.Printf("%+v\n", my_entity)
    // {Entity:{name:my special entity id:my special id version:1.1 Position:{x:100 y:200}} specialField:88.88}

    // NOTE: You can access the fields and change their properties in the same
    // way as a regular struct
    
    my_entity.x = 300
    my_entity.y = 450
    my_entity.name = "Batman entity"
    
    fmt.Printf("%+v\n", my_entity)
    // {Entity:{name:Batman entity id:my special id version:1.1 Position:{x:300 y:450}} specialField:88.88}

}
