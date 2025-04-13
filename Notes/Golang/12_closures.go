//https://gobyexample.com/closures

package main

import "fmt"
// go support anonymous function
// closure: capture variable from surrounding scope, even after outer function finished execute
// use for
//		track state without global
//		crete function factory
//		custom callback: custom behavior that use function as argument

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())

	funcs := []func(){} // when use closure in loop, loop variable is share
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() {
            fmt.Println(i) // to avoid getting 3 3 3 we need to have a new variable newI :=i, and fmt.Print(newI)
        })
    }

    for _, f := range funcs {
        f() // we will get 3 3 3 here
    }
}