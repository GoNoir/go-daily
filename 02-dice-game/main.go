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
)

//create dice, for now - slice (struct or enums could be better - why?)

//roll dice - pseudorandomness from rand library. Good enough for this, but not secure enough for security purpose.
//ask user for the next one - need another library - bufio?

//roll dice again
//print won/lose

//add to score

//repeat 5 times.

func main() {
//keep score
score := 0
//dice is a slice
dice := []int{1, 2, 3, 4, 5, 6}
//roll - we need to extrapolate roll later for DRY - needs to be checked agains the predition, like lastRoll := DiceRoll(dice)
fmt.Println(dice[rand.Intn(len(dice))])
//Ask user to give higher or lower prediction eg. h or l 

//Loop 4 more times
//for i := 1; i < 5; i++ {
//roll dice again
//update score and print for user


//print final score
}
