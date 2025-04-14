//https://gobyexample.com/range-over-channels
package main

import "fmt"
// use for range to iterate over values received from channel

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue { // loop untill channel is close
        fmt.Println(elem) // if not closed, it will keep waiting for data
    }
}