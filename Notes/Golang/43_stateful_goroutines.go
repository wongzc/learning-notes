//https://gobyexample.com/stateful-goroutines
package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)
// instead of using mutex lock,
// can also use built-in synch of goroutines and channels for similar result
type readOp struct { // the state will be own by single goroutine
	// to ensure data not corrupted with concurrent access
	// other goroutine will send messages to this goroutine to read/ write
    key  int
    resp chan int
} //encapsultae read request
type writeOp struct {
    key  int
    val  int
    resp chan bool
} //encapsultae write request

func main() {

    var readOps uint64
    var writeOps uint64

    reads := make(chan readOp) // used by other go routine to issue read and write request
    writes := make(chan writeOp)

    go func() { // goroutine that own the state, the stateful goroutine
        var state = make(map[int]int)
        for {
            select { // repeatly select on read or write channel to respond to request
            case read := <-reads:
                read.resp <- state[read.key] // send on response channel resp to indicate success
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 100; r++ { // creating 100 read
        go func() {
            for {
                read := readOp{ // each read constructing a readOp
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read // send readop over to read channel
                <-read.resp// block untill receied feedback
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}