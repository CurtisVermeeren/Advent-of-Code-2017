package main

import (
	"fmt"
)

func main() {
	// Puzzle Input
	nextA := 703
	nextB := 516

	fmt.Println("Part1", solvePart1(nextA, nextB))
	fmt.Println("Part2", solvePart2(nextA, nextB))
}

// Function solves part 1
func solvePart1(nextA, nextB int) int {
	count := 0
	for i := 0; i < 40000000; i++ {
		nextA = nextValue1(16807, nextA)
		nextB = nextValue1(48271, nextB)
		// Use a bitmask to compare the last 16 bits of A and B
		// 0xFFFF is all ones in binary so we know that is the bit in nextA is a 1 then the result is a 1
		// If the bit in nextB is a 0 then 1x0 = 0 so it is a zero.
		// Each hex character is 4 bits so we use 1111 1111 1111 1111
		if nextA&0xFFFF == nextB&0xFFFF {
			count++
		}
	}
	return count
}

// Function solves part 2
func solvePart2(nextA, nextB int) int {
	count := 0
	for i := 0; i < 5000000; i++ {
		nextA = nextValue2(16807, nextA, 4)
		nextB = nextValue2(48271, nextB, 8)
		// Use a bitmask to compare the last 16 bits of A and B
		if nextA&0xFFFF == nextB&0xFFFF {
			count++
		}
	}
	return count
}

// Function calculates values for part 1
func nextValue1(factor, value int) int {
	val := (value * factor) % (2147483647)
	return val
}

// Function calculates values for part 2
func nextValue2(factor, value, multiple int) int {
	val := (value * factor) % (2147483647)
	// Keep getting values until it's a multiple of the desired value
	for val%multiple != 0 {
		val = (val * factor) % (2147483647)
	}
	return val
}
