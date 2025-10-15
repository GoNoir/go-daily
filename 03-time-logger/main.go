/*
Simple Time Logger
Focus:
Structs, time.Time, duration calculation

Flow:
Start a task (records time)
Stop task
Show how long it took

Strategy:
Create a custom Task type struct which will group time related data
Hypothesis - using methods with a pointer receiver to change instance data(not just a copy as in case of value receiver method)
*/

package main

import (
	//	"time"
	"fmt"
)
/*
type Task struct {
	Name  string
	Start time.Time
	End   time.Time
}

func (t *Task) calcDuration(x, y time.Time ) time.Time {
	return y - x
}
*/

func factorial(x int) int {
	sum := 1
	for i := 2; i < x+1; i++ {
		sum *= i 
	}
	return sum
}

func main() {
	fmt.Println(factorial(5))
}

//methods MIGHT be redundant after all - TBC
