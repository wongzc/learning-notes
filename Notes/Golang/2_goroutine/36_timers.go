//https://gobyexample.com/timers

package main

import (
    "fmt"
    "time"
)
// make go code to run at some point in the future
// or repeatly in some interval
func main() {

    timer1 := time.NewTimer(2 * time.Second)
	// represent single event in the future
	// tell timer to wait for how long
	// it then provide a channel that will be notified at that time

    <-timer1.C // block until timer done
	// C is the field of time.Timer
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop() // timer cancelled
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    time.Sleep(2 * time.Second)
}