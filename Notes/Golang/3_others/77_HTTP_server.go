//https://gobyexample.com/http-server
package main

import (
    "fmt"
    "net/http"
)
// handler: object implementing http.Handler interface
// use http.HandlerFunc adapter on functions to write handle
// function server as handler take a http.responsewriter and http.request as arg
// http.ResponseWriter: is used to fill http response
func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	// read all HTTP request header and echo
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
	// register handler on server route
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
	// listen and serve
}

//curl localhost:8090/hello