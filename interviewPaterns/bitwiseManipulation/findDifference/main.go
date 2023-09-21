package main

/*
Given two strings, str1 and str2, find the index of the extra character that is
present in only one of the strings.
*/

import (
	"fmt"
	"strings"
)

func extraCharacterIndex(str1 string, str2 string) int {

	result := 0
	str1Length := len(str1)
	str2Length := len(str2)

	for i := 0; i < str1Length; i++ {
		result = result ^ int(str1[i])
	}

	for i := 0; i < str2Length; i++ {
		result = result ^ int(str2[i])
	}

	if str1Length > str2Length {
		index := strings.Index(str1, fmt.Sprintf("%d", result))
		return index
	} else {
		index := strings.Index(str2, fmt.Sprintf("%d", result))
		return index
	}
}

// Driver code
func main() {
	string1List := []string{"wxyz", "cbda", "jlkmn", "courae", "hello"}
	string2List := []string{"zwxgy", "abc", "klmn", "couearg", "helo"}

	for i := 0; i < len(string1List); i++ {
		fmt.Printf("%d.\tString 1 = %s\n", i+1, string1List[i])
		fmt.Printf("\tString 2 = %s\n", string2List[i])
		fmt.Printf("\n\tThe extra character is as index: %d\n", extraCharacterIndex(string1List[i], string2List[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
