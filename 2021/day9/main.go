package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := "input.txt"

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
	basins := make([]int, 0)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if isLowPoint(matrix, row, col) {
				basin := make([]string, 0)
				traceBasin(matrix, row, col, &basin)
				basins = append(basins, len(basin))
			}
		}
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	fmt.Println(basins[0])
	fmt.Println(basins[1])
	fmt.Println(basins[2])

	return basins[0] * basins[1] * basins[2]
}

func isLowPoint(matrix [][]int, row, col int) bool {
	cur := matrix[row][col]

	up, down, left, right := getNeighbours(matrix, row, col)

	if cur < up && cur < down && cur < left && cur < right {
		return true
	}

	return false
}

func traceBasin(matrix [][]int, row, col int, basin *[]string) {
	lowPoint := matrix[row][col]
	if contains(*basin, fmt.Sprintf("%d:%d-%d", row, col, lowPoint)) {
		return
	}
	if lowPoint == 9 {
		return
	}

	*basin = append(*basin, fmt.Sprintf("%d:%d-%d", row, col, lowPoint))

	up, down, left, right := getNeighbours(matrix, row, col)

	if up != 10 && up > lowPoint {
		traceBasin(matrix, row-1, col, basin)
	}

	if down != 10 && down > lowPoint {
		traceBasin(matrix, row+1, col, basin)
	}

	if left != 10 && left > lowPoint {
		traceBasin(matrix, row, col-1, basin)
	}

	if right != 10 && right > lowPoint {
		traceBasin(matrix, row, col+1, basin)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
