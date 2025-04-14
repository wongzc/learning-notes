//https://gobyexample.com/if-else
package main

import "fmt"

func main() {

    if 7%2 == 0 { // if else
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 { // if without else
        fmt.Println("8 is divisible by 4")
    }

    if 8%2 == 0 || 7%2 == 0 { // can use operator like || or &&
        fmt.Println("either 8 or 7 are even")
    }

    if num := 9; num < 0 {  // statement can preced conditionals, variable declare in statement are available in branches
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}