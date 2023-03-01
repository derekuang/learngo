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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

func main() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	type jsonGame struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
	}

	jGames := []jsonGame{}
	json.Unmarshal([]byte(data), &jGames)

	games := []game{}
	for _, jGame := range jGames {
		games = append(games, game{item{jGame.ID, jGame.Name, jGame.Price}, jGame.Genre})
	}

	// index the games by id
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : exports the data to json and quits
  > quit   : quits

`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)

		case "save":
			// load the data into the encodable game values
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable,
					jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}

			fmt.Println(string(out))
			return
		}
	}
}
