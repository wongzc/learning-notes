//https://gobyexample.com/base64-encoding
package main

import (
    b64 "encoding/base64"
    "fmt"
)

func main() {

    data := "abc123!?$*&()'-=@~"
	// 2 version of b64
	// standard: 	+,	/,	use padding (=),need escape for url
	// url safe: 	-,	_,	no padding,		urlsafe
    sEnc := b64.StdEncoding.EncodeToString([]byte(data)) // standard encoder
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    uEnc := b64.URLEncoding.EncodeToString([]byte(data)) // URL compatible encoder
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}