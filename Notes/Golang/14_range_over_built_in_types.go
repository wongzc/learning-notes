// https://gobyexample.com/range-over-built-in-types
package main

import "fmt"
// use range to iterate over elements
// for key, value := range map
// for index, value := range list || slice
// for start byte index of rune, rune := range string
// 	rune: single unicode character
//  byte index: go using utf-8, character range from 1-4 byte
// trick for rune: runes:=[]rune(s); fmt.Println(string(runes[1])) to get the complete rune instead of broken 1 !
func main() {

    nums := []int{2, 3, 4} 
    sum := 0
    for _, num := range nums { // range over slice
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums { // range over arrays
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"} // range over map
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs { // range with just the key
        fmt.Println("key:", k)
    }

    for i, c := range "go" { // range over string
        fmt.Println(i, c, c=='g')
    }
}