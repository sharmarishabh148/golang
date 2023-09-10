package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "世界 means world"
	var buf [utf8.UTFMax]byte

	for i, r := range s {
		rl := utf8.RuneLen(r)
		si := i + rl
		copy(buf[:], s[i:si])
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}

	data := "We♥Go"
	fmt.Println(data)
	//c := '❤'
	fmt.Println("Length: ", len(data))
	fmt.Println("Rune Count: ", utf8.RuneCountInString(data))
}

/*
$ go run slices.go
 0: '世'; codepoint: 0x4e16; encoded bytes: []byte{0xe4, 0xb8, 0x96}
 3: '界'; codepoint: 0x754c; encoded bytes: []byte{0xe7, 0x95, 0x8c}
 6: ' '; codepoint:   0x20; encoded bytes: []byte{0x20}
 7: 'm'; codepoint:   0x6d; encoded bytes: []byte{0x6d}
 8: 'e'; codepoint:   0x65; encoded bytes: []byte{0x65}
 9: 'a'; codepoint:   0x61; encoded bytes: []byte{0x61}
10: 'n'; codepoint:   0x6e; encoded bytes: []byte{0x6e}
11: 's'; codepoint:   0x73; encoded bytes: []byte{0x73}
12: ' '; codepoint:   0x20; encoded bytes: []byte{0x20}
13: 'w'; codepoint:   0x77; encoded bytes: []byte{0x77}
14: 'o'; codepoint:   0x6f; encoded bytes: []byte{0x6f}
15: 'r'; codepoint:   0x72; encoded bytes: []byte{0x72}
16: 'l'; codepoint:   0x6c; encoded bytes: []byte{0x6c}
17: 'd'; codepoint:   0x64; encoded bytes: []byte{0x64}
*/
