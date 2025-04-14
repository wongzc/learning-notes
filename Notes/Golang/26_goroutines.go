//https://gobyexample.com/goroutines

package main
// alight weight thread execution
import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct") // usually way of calling

    go f("goroutine") // go routine way

    go func(msg string) { // goroutine can run with anonymous function
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second) // sleep for 1 seconds use 2*time.second for more
    fmt.Println("done")
}