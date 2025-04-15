//https://gobyexample.com/writing-files
package main

import (
    "bufio"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    d1 := []byte("hello\ngo\n")
    err := os.WriteFile("/tmp/dat1", d1, 0644) // dump string into a file
    check(err)

    f, err := os.Create("/tmp/dat2") // or open a file, and write
    check(err)

    defer f.Close() // defer close is good, so it will be closed at the end

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2) // write: only for byte slice
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n") // write string
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f) //buffered write
	// when write, write to buffer
	// only flush to io.writer when buffer is full
	// less syscall!
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush() // nee to flush to ensure all buffered write are done

}