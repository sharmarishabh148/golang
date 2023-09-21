package main

import (
	"fmt"
	"strings"
)

/*
Given two strings, check whether two strings are isomorphic to each other or not.
  Two strings are isomorphic if a fixed mapping exists from the characters of one
  string to the characters of the other string. For example, if there are two
  instances of the character "a"  in the first string, both these instances
  should be converted to another character (which could also remain the same
  character if "a" is mapped to itself) in the second string. This converted
  character should remain the same in both positions of the second string since
  there is a fixed mapping from the character "a" in the first string to the
  converted character in the second string.
*/

func isIsomorphic(string1 string, string2 string) bool {
	mapStr1Str2 := make(map[string]string)
	mapStr2Str1 := make(map[string]string)

	for index := range string1 {
		char1 := string(string1[index])
		char2 := string(string2[index])
		// Returning false if char1 in string1 exists in hashmap and the
		// char1 has different mapping in hashmap
		if _, ok := mapStr1Str2[char1]; ok && mapStr1Str2[char1] != char2 {
			return false
		}

		// Returning false if char2 in string2 exists in hashmap and
		// char2 has different mapping in hashmap
		if _, ok := mapStr2Str1[char2]; ok && mapStr2Str1[char2] != char1 {
			return false
		}

		// Mapping of char of one string to another and vice versa
		mapStr1Str2[char1] = char2
		mapStr2Str1[char2] = char1
	}
	return true
}

// Driver code
func main() {
	A := []string{"egg", "foo", "paper", "badc", "aaeaa"}
	B := []string{"all", "bar", "title", "baba", "uuxyy"}

	for index, value := range A {
		fmt.Printf("%d.\tString 1: \"%s\"\n", index+1, value)
		fmt.Printf("\tString 2: \"%s\"\n", B[index])
		fmt.Printf("\n\tIsomorphic String ? %t\n", isIsomorphic(value, B[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
