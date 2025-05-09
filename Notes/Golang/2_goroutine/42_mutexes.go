//https://gobyexample.com/mutexes
package main

import (
    "fmt"
    "sync"
)
// use mutex to safely access more complex data across multiple goroutine
type Container struct {
    mu       sync.Mutex // add a Mutex to synchronize access
	// mutexs must not be copied, so if passing this struct, need to use pointer
    counters map[string]int
}

func (c *Container) inc(name string) {

    c.mu.Lock() // lock the mutex before accesing counters
    defer c.mu.Unlock() //unlock it at the end of function
    c.counters[name]++
}

func main() {
    c := Container{ // zero value mutex is usable, so no need to initialize

        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }

    wg.Add(3) // run 3 goroutines concurrently, all access same counter
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    wg.Wait()
    fmt.Println(c.counters)
}