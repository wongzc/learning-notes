//https://gobyexample.com/command-line-flags
package main

import (
    "flag"
    "fmt"
)
// flags like the "-l" in "wc -l"
func main() {

    wordPtr := flag.String("word", "foo", "a string")
	// flag declaration
	// flag: word, default value: foo, description: a string

    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")
	// number and fork flag

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
	// declare an option using exiting var declared
	// need to pass in pointer

    flag.Parse()
	// call flag.Parse ti execute CLI

    fmt.Println("word:", *wordPtr)
	// need dereference the pointer to get actual alue
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
// go build 72_command_line_flags.go
// ./72_command_line_flags -word=opt -numb=7 -fork -svar=flag
// ./72_command_line_flags -word=opt a1 a2 a3
// ./72_command_line_flags -h
//		use this for help
// ./72_command_line_flags -wat
//		if provided flag that not specified, will have error