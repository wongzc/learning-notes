//https://gobyexample.com/non-blocking-channel-operations
package main

import "fmt"
// send and receie in channel are blocking
// use select + default to implement non-blocking send, receive, multi-way receive
func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    case messages <- msg: // we cant send, as channel not buffered, and no receiver!
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

	
	messages2 := make(chan string,1)
    select {
    case messages2 <- msg: // can send, we have buffer
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages: // attempting multi-way non-blocking select
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}