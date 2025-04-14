//https://gobyexample.com/json
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct { 
	// only exported field will be encoded, decoded
	// exported field: start with capital letter
	// lowercase will not be encode/decode
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"` // with json tags to customize json key names
}

func main() {

	// encode basic type
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

	// encode slice to json array
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

	// encode map to json objects
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &response2{ // response 2 with json tag, so customize field key
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))
	

	// to decode
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    var dat map[string]interface{} 
	// need to provide variable for JSON to decode into
	// hold a map of string to abritary value

    if err := json.Unmarshal(byt, &dat); err != nil { // decoding
        panic(err)
    }
    fmt.Println(dat)

    num := dat["num"].(float64) // need to convert to appropriate type to use it
    fmt.Println(num)

    strs := dat["strs"].([]interface{})// need series of conversion to access nested data
    str1 := strs[0].(string)
    fmt.Println(str1)

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res) // decode into custom data type also can
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)

    dec := json.NewDecoder(strings.NewReader(str))
    res1 := response2{}
    dec.Decode(&res1)
    fmt.Println(res1)
}