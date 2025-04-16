//https://gobyexample.com/environment-variables

package main

import (
    "fmt"
    "os"
    "strings"
)
// universal way to convery configurations to unix programs
func main() {

    os.Setenv("FOO", "1")
	// set environment variable
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))
	// get the env var

    fmt.Println()
    for _, e := range os.Environ() {
		// os.Environ can list out all env var
		// returned in key=value form, need to split by "="
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}