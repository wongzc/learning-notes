//https://gobyexample.com/line-filters
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
// similar to grep and sed

func main() {

    scanner := bufio.NewScanner(os.Stdin)
	// use bufio to buffer os.Stdin
	// so can use scanner

    for scanner.Scan() {

        ucl := strings.ToUpper(scanner.Text())

        fmt.Println(ucl)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}