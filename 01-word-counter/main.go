/*
   What it does: Counts frequency of each word in a text you type
   Takes string and prints map of words and occurences
   Import fmt for printing and strings library for Split function
   Only simplistic takes on user input 
   Room to refactor and normalize and better whitespace and punctuation handling 
   Read a line of text from user
   Split into words, count occurrences
   Print each word and its count
*/

package main

import (
	"fmt"
	"strings"
)

func wordCounter(s string) map[string]int {

	wordsCounts := make(map[string]int)
	splittedString := strings.Split(s, " ")

	for _, word := range splittedString {

		_, exists := wordsCounts[word]
		if !exists {
			wordsCounts[word] = 1
		} else {
			wordsCounts[word] += 1
		}
	}
	return wordsCounts
}

func main() {
	testUserString := "This is a Hello world test. Hello world!"
	result := wordCounter(testUserString)
	fmt.Println(result)
}
