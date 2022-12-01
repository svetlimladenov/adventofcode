package parse

import "strconv"

// ParseInt parses an integer and panics on error
func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}
