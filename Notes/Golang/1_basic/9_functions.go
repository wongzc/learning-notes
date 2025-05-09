//https://gobyexample.com/functions
package main

import "fmt"

func plus(a int, b int) int {

    return a + b //go requrie explicit return
}

func plusPlus(a, b, c int) int { // multiple consecutive same parameter type, can declare at once
    return a + b + c
}

func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}