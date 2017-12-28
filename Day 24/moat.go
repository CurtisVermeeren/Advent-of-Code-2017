package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// A component has 2 plug values
// and a used value to indicate if it's in a bridge
type component struct {
	plug0 int
	plug1 int
	used  bool
}

// Track the componets, maxStrength and maxLength
var components []component
var maxStrength int
var maxLength int

func main() {
	part1 := solvePart1()
	fmt.Println("Part 1:", part1)

	part2 := solvePart2()
	fmt.Println("Part 2:", part2)
}

func solvePart1() int {
	components = parsePortTypes(getInput())
	maxStrength = -1
	// Start with plug 0, strength 0
	traversePortsPart1(0, 0)
	// Max strength is the strongest bridge possible
	return maxStrength
}

// Function recursively builds all possible bridges and tracks the maximum strength
func traversePortsPart1(plug, strength int) {
	// If the current strength of a bridge is greater than the max change the max
	if maxStrength < strength {
		maxStrength = strength
	}
	for i, component := range components {
		// If the component is already in the bridge skip it
		if component.used {
			continue
		}
		// If one of the plugs matches the other end add it.
		// Add the total strength of this bridge to the total
		// Set the component used and find more pieces in the bridge if possible
		if component.plug0 == plug {
			components[i].used = true
			traversePortsPart1(component.plug1, strength+component.plug0+component.plug1)
			components[i].used = false
		}
		if component.plug1 == plug {
			components[i].used = true
			traversePortsPart1(component.plug0, strength+component.plug0+component.plug1)
			components[i].used = false
		}
	}
}

func solvePart2() int {
	components = parsePortTypes(getInput())
	maxStrength, maxLength = -1, -1
	// Start with plug 0, strength 0, and length 0
	traversePortsPartTwo(0, 0, 0)
	return maxStrength
}

// Function recursively builds all possible bridges and tracks the maximum length
// It also tracks maximum strength if there is a tie in length
func traversePortsPartTwo(plug, strength, length int) {
	// Keep the longest or the longest and strongest bridge built
	if maxStrength < strength && maxLength <= length || maxLength < length {
		maxStrength = strength
		maxLength = length
	}
	for i, component := range components {
		// If the component is already in the bridge skip it
		if component.used {
			continue
		}
		// If one of the plugs matches the other end add it.
		// Add the total strength of this bridge to the total
		// Set the component used and find more pieces in the bridge if possible
		if component.plug0 == plug {
			components[i].used = true
			traversePortsPartTwo(component.plug1, strength+component.plug0+component.plug1, length+1)
			components[i].used = false
		}
		if component.plug1 == plug {
			components[i].used = true
			traversePortsPartTwo(component.plug0, strength+component.plug0+component.plug1, length+1)
			components[i].used = false
		}
	}
}

func parsePortTypes(inputData []string) []component {
	result := []component{}
	// Build components from each of the input strings
	for _, inputItem := range inputData {
		parts := strings.Split(inputItem, "/")
		plug0, _ := strconv.Atoi(parts[0])
		plug1, _ := strconv.Atoi(parts[1])
		// Add them to the component list
		result = append(result, component{plug0: plug0, plug1: plug1, used: false})
	}
	return result
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
