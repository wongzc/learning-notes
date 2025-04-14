// https://gobyexample.com/rate-limiting
package main

import (
    "fmt"
    "time"
)
// implement reate limiting with goroutines, channels, tickers

func main() {

    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)
	// limiter to receive a value every 200 ms

    for req := range requests {
        <-limiter // blocking, so 1 request every 200 ms
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3) // allow burst up to 3 events

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}