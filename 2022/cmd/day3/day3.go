package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rucksacks := strings.Split(input, "\n")

	score := 0
	for _, sack := range rucksacks {
	sackLoop:
		for i := 0; i < len(sack)/2; i++ {
			for j := len(sack) / 2; j < len(sack); j++ {
				if sack[i] == sack[j] {
					if sack[i] > 'Z' {
						fmt.Println(string(sack[i]), int(sack[i]-96))
						score += int(sack[i]) - 96
					} else {
						fmt.Println(string(sack[i]), int(sack[i]-38))
						score += int(sack[i]) - 38
					}
					break sackLoop
				}
			}
		}
	}

	fmt.Println(score)
}
