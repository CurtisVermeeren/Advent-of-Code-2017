package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	strings, err := readFileToSlice("input.txt")
	integers, err := stringToIntSlice(strings)
	if err != nil {
		panic(err)
	}
	integers2 := make([]int, len(integers))
	copy(integers2, integers)
	fmt.Println(solvePart1(integers))
	fmt.Println(solvePart2(integers2))

}

// Functions reads the path to a file and reads each line into a []string
func readFileToSlice(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Function takes a slice of strings and converts to a slice of integers
func stringToIntSlice(strings []string) ([]int, error) {
	var ints = []int{}
	for _, i := range strings {
		integer, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		ints = append(ints, integer)
	}
	return ints, nil
}

func solvePart1(integers []int) int {
	index := 0
	steps := 0
	// While the exit has not been found
	for index < len(integers) {
		// Get the value at the current index
		indexValue := integers[index]
		// Increment the value at the current index
		integers[index] = indexValue + 1
		// Offset the index
		index += indexValue
		// Increment the number of steps needed
		steps++
	}
	return steps
}

func solvePart2(integers []int) int {
	index := 0
	steps := 0
	// While the exit has not been found
	for index < len(integers) {
		// Get the value at the current index
		indexValue := integers[index]
		// If the value was 3 or more decrement the value otherwise increment
		if indexValue >= 3 {
			integers[index] = indexValue - 1
		} else {
			integers[index] = indexValue + 1
		}
		// Offset the index
		index += indexValue
		// Increment the number of steps needed
		steps++
	}
	return steps
}
