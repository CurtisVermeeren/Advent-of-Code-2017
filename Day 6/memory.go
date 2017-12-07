package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// blocks is the input given
	blocks := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}

	fmt.Println(solvePart1and2(blocks))
}

// Function with logic to solve the days problem
func solvePart1and2(blocks []int) int {
	steps := 0
	// Use a map to check if a combination has been used before
	check := make(map[string]bool)
	current := intSliceToString(blocks)
	for !check[current] {
		check[current] = true

		// Find the index with the most blocks
		startIndex := findIndexOfMost(blocks)
		// Determine the number of blocks to move
		movedBlocks := blocks[startIndex]
		blocks[startIndex] = 0
		startIndex++
		// Move the blocks into the next indexes
		for movedBlocks > 0 {
			blocks[(startIndex)%(len(blocks))]++
			movedBlocks--
			startIndex++
		}
		// Create a string from the slice of integers so it can be used as a map key
		current = intSliceToString(blocks)
		steps++
	}

	// When the first pattern is repeated store that value
	matched := current

	steps2 := 0
	// For part 2 we continue moving blocks between indexes until the match is found again
	for {
		check[current] = true

		startIndex := findIndexOfMost(blocks)
		movedBlocks := blocks[startIndex]
		blocks[startIndex] = 0
		startIndex++
		for movedBlocks > 0 {
			blocks[(startIndex)%(len(blocks))]++
			movedBlocks--
			startIndex++
		}
		current = intSliceToString(blocks)

		steps2++
		if current == matched {
			break
		}
	}
	fmt.Println(steps2)
	return steps
}

// Function converts a slice of integers into a string
func intSliceToString(blocks []int) string {
	var stringSlice []string
	for _, value := range blocks {
		stringSlice = append(stringSlice, strconv.Itoa(value))
	}
	return strings.Join(stringSlice, ",")
}

// Returns the index of the bank with the most blocks
func findIndexOfMost(blocks []int) int {
	max := 0
	for index, value := range blocks {
		if value > blocks[max] {
			max = index
			// If two values have the same number of blocks the smaller index is returned
		} else if value == blocks[max] {
			if index < max {
				max = index
			}
		}
	}
	return max
}
