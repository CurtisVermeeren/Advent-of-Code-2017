package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := getInput()
	programs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

	// Perform 1 dance for part 1
	part1 := solvePart1(inputs, programs)
	fmt.Println("Part 1:", strings.Join(part1, ""))

	// Track the pattern of the programs after each dance.
	var iterations []string
	part2 := programs
	part2String := ""
	repeat := 1000000000
	// Perform the dance 1 billion times
	for i := 0; i < repeat; i++ {
		// Get the string for the previous part
		part2String = strings.Join(part2, "")
		// If the order of programs in the dance has already been found (exists in iterations) then we have the cycle value.
		// The cycle value means the dance repeats every i steps
		if v := containsIndex(iterations, part2String); v != -1 {
			// We use 1 billion mod cycle to find what the value is at the 1 billionth dance
			fmt.Println("Part 2:", iterations[repeat%i])
			break
		}
		// If the order didn't exist we add it to the iterations then get the next iteration
		iterations = append(iterations, part2String)
		part2 = solvePart1(inputs, part2)
	}

}

func solvePart1(inputs []string, programs []string) []string {
	for _, input := range inputs {
		instruction := input[:1]

		switch instruction {
		// Spin
		case "s":
			// The number of values to be moves
			valueToMove, _ := strconv.Atoi(input[1:])
			// Get the last number of values
			toSpin := programs[len(programs)-valueToMove:]
			// Get the remaining values
			toRemain := programs[0 : len(programs)-valueToMove]
			// Put the spin values at the front and append the remaining
			programs = append(toSpin, toRemain...)
		// Exchange
		case "x":
			indexes := strings.Split(input[1:], "/")
			index1, err := strconv.Atoi(indexes[0])
			index2, err := strconv.Atoi(indexes[1])
			if err != nil {
				log.Fatal(err)
			}
			// Swap the indexes specified
			programs[index1], programs[index2] = programs[index2], programs[index1]
		// Partner
		case "p":
			indexes := strings.Split(input[1:], "/")
			// Find the indexes of the values specified then swap those indexes
			index1 := containsIndex(programs, indexes[0])
			index2 := containsIndex(programs, indexes[1])
			programs[index1], programs[index2] = programs[index2], programs[index1]
		}
	}
	return programs
}

// Function returns a slice of strings, each string is an instruction for the dance
func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, strings.Split(scanner.Text(), ",")...)
	}
	return inputs
}

// Function checks if a slice of strings str contains the string s
// returns the index of s if found, returns -1 if s is not in str
func containsIndex(str []string, s string) int {
	for i, value := range str {
		if value == s {
			return i
		}
	}
	return -1
}
