package main

import (
	"fmt"
)

func main() {
	stepsize := 349
	buffer := createBuffer(2017, stepsize)
	part1 := findTarget(buffer, 2017)
	fmt.Println("Part 1:", part1)
	part2 := numberAfterZero(50000000, stepsize)
	fmt.Println("Part 2:", part2)
}

// Returns the values in a buffer after the targetNumber
// Returns -1 if targetNumber is not found
func findTarget(buffer []int, targetNumber int) int {
	for i, element := range buffer {
		if element == targetNumber {
			return buffer[(i+1)%len(buffer)]
		}
	}
	return -1
}

func numberAfterZero(maxNumber int, stepsize int) int {
	insertNumber := 1
	currentPosition := 0
	// Track the number after 0
	numberAfterZero := -1
	// Perform maxNumber (50 million) inserts
	for insertNumber <= maxNumber {
		// Calculate the current position as before
		currentPosition = (currentPosition + stepsize) % insertNumber
		currentPosition++
		// We know that 0 is always at position 0 as nothing can be inserted before that
		// Therefore when the current position is index 1 the value after 0 is being changed and we need to update that value
		if currentPosition == 1 {
			numberAfterZero = insertNumber
		}
		// the insertNumber increases by 1 again
		insertNumber++
	}
	// Return the final numberAfterZero after 50 million inserts
	return numberAfterZero
}

// Function builds a buffer with a maximum number of values using stepsize
func createBuffer(maxNumber int, stepsize int) []int {
	// Use an int slice as the splinlocks buffer
	buffer := make([]int, 1)
	// Current position starts at 0
	currentPosition := 0
	buffer[currentPosition] = 0
	for insertNumber := 1; insertNumber <= maxNumber; insertNumber++ {
		// The next currentPosition is the currentPosition plus the step size, then modulo divide by length for circular slice
		currentPosition = (currentPosition + stepsize) % len(buffer)
		// Add 1 because we inserting a new value AFTER the current position
		currentPosition++
		// If the current position is at the end we can simply add the value on to the buffer
		if currentPosition == len(buffer) {
			buffer = append(buffer, insertNumber)
		} else {
			// Rest is all the values after the current position
			rest := make([]int, len(buffer[currentPosition:]))
			copy(rest, buffer[currentPosition:])
			// We make a slice of the insertedNumber then append each of the values after it
			toAdd := append([]int{insertNumber}, rest...)
			// We then add the insertedNumber and the values following it back to the buffer
			buffer = append(buffer[0:currentPosition], toAdd...)
		}
	}
	return buffer
}
