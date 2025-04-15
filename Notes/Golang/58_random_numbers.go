//https://gobyexample.com/random-numbers
package main

import (
    "fmt"
    "math/rand/v2"
)

func main() {

    fmt.Print(rand.IntN(100), ",")
    fmt.Print(rand.IntN(100))// random num, 0<=n<100
    fmt.Println()

    fmt.Println(rand.Float64())

    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    s2 := rand.NewPCG(42, 1024) //create a seed, need 2 uint numbers
    r2 := rand.New(s2) // create the random number
    fmt.Print(r2.IntN(100), ",")
    fmt.Print(r2.IntN(100))
    fmt.Println()

    s3 := rand.NewPCG(42, 1024)
    r3 := rand.New(s3)
    fmt.Print(r3.IntN(100), ",")
    fmt.Print(r3.IntN(100))
    fmt.Println()
}