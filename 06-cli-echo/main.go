package main

import (
	"os"
	"fmt"
)

func main() {
	var s string
	for _, v := range os.Args[1:] {
		s += v + " "
	}
	fmt.Println(s)

}

