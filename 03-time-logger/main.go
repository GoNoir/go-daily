/*
Simple Time Logger
Focus:
Structs, time.Time, duration calculation

Flow:
Start a task (records time)
Stop task
Show how long it took

Strategy:
Using struct to create a custom Task type struct which will group time related data
Hypothesis - using methods with a pointer receiver to change instance data(not just a copy as in case of value receiver method)

*/

package main

import (
	"time"
)

type Task struct {
	Name  string
	Start time.Time
	End   time.Time
}

//tasks to use: functions which require log(n^2) and log(2^n)

//methods MIGHT be redundant after all - TBC
