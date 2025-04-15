//https://gobyexample.com/time-formatting-parsing
package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    t := time.Now()
	p(t)
    p(t.Format(time.RFC3339)) // format time to RFC3339

    t1, e := time.Parse( // parse time using format
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    p(t.Format("3:04PM")) // use example based to parse! 
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)
	// but example based time parsing must use the magic time
	// year: 2006
	// month: 01
	// day: 02
	// hour: 15 or 03
	// minute: 04
	// second: 05
	// PM: PM

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}