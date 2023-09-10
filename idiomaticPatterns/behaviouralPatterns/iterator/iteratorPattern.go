package main

import (
	"errors"
	"log"
)

/*
Accesses elements sequentially,
without exposing underlying representation.
● Can be used with a Decorator!
● io.Reader & sql/database.Row are exampleseWriter
*/

type Tasks []string
type Iterator struct {
	tasks    []string
	position int
}

var ErrEOF = errors.New("EOF")

func (t *Iterator) Next() (int, string, error) {
	t.position++
	if t.position > len(t.tasks) {
		return t.position, "", ErrEOF
	}
	return t.position, t.tasks[t.position-1], nil
}

func NewTasksIterator() Iterator {
	return Iterator{
		tasks: []string{
			"say hello",
			"say goodbye",
		},
	}
}
func main() {
	ti := NewTasksIterator()
	for {
		i, t, err := ti.Next()
		if err != nil {
			log.Fatalf("Unknown error: %s", err)
		}
		if err == ErrEOF {
			log.Printf("Done")
			break
		}
		log.Printf("Task %d:, %s\n", i, t)
	}
}
