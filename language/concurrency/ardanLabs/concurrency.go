package main

import (
	"fmt"
	"runtime"
)

/*func init() {
	runtime.GOMAXPROCS(1)
}
*/

func main() {
	/*. When passing 0, the number of threads the Go program will be using
	is reported. I must make sure that number matches the number of operating
	system threads I have available in my containerized environmen
	*/
	g := runtime.GOMAXPROCS(0)
	fmt.Println("g :", g)
}
