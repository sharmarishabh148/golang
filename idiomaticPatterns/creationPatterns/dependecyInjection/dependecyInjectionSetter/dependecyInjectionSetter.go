package main

import (
	"fmt"
	"io"
	"os"
)

// Creational: Dependency Injection (Setter) pattern
func main() {
	s := NewMyService()
	s.SetWriter(os.Stderr)
	s.WriteHello("world")
}

type MyService struct {
	writer io.Writer
}

func NewMyService() MyService {
	return MyService{}
}
func (s *MyService) SetWriter(w io.Writer) {
	s.writer = w
}

func (s *MyService) WriteHello(m string) {
	fmt.Fprintf(s.writer, "Hello %s\n", m)
}

/*
Constructor - enforce valid dependencies
Setter - allow creating/changing whenever
Which one? Generally - Constructor.
Can use both if required.
*/
