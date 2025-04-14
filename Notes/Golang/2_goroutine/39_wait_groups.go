//https://gobyexample.com/waitgroups
package main

import (
    "fmt"
    "sync"
    "time"
)
// use to wait for multiple goroutine finished

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup
	// waitgroup created to wait for all goroutines launched here to finish
	// if passing waitgroup into fucntions, should be done by pointer

    for i := 1; i <= 5; i++ {
        wg.Add(1) // launch worker and increase waitgroup counter

        go func() { // wrap waorker call in a closure to makesure telling waitgroup worker done
            defer wg.Done() // scheduled to run at the end of function!
            worker(i)
        }()
    }

    wg.Wait() // untill waitgroup count goes to 0

}