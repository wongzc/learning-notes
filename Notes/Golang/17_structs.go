//https://gobyexample.com/structs
package main

import "fmt"
// structs are typed collection of fields
type person struct {
    name string // person struct, with name and age fields
    age  int
}

func newPerson(name string) *person {// function to construct a struct with given name

    p := person{name: name}
    p.age = 42
    return &p // retuned a pointer to local variable, wont clean by go garbage collector
}

func main() {

    fmt.Println(person{"Bob", 20}) // normal way to initialize

    fmt.Println(person{name: "Alice", age: 30}) // can name the field

    fmt.Println(person{name: "Fred"}) // skipped field will be zero value

    fmt.Println(&person{name: "Ann", age: 40}) // use & to point to

    fmt.Println(newPerson("Jon")) // use function for new struct creation

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name) // use a "." to access struct field

    sp := &s
    fmt.Println(sp.age) // use "." on pointer is ok, it auto derefernced

    sp.age = 51
    fmt.Println(sp.age) // structs are mutable

    dog := struct { // can use anonymous struct if we only going to use it once
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}