//https://gobyexample.com/multiple-return-values

package main

import "fmt"

func vals() (int, int) { // (int,int) means return 2 int
    return 3, 7
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}