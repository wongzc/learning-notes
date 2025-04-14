//https://gobyexample.com/worker-pools
package main

import (
    "fmt"
    "time"
)
// implement worker pool using goroutine & channel

func worker(id int, jobs <-chan int, results chan<- int) {// worker
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

    const numJobs = 5
    jobs := make(chan int, numJobs) // channel to send work
    results := make(chan int, numJobs) // channel to get result

    for w := 1; w <= 3; w++ { // create 3 workers, initially bloacked due to no job
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ { // send 5 jobs
        jobs <- j
    }
    close(jobs)// close to indicate completion for worker

    for a := 1; a <= numJobs; a++ {
        <-results // collect all results of work
    }
}