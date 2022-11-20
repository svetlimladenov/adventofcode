package main

import (
	"bufio"
	"fmt"
	"os"
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

func findLowPoints(matrix [][]int) []int {
	result := make([]int, 0)

	sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			cur := matrix[row][col]

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

			if cur < up && cur < down && cur < left && cur < right {
				fmt.Printf("Low point - %d\n", cur)
				sum += cur + 1
			}
		}
	}
	fmt.Println(sum)
	result = append(result, sum)

	return result
}

func parseLine(line string) []int {
	numbers := strings.Split(line, "")
	var result []int
	for _, n := range numbers {
		parsed, err := strconv.Atoi(n)
		if err != nil {

		}

		result = append(result, parsed)
	}

	return result
}
