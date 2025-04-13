// https://gobyexample.com/variables
package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2 // var declare multiple variable
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int // declare only, will be zero value
    fmt.Println(e)

    f := "apple" // walrus to declare and initialize
    fmt.Println(f)
}