//https://gobyexample.com/channel-buffering
package main

import "fmt"
// channel are unbuffered by default
// will only accept send if there are corresponding receiver ready
// buffered: accept limited number of values without corresponding receiver
func main() {

    messages := make(chan string, 2) // string channel with buffer=2

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}