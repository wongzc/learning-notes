//https://gobyexample.com/interfaces
package main

import (
    "fmt"
    "math"
)
// interface is named collections of method signature
type geometry interface { // defined interface with 2 method
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// to implement interface, just meed to implement all methods in interface
// any named type that implement all method of interface auto satisfied interface
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) { // variable of interface type: we can call methods that are in interface
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func detectCircle(g geometry) {
    if c, ok := g.(circle); ok { 
		// syntax: value, ok := interfaceVar.(ConcreteType)
		// check if an interface variable holding a type
		// example: 
		// 		check if int: v, ok := x.(int) 
		//		check funcion: fn, ok := f.(func(int) int)
        fmt.Println("circle with radius", c.radius)
    }
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)

    detectCircle(r)
    detectCircle(c)
}