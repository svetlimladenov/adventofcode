package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	pairs := strings.Split(input, "\n")
	counter := 0
	for _, pair := range pairs {
		assignments := strings.Split(pair, ",")
		firstAssigmentRange := strings.Split(assignments[0], "-")
		secondAssigmentRange := strings.Split(assignments[1], "-")

		firstAssigmentStart, _ := strconv.Atoi(firstAssigmentRange[0])
		firstAssigmentEnd, _ := strconv.Atoi(firstAssigmentRange[1])
		secondAssigmentStart, _ := strconv.Atoi(secondAssigmentRange[0])
		secondAssigmentEnd, _ := strconv.Atoi(secondAssigmentRange[1])

		if firstAssigmentStart >= secondAssigmentStart && firstAssigmentEnd <= secondAssigmentEnd {
			fmt.Println(assignments)
			counter++
			continue
		}

		if secondAssigmentStart >= firstAssigmentStart && secondAssigmentEnd <= firstAssigmentEnd {
			counter++
			fmt.Println(assignments)
			continue
		}
	}
	fmt.Println(counter)
}
