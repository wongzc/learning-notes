//https://gobyexample.com/url-parsing
package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {

    s := "postgres://user:pass@host.com:5432/path?k=v#f"
	// include scheme, authentication info, host, port, path, query params, query fragment

    u, err := url.Parse(s) // parse url
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    fmt.Println(u.Host) // contain both host & port
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery) // to get query parames in k=v format
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}