package main

import (
	"fmt"
	"unicode/utf8"
)

//type byte = uint8
// represents ascii chars

//type rune = int32
// represents set of unicode charters in Utf-8 Encoding
// can be of 1to 4 bytes

func main() {
	data := "We♥Go"
	fmt.Println(data)
	//c := '❤'
	fmt.Println("Length: ", len(data))
	fmt.Println("Rune Count: ", utf8.RuneCountInString(data))
}
