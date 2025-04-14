//https://gobyexample.com/atomic-counters
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)
// using sync/atomic package to manager state
func main() {

    var ops atomic.Uint64 // represent a counter

    var wg sync.WaitGroup // waitgroup to wait for all goroutines finish

    for i := 0; i < 50; i++ { // start 50 goroutine
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ { // each increase counter by 1000

                ops.Add(1)
            }

            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("ops:", ops.Load())
	// use Load to safely read a value while other goroutine updating it
	// but now actually is finished
}