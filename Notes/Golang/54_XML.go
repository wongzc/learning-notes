//https://gobyexample.com/xml
package main

import (
    "encoding/xml"
    "fmt"
)

type Plant struct {
    XMLName xml.Name `xml:"plant"`
	// XMLName is special field ot indicate XML element name
	// if omitted, the elemnt name will be stuct name <Plant/>, not <plant/>
    Id      int      `xml:"id,attr"`
	//id,attr means Id field is XML attribute not nested element
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)
}

func main() {
    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

    out, _ := xml.MarshalIndent(coffee, " ", "  ")
	// xml.MarshallIndent to print out XML in more readable form
    fmt.Println(string(out))
	fmt.Println("----------")
    fmt.Println(xml.Header + string(out))
	// append a generic XML header
	fmt.Println("----------")

    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
		// Unmarshal to parse XML into a data structure
		// will have error if XML malformed or cannot map to Plant
        panic(err)
    }
    fmt.Println(p)
	fmt.Println("----------")

    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`// tell encoder to nest plant under <parent><child>
    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
	fmt.Println("----------")

	tryy := &Plant{Id: 999, Name: "hahaha", Origin: []string{"A","b","C"}}
	type MyNest struct {
		Huh []*Plant `xml:"a>hah>k"` // anything can be use for "Huh", just need cap
	}
	trynest := &MyNest{}
	trynest.Huh = []*Plant{tryy}

	out, _ = xml.MarshalIndent(trynest, " ", "  ")
    fmt.Println(string(out))
	fmt.Println("----------")
}