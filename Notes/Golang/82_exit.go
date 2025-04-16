//https://gobyexample.com/exit
package main

import (
    "fmt"
    "os"
)
// to exit wiht a given status
func main() {

    defer fmt.Println("!")

    os.Exit(3)// when exit, defer will not be called!!!
}