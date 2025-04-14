//https://gobyexample.com/select
package main

import (
    "fmt"
    "time"
)
// use select to wait on multiple channel operations
// people usually combine channel with select
func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
		time.Sleep(1 * time.Second)
		c2 <- "haha"
    }()

    for i := 0; i < 3; i++ {
        select { // using select to await both operations
			// select can wait simultaneously
			// we cant just use for loop, as channel is blocking. i.e., msg1 will block the check for msg2
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}