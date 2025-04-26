//https://gobyexample.com/panic
package main

import "os"
//something went unexpectedly wrong
// to fail fast on errors that shouldnt occur 
func main() {

    panic("a problem")

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err) // use panic to abort when a function retuned error
    }
}