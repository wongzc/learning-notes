//https://gobyexample.com/sorting-by-functions
package main

import (
    "cmp"
    "fmt"
    "slices"
)
// custom sort like using length etc
func main() {
    fruits := []string{"peach", "banana", "kiwi"}

    lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
		// implement a comparison function with cmp.Compare
    }

    slices.SortFunc(fruits, lenCmp) // use slice.SortFunc for custom sort
    fmt.Println(fruits)

    type Person struct {
        name string
        age  int
    }

    people := []Person{
        Person{name: "Jax", age: 37},
        Person{name: "TJ", age: 25},
        Person{name: "Alex", age: 72},
    }

    slices.SortFunc(people, // custom sort for non built-in type
        func(a, b Person) int {
            return cmp.Compare(a.age, b.age)
        })
    fmt.Println(people)
}