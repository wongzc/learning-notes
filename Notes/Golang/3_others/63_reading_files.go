//https://gobyexample.com/reading-files
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func check(e error) { // helper to streamline error checks
    if e != nil {
        panic(e)
    }
}

func main() {

    dat, err := os.ReadFile("/tmp/dat") // read entire into memory
    check(err)
    fmt.Print(string(dat))

    f, err := os.Open("/tmp/dat") // open the file as a f
    check(err)

    b1 := make([]byte, 5) // read first 5 byte from f
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    o2, err := f.Seek(6, io.SeekStart)
	// move file pointer to 6 byte
	// io.SeekStart: 6 byte from there
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    _, err = f.Seek(2, io.SeekCurrent)
	//io.SeekCurrent: from current cursor position
    check(err)

    _, err = f.Seek(-4, io.SeekEnd)
    check(err)

    o3, err := f.Seek(6, io.SeekStart)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
	// keep calling read untill getting minimum number of byte
	// fixed-size header
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    _, err = f.Seek(0, io.SeekStart)
	// use for rewind
    check(err)

    r4 := bufio.NewReader(f)
	// efficeint for small read like char by char, line by line
	// buffio read a large chunk at once
	// lesser syscall or disk/network access
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close() // close file
}