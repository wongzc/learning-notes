//https://gobyexample.com/struct-embedding
package main

import "fmt"
// embed stucts and interface to express composition of types
type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct { //container embed base
    base // the struct that we embed in
    str string
}

func main() {

    co := container{ // createing structs with embed
        base: base{ // need to explicit intialize the embedding
            num: 1,
        },
        str: "some name",
    }

    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // can access base field directly

    fmt.Println("also num:", co.base.num) // or using full path to access it

    fmt.Println("describe:", co.describe()) //container can use base method as well

    type describer interface {
        describe() string
    }

    var d describer = co // embeding base into container, so it implement interface
    fmt.Println("describer:", d.describe())
}