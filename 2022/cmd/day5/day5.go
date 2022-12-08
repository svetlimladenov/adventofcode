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

	commands := strings.Split(inputParts[1], "\n")

	for _, cmd := range commands {
		commandParts := strings.Split(cmd, " ")
		from, _ := strconv.Atoi(commandParts[3])
		to, _ := strconv.Atoi(commandParts[5])
		count, _ := strconv.Atoi(commandParts[1])

		for i := 0; i < count; i++ {
			fromLastElementIndex := len(data[from]) - 1
			elementToMove := data[from][fromLastElementIndex]
			data[to] = append(data[to], elementToMove)
			data[from] = data[from][:fromLastElementIndex]
		}
	}

	// map iteration with range is not stable !!!
	for i := 1; i < len(data)+2; i++ {
		array := data[i]
		fmt.Print(string(array[len(array)-1]))
	}

	fmt.Println()
}

func print(data map[int][]byte) {
	for i, v := range data {
		for _, c := range v {
			fmt.Println(string(c), i)
		}
		fmt.Println()
	}
}
