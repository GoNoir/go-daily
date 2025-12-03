//A simple little program to practice pointers. 
//swapValues takes two pointers, dereferencing and swat the values around. Testing by print.

package main

import "fmt"

//takes two pointers
func swapValues(a, b *int) {
 
	*a, *b = *b, *a//idiomatic go swap

}

func main() {
	x := 5
	y := 10

	fmt.Println("Before swap:", x, y)

	//swapping
	swapValues(&x, &y)

	fmt.Println("After swap:", x, y)

}

