package main

import (
	"fmt"
	"regexp"
	"strings"
)

// reverseWords is our challenge function
func reverseWords(sentence string) string {
	// remove leading, trailing and multiple spaces
	trimedSentence := trimString(sentence)

	// We need to convert the input strings to lists of characters as strings are immutable in Go
	sentenceBytes := []byte(trimedSentence)

	strLen := len(sentenceBytes)

	// We will first reverse the entire string.
	sentenceBytes = strRev(sentenceBytes, 0, strLen-1)

	// Now all the words are in the desired location, but
	// in reverse order: "Hello World" -> "dlroW olleH".

	start, end := 0, 0

	for start < strLen {

		for end < strLen && sentenceBytes[end] != ' ' {
			end += 1
		}

		// Let's call our helper function to reverse the word in-place.
		strRev(sentenceBytes, start, end-1)
		start = end + 1
		end += 1
	}
	return string(sentenceBytes)
}

func trimString(str string) string {
	// Trim extra spaces at the beginning and end of the string
	str = strings.TrimSpace(str)

	// Replace multiple spaces with a single space
	regex := regexp.MustCompile(`\s+`)
	str = regex.ReplaceAllString(str, " ")

	return str
}

// strRev function that reverses a whole sentence character by character
func strRev(str []byte, startRev int, endRev int) []byte {
	// Starting from the two ends of the list, and moving
	// in towards the centre of the string, swap the characters
	for startRev < endRev {
		temp := str[startRev]       // temp store for swapping
		str[startRev] = str[endRev] // swap step 1
		str[endRev] = temp          // swap step 2

		startRev += 1 // move forwards towards the middle
		endRev -= 1   // move backwards towards the middle
	}
	return str
}

// Driver code
func main() {
	stringToReverse := []string{
		"Hello World", "We love GoLang",
		"The quick brown fox jumped over the lazy dog",
		"Hey", "To be or not to be", "AAAAA", "Hello     World"}

	for i, str := range stringToReverse {
		fmt.Printf("%d.\tActual string:      \"%s\"\n", i+1, str)
		fmt.Printf("\tReversed string:    \"%s\"\n", reverseWords(str))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
