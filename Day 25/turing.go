package main

import (
	"fmt"
)

func main() {
	solvePart1()
}

func solvePart1() {
	// Track the 1 and 0 values
	tape := make(map[int]int)
	// Track current state A,B,C,D,E,F
	currentState := "A"
	// Track the current slot
	// increasing moves to the right, decreasing to the left
	currentSlot := 0
	// Perform all the steps
	for steps := 0; steps < 12683008; steps++ {
		switch currentState {
		// Perform all the cases accoding to the input
		case "A":
			if tape[currentSlot] == 0 {
				tape[currentSlot] = 1
				currentSlot++
				currentState = "B"
			} else {
				tape[currentSlot] = 0
				currentSlot--
				currentState = "B"
			}
		case "B":
			if tape[currentSlot] == 0 {
				tape[currentSlot] = 1
				currentSlot--
				currentState = "C"
			} else {
				tape[currentSlot] = 0
				currentSlot++
				currentState = "E"
			}
		case "C":
			if tape[currentSlot] == 0 {
				tape[currentSlot] = 1
				currentSlot++
				currentState = "E"
			} else {
				tape[currentSlot] = 0
				currentSlot--
				currentState = "D"
			}
		case "D":
			if tape[currentSlot] == 0 {
				tape[currentSlot] = 1
				currentSlot--
				currentState = "A"
			} else {
				tape[currentSlot] = 1
				currentSlot--
				currentState = "A"
			}
		case "E":
			if tape[currentSlot] == 0 {
				tape[currentSlot] = 0
				currentSlot++
				currentState = "A"
			} else {
				tape[currentSlot] = 0
				currentSlot++
				currentState = "F"
			}
		case "F":
			if tape[currentSlot] == 0 {
				tape[currentSlot] = 1
				currentSlot++
				currentState = "E"
			} else {
				tape[currentSlot] = 1
				currentSlot++
				currentState = "A"
			}
		}
	}
	// Count all the 1 values in the map
	count := 0
	for _, value := range tape {
		if value == 1 {
			count++
		}
	}
	fmt.Println("Part 1:", count)
}
