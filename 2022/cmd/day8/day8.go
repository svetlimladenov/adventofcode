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
	matrix := parseInput(input)
	// visible := make([]string, 0)

	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[col])-1; col++ {
			before := matrix[row][0:col]
			after := matrix[row][col+1 : len(matrix[row])]
			fmt.Printf("before: %v\n", before)
			fmt.Printf("after: %v\n", after)
			fmt.Printf("matrix[row][col]: %v\n", matrix[row][col])

		}
		fmt.Println()
	}
}

func parseInput(input string) [][]int {
	var matrix [][]int

	rows := strings.Split(input, "\n")
	for _, line := range rows {
		matrix = append(matrix, parseLine(line))
	}

	return matrix
}

func parseLine(line string) []int {
	numbers := strings.Split(line, "")
	var result []int
	for _, n := range numbers {
		parsed, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		result = append(result, parsed)
	}

	return result
}

func getNeighbours(matrix [][]int, row, col int) (int, int, int, int) {
	up := -1
	down := -1
	left := -1
	right := -1

	if row != 0 {
		up = matrix[row-1][col]
	}

	if col != 0 {
		left = matrix[row][col-1]
	}

	if row != len(matrix)-1 {
		down = matrix[row+1][col]
	}

	if col != len(matrix[row])-1 {
		right = matrix[row][col+1]
	}

	return up, down, left, right
}

func isVisible(matrix [][]int, row, col int) bool {
	cur := matrix[row][col]

	up, down, left, right := getNeighbours(matrix, row, col)

	if cur < up && cur < down && cur < left && cur < right {
		return true
	}

	return false
}
