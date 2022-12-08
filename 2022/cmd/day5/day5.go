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
	inputParts := strings.Split(input, "\n\n")

	stacks := strings.Split(inputParts[0], "\n")

	data := make(map[int][]byte)

	for i := len(stacks) - 2; i >= 0; i-- {
		line := stacks[i]
		pile := 1
		for i := 1; i < len(line); i += 4 {
			if string(line[i]) != " " {
				data[pile] = append(data[pile], line[i])
			} else if len(data[pile]) == 0 {
				data[pile] = make([]byte, 0)
			}
			pile++
		}
	}

	a := data[2][2:3]
	fmt.Println(a)
	commands := strings.Split(inputParts[1], "\n")

	for _, cmd := range commands {
		commandParts := strings.Split(cmd, " ")
		from, _ := strconv.Atoi(commandParts[3])
		to, _ := strconv.Atoi(commandParts[5])
		count, _ := strconv.Atoi(commandParts[1])

		fromLastElementIndex := len(data[from])
		elementsToMove := data[from][fromLastElementIndex-count : fromLastElementIndex]
		data[to] = append(data[to], elementsToMove...)
		data[from] = data[from][:fromLastElementIndex-count]
	}

	// map iteration with range is not stable !!!
	for i := 1; i < len(data)+1; i++ {
		array := data[i]
		fmt.Print(string(array[len(array)-1]))
	}

	fmt.Println()
}
