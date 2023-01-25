// Sample program to show how maps are not safe for concurrent use by default.
// The runtime will detect concurrent writes and panic.
package main

import (
	"fmt"
	"sync"
)

// scores holds values incremented by multiple goroutines.
var scores = make(map[string]int)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final scores:", scores)
}

/*
$ go run five.go
fatal error: concurrent map writes

goroutine 19 [running]:
runtime.throw({0x5ecd0d, 0x15})
        C:/Program Files (x86)/Go/src/runtime/panic.go:1198 +0x64 fp=0x1182ef88 sp=0x1182ef74 pc=0x581564
runtime.mapassign_faststr(0x5df160, 0x1188a100, {0x5e94b4, 0x1})
        C:/Program Files (x86)/Go/src/runtime/map_faststr.go:211 +0x36e fp=0x1182efcc sp=0x1182ef88 pc=0x55e4de
main.main.func1(0x118a2070)
        C:/Users/sharmari/go/src/github.com/rishabh.sharma/Golang/language/dataraces/example5/five.go:19 +0x47 fp=0x1182efe8 sp=0x1182efcc pc=0x5d5a77
runtime.goexit()
        C:/Program Files (x86)/Go/src/runtime/asm_386.s:1319 +0x1 fp=0x1182efec sp=0x1182efe8 pc=0x5a8321
created by main.main
        C:/Users/sharmari/go/src/github.com/rishabh.sharma/Golang/language/dataraces/example5/five.go:17 +0x5d

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x118a2078)
        C:/Program Files (x86)/Go/src/runtime/sema.go:56 +0x36
sync.(*WaitGroup).Wait(0x118a2070)
        C:/Program Files (x86)/Go/src/sync/waitgroup.go:130 +0x87
main.main()
        C:/Users/sharmari/go/src/github.com/rishabh.sharma/Golang/language/dataraces/example5/five.go:31 +0x87
exit status 2

sharmari@IN-PF1W6G3T MINGW64 ~/go/src/github.com/rishabh.sharma/Golang/language/dataraces/example5
$
*/

/*
$ go run five.go
Final scores: map[A:1000 B:1000]
*/
