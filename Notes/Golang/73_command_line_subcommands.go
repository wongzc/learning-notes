//https://gobyexample.com/command-line-subcommands
package main

import (
    "flag"
    "fmt"
    "os"
)
// example git commit -m "soemthing"
// git is main command
// commit is subcommand

func main() {

    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// subcommand created with NewFlagSet
    fooEnable := fooCmd.Bool("enable", false, "enable")
    fooName := fooCmd.String("name", "", "name")
	// define flags fro subcommand

    barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	// another subcommand
    barLevel := barCmd.Int("level", 0, "level")

    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {

    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}

// go build 73_command_line_subcommands.go
// ./73_command_line_subcommands foo -enable -name=joe a1 a2
// ./73_command_line_subcommands bar -level 8 a1
// ./73_command_line_subcommands bar -enable a1 