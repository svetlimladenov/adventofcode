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

	secondPart(rucksacks)
}

func secondPart(rucksacks []string) {
	result := 0

	for i := 0; i < len(rucksacks); i += 3 {
	group:
		for firstElfIndex := 0; firstElfIndex < len(rucksacks[i]); firstElfIndex++ {
			for secondElfIndex := 0; secondElfIndex < len(rucksacks[i+1]); secondElfIndex++ {
				if rucksacks[i][firstElfIndex] == rucksacks[i+1][secondElfIndex] {
					for thirdElfIndex := 0; thirdElfIndex < len(rucksacks[i+2]); thirdElfIndex++ {
						if rucksacks[i][firstElfIndex] == rucksacks[i+2][thirdElfIndex] {
							result += calculateTypeScore(int(rucksacks[i][firstElfIndex]))
							break group
						}
					}
				}
			}
		}
	}
	fmt.Println(result)
}

func firstPart(rucksacks []string) {
	score := 0
	for _, sack := range rucksacks {
	sackLoop:
		for i := 0; i < len(sack)/2; i++ {
			for j := len(sack) / 2; j < len(sack); j++ {
				if sack[i] == sack[j] {
					score += calculateTypeScore(int(sack[i]))
					break sackLoop
				}
			}
		}
	}

	fmt.Println(score)
}

func calculateTypeScore(itemType int) int {
	if itemType > 'Z' {
		return int(itemType) - 96
	} else {
		return int(itemType) - 38
	}
}
