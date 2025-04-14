// https://gobyexample.com/recursion

package main

import "fmt"

func fact(n int) int { // recursion in untill n==0
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))

    var fib func(n int) int // if recusion with anonymous function, need declare with var!

    fib = func(n int) int {
        if n < 2 {
            return n
        }

        return fib(n-1) + fib(n-2) // declared fib, so go know which function to call with fib
    }

    fmt.Println(fib(7))
}