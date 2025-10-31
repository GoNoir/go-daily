//A simple little program to practice pointers. 
//swapValues takes two pointers, dereferencing and swat the values around. Testing by print.

package main

import "fmt"

func swapValues(p1, p2 *int) {
	value1 := *p1
	value2 := *p2
	fmt.Println(value1, value2)

	*p1 = value2
	*p2 = value1

	fmt.Println(p1, p2)
	fmt.Println(*p1, *p2)
}

func main() {
	x := 5
	y := 10
	xpnt := &x
	ypnt := &y

swapValues(xpnt, ypnt)

}

