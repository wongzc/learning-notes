//https://gobyexample.com/constants

package main

import (
    "fmt"
    "math"
)

const s string = "constant" // use const to declare constnat

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d)) // numeric constant dont have type until explicit conversion

    fmt.Println(math.Sin(n))
}