package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	elfs := strings.Split(input, "\n\n")

	max := make([]int, len(elfs))

	for _, elf := range elfs {
		cals := sum(elf)
		max = append(max, cals)
	}

	sort.Slice(max, func(i, j int) bool {
		return max[i] > max[j]
	})

	fmt.Println(max)
	fmt.Println(max[0] + max[1] + max[2])

}

func sum(elf string) int {
	sum := 0
	for _, cal := range strings.Split(elf, "\n") {
		parsedCal, _ := strconv.Atoi(cal)
		sum += parsedCal
	}
	return sum
}
