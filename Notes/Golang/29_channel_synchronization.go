// https://gobyexample.com/channel-synchronization

package main

import (
    "fmt"
    "time"
)
// use channel to synchronize execution across goroutines
// use blocking receive for this
// when waiting for multiple goroutines, can use WaitGroup
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done // block here untill we receive a notification from worker
	// if the above line removed, program will exit before worker started!
}