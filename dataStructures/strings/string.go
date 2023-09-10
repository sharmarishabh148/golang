package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var zeroValueString string
	fmt.Println("----demo string concatenation----")
	fmt.Println(zeroValueString)
	zeroValueString = "Hello world"

	//Concatenation
	str1 := "hello"
	str2 := "world"
	str3 := str1 + str2
	fmt.Println(str3)

	// Formatting using fmt.sprintf
	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)

	//string to numeric
	fmt.Println("----demo string to Numeric----")
	str_val1 := "5"
	str_val2 := "2.8790"

	int_val1, err := strconv.ParseInt(str_val1, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(int_val1)
	}

	float_val1, err := strconv.ParseFloat(str_val2, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(float_val1)
	}

	fmt.Println("----demo Numeric to string----")
	str_val1 = fmt.Sprintf("%d", int_val1)
	fmt.Println(str_val1)
	str_val2 = fmt.Sprintf("%f", float_val1)
	fmt.Println(str_val2)

	// Sring parser
	fmt.Println("----demo Numeric to string----")
	data := "Berlin Amsterdam London Tokyo"
	fmt.Println(data)
	cities := strings.SplitAfter(data, " ")
	fmt.Println(cities)
	//conersion
	fmt.Println(strings.ToUpper(data))
}
