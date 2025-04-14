//https://gobyexample.com/timeouts
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1) // make a buffered channel
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1" // send in data after sleep 2 sec
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second): // awaits a value to be sent after receive timeout of 1s
	// if operation take >1 sec, will go to here
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}