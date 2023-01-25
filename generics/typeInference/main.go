package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

var kinetecoPrint string = "kineteco Deal:"

type Number interface {
	constraints.Float | constraints.Integer
}

func Cost[T Number](usage, netto T) T {
	cost := usage * netto
	return cost
}

type Solar struct {
	Name  string
	Netto float64
}

type solarSlice []Solar
type Wind struct {
	Name  string
	Netto float64
}

type Energy interface {
	Solar | Wind
	Cost() float64
}

func (s Solar) Cost() float64 {
	return s.Netto * 0.4
}

func (w Wind) Cost() float64 {
	return w.Netto * 0.4
}

func (s *Solar) Print() string {
	return fmt.Sprintf("%s-%v\n", kinetecoPrint, *s)
}

func (w *Wind) Print() string {
	return fmt.Sprintf("%s-%v\n", kinetecoPrint, *w)
}

func PrintGeneric[T any](t T) string {
	return fmt.Sprintf("%s-%v\n", kinetecoPrint, t)
}

func PrintSlice2[T Energy, S ~[]T](tt S) {
	fmt.Printf("type of tt : %T\n", tt)
	for i, t := range tt {
		fmt.Printf("%d: %s\n", i, PrintGeneric[T](t))
	}
}

func PrintSlice[T any](tt []T) {
	fmt.Printf("type of tt : %T\n", tt)
	for i, t := range tt {
		fmt.Printf("%d: %s\n", i, PrintGeneric[T](t))
	}
}

func main() {
	solar2k := Solar{"Solar 2000", 4.500}

	fmt.Println("Printing Solar2k using generics with Type Inference\n",
		PrintGeneric(solar2k))
	//Notice type is missing. if compiler able to deduce , success
	//Otherwise compiler error

	fmt.Printf("type Inference example :", Cost(10, solar2k.Netto))
	//fmt.Printf("type Inference example :", Cost(0.45, 10))
	fmt.Printf("type Inference example :", Cost[float64](0.45, 10))

	fmt.Println()
	ss2 := solarSlice{solar2k, solar2k}
	//type of tt : []main.Solar
	PrintSlice(ss2)
	// type of tt : main.solarSlice
	PrintSlice2(ss2)
}
