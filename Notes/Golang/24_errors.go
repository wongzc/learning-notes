//https://gobyexample.com/errors

package main
// idiomatic to communicate errors via explicit seperate return value in go
import (
    "errors"
    "fmt"
)

func f(arg int) (int, error) { // by convention, put error as last return value
    if arg == 42 {

        return -1, errors.New("can't work with 42") // errors.New construct a basic error with message
    }

    return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available") // sentinel error: predeclared variable for error
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {

        return fmt.Errorf("making tea: %w", ErrPower) // wrap error with higher-level error to add context
    }
    return nil
}

type MyError struct {
    Code int
}

func (e MyError) Error() string {
    return fmt.Sprintf("error with code %d", e.Code)
}

func main() {
    for _, i := range []int{7, 42} {

        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }

    for i := range 5 {
        if err := makeTea(i); err != nil { // use inline error to check

            if errors.Is(err, ErrOutOfTea) {// can use errors.is to check for error 
				fmt.Printf("error: %s\n", err)
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {// can use errors.is to check value for high lv error also
				fmt.Printf("error: %s\n", err)
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }

        fmt.Println("Tea is ready!")
    }



	err := fmt.Errorf("wrapped: %w", MyError{Code: 404})
	var myErr MyError
	if errors.As(err, &myErr) {// use erroes.As to check type
		fmt.Println("Got MyError with code", myErr.Code)
	}
}