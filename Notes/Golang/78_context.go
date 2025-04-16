//https://gobyexample.com/context
package main

import (
    "fmt"
    "net/http"
    "time"
)
// context: cancel long running goroutine, timeout for external call, pass request scoped value
func hello(w http.ResponseWriter, req *http.Request) {

    ctx := req.Context() // create a context, tied to http request
	// ctx will be auto canceled if client disconnect ot request timeout 
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done(): //ctx.Done() means client disconnected or request cancelled

        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}

//curl localhost:8090/hello
// we can also use to cancel DB queries if it is too long
//		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// then cts.DOne becomes readable
// requets scope value
//		ctx = context.WithValue(ctx, "userID", 123)
