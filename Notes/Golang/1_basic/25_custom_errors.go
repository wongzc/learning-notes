//https://gobyexample.com/custom-errors
package main

import (
    "errors"
    "fmt"
)

type argError struct { // usually put "Error" as suffix for custom error
    arg     int
    message string
}

// using *argError, so only pointer implement the error interface
func (e *argError) Error() string { // adding Error method so that it implement error interface
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"} // create new error and return a pointer to the struct
														// which implement error interface, so actually return error
    }
    return arg + 3, nil
}

func main() {

    _, err := f(42)
    var ae *argError // ae type is pointer to argError type
	// not &argError, as we cant point to address of type
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}