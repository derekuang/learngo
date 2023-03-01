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
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	genre string
	item
}

type jgame struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
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
  > save   : exports the data to json and quits
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
				fmt.Printf("#%d: %-15q %-20s $%d\n", game.id, game.name, "("+game.genre+")", game.price)
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
				fmt.Printf("#%d: %-15q %-20s $%d\n", game.id, game.name, "("+game.genre+")", game.price)
			} else {
				fmt.Println("sorry. I don't have the game")
			}

		case "save":
			jgames := []jgame{}
			for _, game := range games {
				jgames = append(jgames, jgame{Id: game.id, Name: game.name, Genre: game.genre, Price: game.price})
			}
			out, _ := json.MarshalIndent(jgames, "", "\t")
			fmt.Println(string(out))
			return

		case "quit":
			fmt.Println("bye!")
			return
		}
	}
}
