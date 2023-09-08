package main

import "fmt"

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}

func main() {

	slice := make([]string, 5, 8)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"
	inspectSlice(slice)

	var data []string
	for record := 1; record <= 10; record++ {
		data = append(data, fmt.Sprintf("Rec: %d", record))
	}
	inspectSlice(data)
}

/*
$ go run slices.go
Length[5] Capacity[8]
[0] 0xc00010e000 Apple
[1] 0xc00010e010 Orange
[2] 0xc00010e020 Banana
[3] 0xc00010e030 Grape
[4] 0xc00010e040 Plum
Length[10] Capacity[16]
[0] 0xc000110000 Rec: 1
[1] 0xc000110010 Rec: 2
[2] 0xc000110020 Rec: 3
[3] 0xc000110030 Rec: 4
[4] 0xc000110040 Rec: 5
[5] 0xc000110050 Rec: 6
[6] 0xc000110060 Rec: 7
[7] 0xc000110070 Rec: 8
[8] 0xc000110080 Rec: 9
[9] 0xc000110090 Rec: 10
*/
