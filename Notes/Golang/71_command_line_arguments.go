//https://gobyexample.com/command-line-arguments
package main

import (
    "fmt"
    "os"
)
// argument for parameterize go program
func main() {

    argsWithProg := os.Args
	// provide access to raw CL arguments
    argsWithoutProg := os.Args[1:]
	// os.Args[0] is the path to program
	// the other are argument

    arg := os.Args[3] // get individual argument with  indexing

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}

// go build 71_command_line_arguments.go
// run in CLI with something like
// ./71_command_line_arguments a b c d