//https://gobyexample.com/file-paths
package main

import (
    "fmt"
    "path/filepath" // the filepath package
    "strings"
)

func main() {

    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	// always use join to create path!
	// it remove superfluous separators and directory change
	
	// using dir or base to get filename or dir
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

	// can split file name and dir
	a,b := filepath.Split(p)
	fmt.Println("a: ",a)
	fmt.Println("b: ",b)

	// check if file is abs
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

    ext := filepath.Ext(filename) // split extension out
    fmt.Println(ext)

    fmt.Println(strings.TrimSuffix(filename, ext))

    rel, err := filepath.Rel("a/b", "a/b/t/file") // find relative path btw 2
    if err != nil {
        panic(err) // return error if cannot made relative
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}