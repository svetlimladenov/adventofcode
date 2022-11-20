package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := "easy.txt"

	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(file)

	var matrix [][]int

	for s.Scan() {

		matrix = append(matrix, parseLine(s.Text()))
	}

	result := findLowPoints(matrix)
	fmt.Println(result)
}

func findLowPoints(matrix [][]int) int {
	sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if isLowPoint(matrix, row, col) {
				sum += matrix[row][col] + 1
			}
		}
	}
	return sum
}

func isLowPoint(matrix [][]int, row, col int) bool {
	cur := matrix[row][col]

	up, down, left, right := getNeighbours(matrix, row, col)

	if cur < up && cur < down && cur < left && cur < right {
		return true
	}

	return false
}

func traceBasin(matrix [][]int, row, col int) {
	// lowPoint := matrix[row][col]

	// up := 10
	// down := 10
	// left := 10
	// right := 10

	// if row != 0 {
	// 	up = matrix[row-1][col]
	// }

	// if col != 0 {
	// 	left = matrix[row][col-1]
	// }

	// if row != len(matrix)-1 {
	// 	down = matrix[row+1][col]
	// }

	// if col != len(matrix[row])-1 {
	// 	right = matrix[row][col+1]
	// }

	// if up == lowPoint+1 {
	// 	traceBasin(matrix, row+1, col)
	// }

}

func getNeighbours(matrix [][]int, row, col int) (int, int, int, int) {
	up := 10
	down := 10
	left := 10
	right := 10

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
