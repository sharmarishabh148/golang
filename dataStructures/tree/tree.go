package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) insert(d int) {
	if t != nil {
		//check left side
		if t.Left == nil {
			t.Left = &Tree{nil, d, nil}
		} else {
			if t.Right == nil {
				t.Right = &Tree{nil, d, nil}
			} else {
				if t.Left != nil {
					t.Left.insert(d)
				} else {
					t.Right.insert(d)
				}
			}
		}

	} else {
		t = &Tree{nil, d, nil}
	}
}

// print method for printing a Tree
func print(tree *Tree) {
	if tree != nil {
		fmt.Println(" Value", tree.Value)
		fmt.Printf("Tree Node Left")
		print(tree.Left)
		fmt.Printf("Tree Node Right")
		print(tree.Right)
	} else {
		fmt.Printf("Nil\n")
	}
}
func main() {
	var tree *Tree = &Tree{nil, 1, nil}
	print(tree)
	tree.insert(3)
	print(tree)
	tree.insert(5)
	print(tree)
	tree.Left.insert(7)
	print(tree)
}
