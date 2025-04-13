// https://gobyexample.com/arrays
package main

import "fmt"

func main() {
//array is typed by element type and number of element
    var a [5]int //initialize array with size only (var)
    fmt.Println("emp:", a)

    a[4] = 100 //set value via index
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5} //initialize with size+value (walrus)
    fmt.Println("dcl:", b)

    b2 := [...]int{1, 2, 3, 4, 5} // initialize with value (walrus) size will be based on values lisy
    fmt.Println("dcl:", b2)

    b3 := [...]int{100, 3: 400, 500} // initialize with value (walrus), 3:400 means place 400 at index 3, value after 400 will be 4, 5, ...
    // those without value will be 0
    fmt.Println("idx:", b3)

    var twoD [2][3]int // initialize 2d array with size only (var)
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    twoD2 := [2][3]int{ // initialize 2d array with size & value (walrus)
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD2)

    twoD3 := [...][3]int{ // initialize 2d array with value (walrus)
        {1, 2, 3}, // go only able ti infer length of outer, so need to provide inner which is 3
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD3)

    var dp [3+1][4] bool
    fmt.Println(dp)
}