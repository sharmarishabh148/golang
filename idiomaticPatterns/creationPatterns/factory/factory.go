package main

import (
	"fmt"
	"io"
	"os"
)

/*
Initialize a value for an interface,
but let the Factory decide which implementation
to instantiate
*/
type MyWriter struct {
}

func (w MyWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s", p)
	return len(p), nil
}

func NewWriter(kind string) (io.Writer, error) {
	switch kind {
	case "mywriter":
		return MyWriter{}, nil
	case "stderr":
		return os.Stderr, nil
	default:
		return nil, fmt.Errorf("kind invalid: %s", kind)
	}
}
func main() {
	kind := "mywriter"
	if len(os.Args) > 1 && os.Args[1] == "stderr" {
		kind = "stderr"
	}
	//create writer and
	writer, _ := NewWriter(kind)
	writer.Write([]byte("Hello world\n"))

}
