package main

import "fmt"

var kinetecoPrint string = "kineteco Deal:"

type Solar struct {
	Name  string
	Netto float64
}

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

func PrintSlice[T any](tt []T) {
	for i, t := range tt {
		fmt.Printf("%d: %s\n", i, PrintGeneric[T](t))
	}
}

func main() {
	solar2k := Solar{"Solar 2000", 4.500}
	solar3k := Solar{"Solar 3000", 4.000}
	windwest := Wind{"Wind West", 3.950}

	fmt.Println(solar2k.Print())
	fmt.Println(solar3k.Print())
	fmt.Println(windwest.Print())

	fmt.Println("Printing Solar2k using generics\n", PrintGeneric[Solar](solar2k))
	ss := []Solar{solar2k, solar3k}

	fmt.Println("Printing Solar arrays using generics")
	PrintSlice[Solar](ss)

	fmt.Println("Printing Wind arrays using generics")
	PrintSlice[Wind]([]Wind{windwest, windwest})
}
