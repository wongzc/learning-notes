//https://gobyexample.com/methods
package main
// go support methods defined on struct type
import "fmt"

type rect struct {
    width, height int
}
type MyInt int
// not just struct, any named type can have method
func (m MyInt) Double() int {
    return int(m) * 2
}
func (r *rect) area() int { // func (y *struct) name() xx {}, this is creating method for func
    return r.width * r.height
} // pointer receiver, use when method modify struct

func (r rect) perim() int {
    return 2*r.width + 2*r.height
} // value receiver, use when method doent modify struct

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r // if we pass pointer, go auto convert it
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}