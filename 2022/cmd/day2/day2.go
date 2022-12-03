package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	games := strings.Split(input, "\n")
	score := 0
	for _, game := range games {
		points := getGamePoints(game)
		fmt.Println(game, points)
		score += points
	}

	fmt.Println(score)
}

var allPossible = map[string]int{
	"A X": 3 + 0, // scissors lose
	"A Y": 3 + 1, // rock + draw
	"A Z": 2 + 6, // paper + win

	"B X": 1 + 0, // rock + lose
	"B Y": 3 + 2, // draw
	"B Z": 3 + 6, // scissors win

	"C X": 2 + 0, // paper + lose
	"C Y": 3 + 3,
	"C Z": 1 + 6, // rock + win
}

func getGamePoints(game string) int {
	return allPossible[game]
}
