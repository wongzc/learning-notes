//https://gobyexample.com/strings-and-runes
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
// in go, concept of character is rune
// it is integer that repesent Unicode code point
    const s = "สวัสดี"

    fmt.Println("Len:", len(s)) //  this actually give length of raw byte

    for i := 0; i < len(s); i++ { 
        fmt.Printf("%x ", s[i]) // this will print the hex values that constitute the code point in s
    }
    fmt.Println()

    fmt.Println("Rune count:", utf8.RuneCountInString(s)) // need to use unicode/utf8 package to count number of rune

    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
}

func examineRune(r rune) {

    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}