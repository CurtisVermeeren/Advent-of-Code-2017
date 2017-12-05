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
	checksum := 0
	evenSum := 0

	// Open input file to read
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan line by line of input file
	for scanner.Scan() {
		// strings.Fields converts a string into []string around any length of whitespace
		numberLine := strings.Fields(scanner.Text())

		// PART 1
		// --------------------------------------------------------
		largest := 0
		smallest, _ := strconv.Atoi(numberLine[0])
		for _, value := range numberLine {
			// Convert value to integer
			intValue, _ := strconv.Atoi(value)

			// Check largest value
			if intValue > largest {
				largest = intValue
			}

			// Check smallest value
			if intValue < smallest {
				smallest = intValue
			}
		}
		// Add the difference of largest and smallest in a line to checksum
		checksum += largest - smallest

		// PART 2
		// --------------------------------------------------------
		for index, value := range numberLine {
			for index2, value2 := range numberLine {
				// Skip same values. They will always divide themselves
				if index == index2 {
					continue
				}

				intValue, _ := strconv.Atoi(value)
				intValue2, _ := strconv.Atoi(value2)

				// If no remainder the two values divide therefore we add their division to the total answer
				if (intValue % intValue2) == 0 {
					evenSum += intValue / intValue2
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(checksum)
	fmt.Println(evenSum)
}
