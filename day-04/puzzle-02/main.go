package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	min, max := parseInt(os.Args[1]), parseInt(os.Args[2])

	fmt.Println(len(combos(min, max)))
}

func parseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Could not convert '%s': %v\n", input, err)
		os.Exit(1)
	}

	return result
}

func combos(min, max int) []int {
	var result []int

	for i := min; i < max; i++ {
		if isValid(i) {
			result = append(result, i)
		}
	}

	return result
}

func isValid(i int) bool {
	if i < 100000 || i > 999999 {
		return false
	}

	current := 10
	hasAdjacents := false
	adjacentCount := 0

	for i > 0 {
		remainder := i % 10

		//fmt.Printf("i = %d, remainder = %d, current = %d, hasAdjacent = %v, adjacentCount = %d\n", i, remainder, current, hasAdjacents, adjacentCount)

		if remainder > current {
			return false
		}

		if remainder == current {
			adjacentCount++
		} else {
			if adjacentCount == 2 {
				hasAdjacents = true
			}
			adjacentCount = 1
		}

		current = remainder
		i = i / 10
	}

	return hasAdjacents || adjacentCount == 2
}
