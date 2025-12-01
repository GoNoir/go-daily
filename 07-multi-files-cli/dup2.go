// Dup2 is reading from a file or command line, then count all the lines which are identical and print back the list of lines and the number of occurences.
// In the absence of the specified file, Dub2 will instead give option to input lines directly into CLI and will do the same computation for them
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//counts is just a map with key bing lines and value being number of occurences
	counts := make(map[string]int)
	//now, we need to actually create a new slice containing all the files given
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// we are iterating through the slice of files
		for _, arg := range files {
			// f is the contect of the file (arg)
			f, err := os.Open(arg)
			if err != nil {
				// os.Stderr allows to convieniently write err to a file
				fmt.Fprintf(os.Stderr, "Dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// now, lest check counts and print duplicates
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Dup2: %d:\t%s\n", n, line)
		}
	}
}

// Actual count lines funcion data is with a streaming buffer for a better performance
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// We are ignring unlikely errors for brevity sake
}
