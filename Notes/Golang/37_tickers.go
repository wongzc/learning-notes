//https://gobyexample.com/tickers
package main

import (
    "fmt"
    "time"
)
// timer: do something at the future
// ticker: do soemthing repeatly at regular intervals
func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
	// tikcer also using channel to send value
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop() // ticker can be stopped like timer
	// once stopped, no more value receive on its channel
    done <- true
    fmt.Println("Ticker stopped")
}