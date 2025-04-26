//https://gobyexample.com/embed-directive
package main

import (
    "embed" 
	// a compiler directive to include abritary file/ folder in go binary
	// at build time
	// so no need to manually provide the txt or html or json etc
)


// below go:embed.... must appear aboe the variable!
//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

// embed multiple files or folder with wildcard
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

    print(fileString)
    print(string(fileByte))

	// then can read file
    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))

    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
}