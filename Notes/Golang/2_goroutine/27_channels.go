//https://gobyexample.com/channels

package main

import "fmt"
// channel are pipe that connect concurrent goroutines
// goroutine can send value into channel
// goroutine can receie value from channel

func main() {

    messages := make(chan string) // create channel, channel type decide by value they carry

    go func() { messages <- "ping" }() // channel <- syntax means send to channel

    msg := <-messages // <-channel means received from channel
    fmt.Println(msg)
	// by default send and receive block untill both sender and receiver ready
	// means at that moment 1 want to send , another want to receive
}