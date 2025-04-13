//https://gobyexample.com/pointers

package main
// pass references to values and record
import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) { //*int means it is a pointer to int
    *iptr = 0 //dereference pointer from memory address to value, get value from addres
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i) // & syntax give memory address, the pointer
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}