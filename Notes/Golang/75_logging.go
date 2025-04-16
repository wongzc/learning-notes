//https://gobyexample.com/logging
package main

import (
    "bytes"
    "fmt"
    "log" // for free form output
    "os"

    "log/slog" // for structured output
)

func main() {

    log.Println("standard logger")

    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    log.Println("with micro")

    log.SetFlags(log.LstdFlags | log.Lshortfile)
	// use flags to configure output format
    log.Println("with file/line")

    mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	// create custom logger and pass around
    mylog.Println("from mylog")

    mylog.SetPrefix("ohmy:")
    mylog.Println("from mylog")

    var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)

    buflog.Println("hello")

    fmt.Print("from buflog:", buf.String())

    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)
    myslog.Info("hi there")

    myslog.Info("hello again", "key", "val", "age", 25)
}