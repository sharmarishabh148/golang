package main

import "fmt"

type sigma struct {
	name        string
	dob         int
	fathersName string
}

type rishabh struct {
	name        string
	dob         int
	fathersName string
}

func main() {
	fmt.Print("Hello World")

	var s sigma
	var r rishabh

	s = sigma(r)
	fmt.Println(s, r)

}
