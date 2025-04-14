//https://gobyexample.com/for

package main

import "fmt"

func main() {

    i := 1 // basic for loop with just condition
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j < 3; j++ {// clasic initial;condition;action
        fmt.Println(j)
    }

    for i := range 3 { // using range
        fmt.Println("range", i)
    }

    for { // run repeatly untill encounter break or return
        fmt.Println("loop")
        break
    }

    for n := range 6 { 
        if n%2 == 0 {// use continue to continue to next iteration
            continue
        }
        fmt.Println(n)
    }
}