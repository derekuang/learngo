// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	genre string
	item  item
}

func main() {
	games := map[int]game{
		1: {
			genre: "action adventure",
			item: item{
				id:    1,
				name:  "god of war",
				price: 50,
			},
		},
		2: {
			genre: "strategy",
			item: item{
				id:    2,
				name:  "x-com 2",
				price: 30,
			},
		},
		3: {
			genre: "sandbox",
			item: item{
				id:    3,
				name:  "minecraft",
				price: 20,
			},
		},
	}

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > quit   : quits

`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		args := strings.Fields(in.Text())
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "list":
			for _, game := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n", game.item.id, game.item.name, "("+game.genre+")", game.item.price)
			}

		case "id":
			if len(args) != 2 {
				fmt.Println("wrong id")
				break
			}

			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("wrong id")
				break
			}

			if game, ok := games[id]; ok {
				fmt.Printf("#%d: %-15q %-20s $%d\n", game.item.id, game.item.name, "("+game.genre+")", game.item.price)
			} else {
				fmt.Println("sorry. I don't have the game")
			}

		case "quit":
			fmt.Println("bye!")
			return
		}
	}
}
