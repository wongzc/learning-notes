// https://gobyexample.com/maps
package main

import (
    "fmt"
    "maps"
)

func main() {
// maps is like dict in python
	var mvar map[string]string // create map (var)
	// but this cant be assign value with mvar["abc"]="cdf", as no memory allocated, it will panic
	// can only use to reassign later, if xxx{ mvar = make(map[string]string)}
	fmt.Println(mvar)
    m := make(map[string]int) // create map using make(map[key type]value type)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m) // print map show all kv pair in it

    v1 := m["k1"] //get value from dict
    fmt.Println("v1:", v1)

    v3 := m["k3"] // if key not eixts, return zero value of value type, 0; ""; nil ( for slice, pointer, struct,...)
    fmt.Println("v3:", v3)

    fmt.Println("len:", len(m)) // number of kv pair in map

    delete(m, "k2") // delete a kv pair
    fmt.Println("map:", m)

    clear(m) // remove all kv pair
    fmt.Println("map:", m)

    _, prs := m["k2"] // optional 2nd retuen value can tell is the key not exist, or value is really 0
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2} // initialize a new map with value
    fmt.Println("map:", n)

    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) { // compare map
        fmt.Println("n == n2")
    }
}