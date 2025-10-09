/*
Dice Game
Focus: Random numbers, for loop, switch and if, user input in cli: bufio/scan
What it does: Simple "guess if next roll is higher/lower" game

Roll a die (1-6), show number
Ask user: will next roll be higher/lower?
Roll again, tell them if they won
Keep score, play 5 rounds
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func DiceRoll(s []int) int {
	return s[rand.Intn(len(s))]
}

func main() {
	score := 0
	for i := 1; i < 6; i++ {
		dice := []int{1, 2, 3, 4, 5, 6}
		x := DiceRoll(dice)
		fmt.Printf("First roll was: %d:", x)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Will next roll be [h]igher or [l]ower?:")
		scanner.Scan()
		input := scanner.Text()
		if input != "h" && input != "l" {
			fmt.Printf("Sorry, can only accept 'h' for higher or 'l' for lower:")
			scanner.Scan()
			input = scanner.Text()
		}
		fmt.Printf("Yout typed: %q, let's roll the dice...", input)

		y := DiceRoll(dice)

		switch input {
		case "h":
			if y > x {
				score++
				fmt.Println("You won!")
			} else {
				fmt.Println("You lose :/")
			}
		case "l":
			if y < x {
				score++
				fmt.Println("You won!")
			} else {
				fmt.Println("You lose :/")
			}

		}

	}
	fmt.Printf("Your final score is: %d", score)
}
