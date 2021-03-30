package main

import (
    "reflect"
    "fmt"
)


func test1() {
    const tagName = "json"

    type User struct {
        Id int `validate:"-"`
        Name string `validate:"presence,min=2,max=32"`
        Email string `validate:"email,required" json:"email"`
    }

    user := User{
        Id: 1,
        Name: "John Doe",
        Email: "john@example",
    }

    // TypeOf returns the reflection Type that represents the dynamic type of variable.
    // If variable is a nil interface value, TypeOf returns nil.
    t := reflect.TypeOf(user)

    //Get the type and kind of our user variable
    fmt.Println("Type: ", t.Name())
    fmt.Println("Kind: ", t.Kind())

    for i := 0; i < t.NumField(); i++ {
        // Get the field, returns https://golang.org/pkg/reflect/#StructField
        field := t.Field(i)

        //Get the field tag value
        tag := field.Tag.Get(tagName)

        fmt.Printf("%d. %v(%v), tag:'%v'\n", i+1, field.Name, field.Type.Name(), tag)
    }    
}

func test2() {
    type T1 struct {
        Id string
    }
    type T2 struct {
        T1
        Name string
    }

    t2 := T2 {}
    t2.Id = "aaa"
    fmt.Println(t2)
}

type A struct {
    id int
}

func newDefaultA() A {
    return A{id:10}
}    
    
func test3() {
    a := newDefaultA()
    fmt.Println(a)
}

func main() {
    test3()
}
