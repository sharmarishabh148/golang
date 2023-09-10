package main

import (
	"fmt"
	"os"
)

func writeToFile(m string, f string) {
	bytes := []byte(m)
	os.WriteFile(f, bytes, 0644)
	fmt.Println("writing data into a file")
}

func readFromFile(f string) {
	fmt.Println("reading data from a file")
	content, _ := os.ReadFile(f)
	fmt.Println(string(content))
}
func main() {

	fmt.Println("writing data into a file")
	filename := "9sept.txt"
	writeToFile("welcome to Go", filename)
	readFromFile(filename)

}
