//https://gobyexample.com/range-over-iterators
package main

import (
    "fmt"
    "iter"
    "slices"
)

type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) All() iter.Seq[T] { // this function return an iterator of type T, which is func(func(T))
    return func(yield func(T) bool) {

        for e := lst.head; e != nil; e = e.next {
            if !yield(e.val) {
                return
            }
        }
    }
}
// iter.Seq[T] == type Seq[T any] func(yield func(T) bool)

func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 1, 1

        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}

func main() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)

    for e := range lst.All() { // use iterator to range over
        fmt.Println(e)
    }

    all := slices.Collect(lst.All()) // Collect is a slices function that convert iterator into slice
    fmt.Println("all:", all)

    for n := range genFib() {

        if n >= 10 { // once hits break or early return, yield function in iterator will return false
            break
        }
        fmt.Println(n)
    }
}