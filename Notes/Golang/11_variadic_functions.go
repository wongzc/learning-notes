//https://gobyexample.com/variadic-functions

package main
// variadic function: function can be called with any number of trailing arguments
// like fmt.Println

import "fmt"

func sum(nums ...int) { // the nums is actually equivalent to []int, except when we input to func, we dont use slice!
    fmt.Print(nums, " ") // can call the len of nums
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...) // use ... to spread the element in a slice
}