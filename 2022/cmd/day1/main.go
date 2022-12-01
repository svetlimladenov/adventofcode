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
	elfs := strings.Split(input, "\n\n")

	max := 0
	for _, elf := range elfs {
		cals := sum(elf)
		if cals > max {
			max = cals
		}
	}

	fmt.Println(max)
}

func sum(elf string) int {
	sum := 0
	for _, cal := range strings.Split(elf, "\n") {
		parsedCal, _ := strconv.Atoi(cal)
		sum += parsedCal
	}
	return sum
}
