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

		// .23......  2-3
		// ...45....  4-5

		// .....678.. 6-8
		// 123....... 1-3
		if secondAssigmentStart > firstAssigmentEnd {
			fmt.Println(assignments)
			continue
		}

		if secondAssigmentEnd < firstAssigmentStart {
			fmt.Println(assignments)
			continue
		}
		counter++
	}
	fmt.Println(counter)
}

func isTotalOverlap(firstAssigmentStart, firstAssigmentEnd, secondAssigmentStart, secondAssigmentEnd int) bool {
	if firstAssigmentStart >= secondAssigmentStart || firstAssigmentEnd <= secondAssigmentEnd {
		return true
	}

	if secondAssigmentStart >= firstAssigmentStart || secondAssigmentEnd <= firstAssigmentEnd {
		return true
	}
	return false
}
