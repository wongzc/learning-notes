//https://gobyexample.com/temporary-files-and-directories
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    f, err := os.CreateTemp("", "sample")
	// "" means create file in default location
    check(err)

    fmt.Println("Temp file name:", f.Name())

    defer os.Remove(f.Name())
	// os likely will clean up itself, but still good for us to clean

    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)

    dname, err := os.MkdirTemp("", "sampledir")
	// temporary dir if many temp file needed
    check(err)
    fmt.Println("Temp dir name:", dname)

    defer os.RemoveAll(dname)

    fname := filepath.Join(dname, "file1")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}