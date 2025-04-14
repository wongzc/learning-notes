//https://gobyexample.com/string-formatting
package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
    fmt.Printf("struct1: %v\n", p) //print struct

    fmt.Printf("struct2: %+v\n", p) // print struct with field name

    fmt.Printf("struct3: %#v\n", p) //prints a Go syntax representation of the value
	// the source code snippet that would produce the alue

    fmt.Printf("type: %T\n", p) // print type

    fmt.Printf("bool: %t\n", true)

    fmt.Printf("int: %d\n", 123) // base-10

    fmt.Printf("bin: %b\n", 14) // binary

    fmt.Printf("char: %c\n", 33) // print char of corresponding int

    fmt.Printf("hex: %x\n", 456) // hex encoding

    fmt.Printf("float1: %f\n", 78.9) //float

    fmt.Printf("float2: %e\n", 123400000.0) // scientific notation
    fmt.Printf("float3: %E\n", 123400000.0) // same

    fmt.Printf("str1: %s\n", "\"string\"")

    fmt.Printf("str2: %q\n", "\"string\"") // print string as double quoted

    fmt.Printf("str3: %x\n", "hex this") // string in base-16

    fmt.Printf("pointer: %p\n", &p) //represent pointer

    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}