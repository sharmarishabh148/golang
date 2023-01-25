package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
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

func SortByCost[T Energy](a []T) {
	slices.SortFunc(a, func(a, b T) bool {
		return a.Cost() < b.Cost() || math.IsNaN(a.Cost()) && !math.IsNaN(b.Cost())
	})
}
func main() {
	solar2k := Solar{"Solar 2000", 4.500}
	solar3k := Solar{"Solar 3000", 4.000}
	fmt.Println()
	ss2 := solarSlice{solar2k, solar3k}
	//type of tt : []main.Solar
	//PrintSlice(ss2)
	// type of tt : main.solarSlice
	//PrintSlice2(ss2)
	fmt.Printf("index: %d\n", slices.Index(ss2, solar2k))
	SortByCost(ss2)
	//
	fmt.Printf("index changed now at: %d\n", slices.Index(ss2, solar2k))

}
