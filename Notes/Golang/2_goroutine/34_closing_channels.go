//https://gobyexample.com/closing-channels
package main

import "fmt"
// closed it so no more value can be sent on it
// use for communicate completion to receiver

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for { // run for untill return
            j, more := <-jobs // receiving  from jobs channel
			// j is the alue from channel
			// more=false, if channel closed
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ { // send 3 jobs over
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs) // then it was closed after sent jobs
    fmt.Println("sent all jobs")

    <-done // await worker to synchronize

    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)
}