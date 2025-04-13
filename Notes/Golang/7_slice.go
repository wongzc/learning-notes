// https://gobyexample.com/slice
package main

import (
    "fmt"
    "slices"
)

func main() {
// slice is onyl type by element type, not element number
    var s []string // uninitialized slice, equal to nil, length 0
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

    s = make([]string, 3) // create slice with length (make)
	// by defualt cap is length, can give more cap if know going to grow
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))


    s[0] = "a" // slice can set & get value like array
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d") // can append for slice, need to accept returned value
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s) // use copy(new,ori) to copy
    fmt.Println("cpy:", c)

    l := s[2:5] // support slicing
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"} // declare slice
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) { // compare slices
        fmt.Println("t == t2")
    } // slice cannot compare with == ( but can use slice==nil), or use for loop to check
	

	checkEqual:= func (a,b [] string) bool { // using anonymous way to declare
		if len(a)!=len(b){ // xx := func (a,b) x {} but not func xx(a,b) x {}
			return false
		}
		for i:=range a {
			if a[i]!=b[i] {
				return false
			}
		}
		return true
	}

	if checkEqual(t,t2) {
		fmt.Println(("equal!"))
	}

    twoD := make([][]int, 3) // 3 is outer length, need for loop to define for inner length
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}