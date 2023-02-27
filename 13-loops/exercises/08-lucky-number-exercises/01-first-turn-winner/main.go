// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
	notNumErr = `Not a number.`
	notPosErr = `Please pick a positive number.`
	firstMsg  = `"🥇 FIRST TIME WINNER!!!"`
	winMsg    = `🎉  YOU WON!`
	lostMsg   = `☠️  YOU LOST... Try again?`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		println(usage)
		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		println(notNumErr)
		return
	}

	if guess <= 0 {
		println(notPosErr)
	}

	for turn := 1; turn <= maxTurns; turn++ {
		res := rand.Intn(guess) + 1

		if guess != res {
			continue
		}

		if turn == 1 {
			println(firstMsg)
		} else {
			println(winMsg)
		}
		return
	}
	println(lostMsg)
}
