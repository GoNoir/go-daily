/*
Dice Game
Focus: Random numbers, loops, conditionals
What it does: Simple "guess if next roll is higher/lower" game

Roll a die (1-6), show number (create a dice - enums?) 
Ask user: will next roll be higher/lower?
Roll again, tell them if they won
Keep score, play 5 rounds
*/

package main

import (
	"fmt"
	"math/rand"
	//"bufio"
	//"os"
)

//create dice, for now - slice (struct or enums could be better - why?)

//roll dice - pseudorandomness from rand library. Good enough for this, but not secure enough for security purpose.

func DiceRoll (s []int) int {
return rand.Intn(len(s))
}

//ask user for the next one - need another library - bufio?

//roll dice again
//print won/lose

//add to score

//repeat 5 times.

func main() {
//keep score
score := 0
//loop 5 times

//for i := 1; i < 5; i++ {
//dice is a slice
dice := []int{1, 2, 3, 4, 5, 6}
x := DiceRoll(dice)
fmt.Println(x)
//ask user for prediction: need bufio/scanner for command line input
//fmt.Printf("The last score was: %d", x)
//fmt.Printf("Will next roll be higher or lower?"
//takes h or l user input
//y := DiceRoll(dice)
switch
case h
//if //user input h and y > x {
score++
fmt.Println("You won!")
//else {
fmt.Println("You lose :("
case l
//

//update score and print for user
//print final score
}
