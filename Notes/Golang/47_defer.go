//https://gobyexample.com/defer
package main

import (
    "fmt"
    "os"
)
// ensure a function call is performed later in a program execution
// usually for cleanup
// similar to "finally"
func main() {

    f := createFile("/tmp/defer.txt")
    defer closeFile(f) // will be execute at the end of exclosing function, which is main here
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        panic(err)
    }
}