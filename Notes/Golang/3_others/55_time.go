//https://gobyexample.com/time
package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    now := time.Now() // current date time
    p(now)

    then := time.Date( // build a time struct by providing y,m,d, etc
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year()) // can extract from time
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now)) // compare time
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then) // duration between now and then
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff)) // add - to move backward
}