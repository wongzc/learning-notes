//https://gobyexample.com/recover
package main

import "fmt"
// to recover from panic
// usecase: server close connection of client with critical error, instead of crash
func mayPanic() {
    panic("a problem")
}

func main() {

    defer func() { // recover must be called in defer
		// when enclosing function panic, this will be activated
        if r := recover(); r != nil {

            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    fmt.Println("After mayPanic()") // this will not run as alr panic!
}