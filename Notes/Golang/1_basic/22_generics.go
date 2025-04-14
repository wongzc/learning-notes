//https://gobyexample.com/generics
package main

import "fmt"

// generic function: function work with different type!!
// E comparable: means can be compare using == or != with it
//		int, float64, boolean, pointer. not slice, map function.
//		but array and struct ok if underlying are comparable
// S ~[]E: S is type that behave like slice of E, ~: allow S to be anytype with the underlying type []E
// s S: slice to search
// v E: value to find
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}

type List[T any] struct { //defined a Struct type with head and tail, which point to element[T]
    head, tail *element[T] // T can be any value
}

type element[T any] struct { // defines a Struct of element, with next ( a pointer to element) and a val
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) { // T is from the struct, need to be same as struct
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

// another simple example of generic function
func GetFirst[T any](s []T) T {
    return s[0]
}

func main() {
    var s = []string{"foo", "bar", "zoo"}
    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	var k =[]int{1,9,4,0,5,3,2}
    fmt.Println("index of 0:", SlicesIndex(k, 0))

    _ = SlicesIndex[[]string, string](s, "zoo") // we can specify the input type explicitly, but usually skip

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())

	fmt.Println(GetFirst([]int{1, 2, 3})) // work with []int
	fmt.Println(GetFirst([]string{"a", "b"})) // work with []string
}