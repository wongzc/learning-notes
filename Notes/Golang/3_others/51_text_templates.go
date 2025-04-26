//https://gobyexample.com/text-templates
package main

import (
    "os"
    "text/template"
)

func main() {

    t1 := template.New("t1") // template is a text doc with placeholder
    t1, err := t1.Parse("Value is {{.}}\n") // parse the template. {{.}} is placeholder
						// for data that pass later when execute template
	t1.Execute(os.Stdout, 5)
    if err != nil {
        panic(err)
    }

    t1 = template.Must(t1.Parse("Value: {{.}}\n"))
	// parse with template.Must in case Parse returns error it will panic

    t1.Execute(os.Stdout, "some text") // the {{.}} will be replaced after execute
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })

    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

    t2 := Create("t2", "Name: {{.Name}}\n") // can use .field to access data in struct

    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})

    t2.Execute(os.Stdout, map[string]string{
        "Name": "Mickey Mouse",
    })

    t3 := Create("t3",
        "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")

    t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
            "Go",
            "Rust",
            "C++",
            "C#",
        })
}