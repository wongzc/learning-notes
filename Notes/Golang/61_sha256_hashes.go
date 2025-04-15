//https://gobyexample.com/sha256-hashes
package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    s := "sha256 this string"

    h := sha256.New()

    h.Write([]byte(s)) // write byte

    bs := h.Sum(nil) // get final hash

    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}