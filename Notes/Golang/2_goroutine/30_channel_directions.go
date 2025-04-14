// https://gobyexample.com/channel-directions
package main

import "fmt"
// when passing channel as parameter to function
// can specify if channel meant to send or receive
// make it more type safe
func ping(pings chan<- string, msg string) { // only accept channel for send
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // channel for send & receive
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}