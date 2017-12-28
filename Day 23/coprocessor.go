package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := getInput()
	part1 := solvePart1(instructions)
	fmt.Println(part1)
	part2 := solvePart2()
	fmt.Println(part2)
}

func solvePart1(inputs []string) int {
	registers := make(map[string]int)
	count := 0
	for i := 0; i > -1 && i < len(inputs); i++ {
		// Split the instruction into parts for parsing
		input := inputs[i]
		parts := strings.Split(input, " ")
		// Use a switch to detect the instruction type
		switch parts[0] {
		// Perform the operations based on the instruction.
		// Works the same as Day 18's duet
		case "set":
			if isNumber(parts[2]) {
				registers[parts[1]] = toNumber(parts[2])
			} else {
				registers[parts[1]] = registers[parts[2]]
			}
		case "sub":
			if isNumber(parts[2]) {
				registers[parts[1]] -= toNumber(parts[2])
			} else {
				registers[parts[1]] -= registers[parts[2]]
			}
		case "mul":
			if isNumber(parts[2]) {
				registers[parts[1]] = registers[parts[1]] * toNumber(parts[2])
			} else {
				registers[parts[1]] = registers[parts[1]] * registers[parts[2]]
			}
			// Keep track of all the multiplication operations
			count++
		case "jnz":
			if isNumber(parts[1]) {
				if toNumber(parts[1]) != 0 {
					if isNumber(parts[2]) {
						i += toNumber(parts[2]) - 1
					} else {
						i += registers[parts[2]] - 1
					}
				}
			} else {
				if registers[parts[1]] != 0 {
					if isNumber(parts[2]) {
						i += toNumber(parts[2]) - 1
					} else {
						i += registers[parts[2]] - 1
					}
				}
			}
		}
	}
	// Returns the number of multiplications performed
	return count
}

// Function builds a slice of strings from the input file
func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	inputs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}
	return inputs
}

// Function converts a string to an integer with reusable error checking
func toNumber(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

// Returns true is str is number, false otherwise
func isNumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

func solvePart2() int {
	// By analyzing the instructions we can refractor and determine that h only increments when b is not a prime number
	// b starts at 105700 after the first few instructions and it runs until 12270 becuase c = b then c += 17000.
	// The instruction sub b -17 means we add 17 to b each loop
	h := 0
	for b := 105700; b <= 122700; b += 17 {
		if !isPrimeSqrt(b) {
			h++
		}
	}
	return h
}

// Function determines if the value is a prime number
func isPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
