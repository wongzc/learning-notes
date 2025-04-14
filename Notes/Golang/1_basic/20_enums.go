//https://gobyexample.com/enums
package main
// type with fixed number of possible values
// and each with distinct name
import "fmt"

type ServerState int

const (
	// create const, with type and value from 0
    StateIdle ServerState = iota //iota: special identifier, auto increment with each line inside const block
    StateConnected 	// ServerState(1)
    StateError		// ServerState(2)
    StateRetrying	// ServerState(3)
)

var stateName = map[ServerState]string{ // map from server state to string
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

func (ss ServerState) String() string { // String() method, will be auto called in Println(), Printf(),Sprintf()
    return stateName[ss]
}

func main() {
    ns := transition(StateIdle)
    fmt.Println(ns)

    ns2 := transition(ns)
    fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:

        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}