package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Type node holds x and y coordinates
type node struct {
	x, y int
}

// Travel directions.
const (
	north int = iota
	east
	south
	west
)

// States of infection
const (
	clean int = iota
	weakened
	infected
	flagged
)

func main() {
	lines := getInput()
	infections := getInfectionMap(lines)
	infections2 := getInfectionMap(lines)
	part1 := solvePart1(infections)
	fmt.Println("Part 1:", part1)
	part2 := solvePart2(infections2)
	fmt.Println("Part 2", part2)
}

func solvePart1(infections map[node]int) int {
	currentNode := node{0, 0}
	direction := north
	infectionCount := 0

	for i := 0; i < 10000; i++ {

		// Check the current node
		switch infections[currentNode] {
		case clean:
			// Turn left and infect the current node and increment the counter of infected
			direction = (direction + 3) % 4
			infectionCount++
			infections[currentNode] = infected
		case infected:
			// Turn right and clean the current node
			direction = (direction + 1) % 4
			infections[currentNode] = clean
		}

		// Move the current node
		switch direction {
		case north:
			currentNode = node{currentNode.x, currentNode.y - 1}
		case east:
			currentNode = node{currentNode.x + 1, currentNode.y}
		case south:
			currentNode = node{currentNode.x, currentNode.y + 1}
		case west:
			currentNode = node{currentNode.x - 1, currentNode.y}
		}

	}
	return infectionCount
}

func solvePart2(infections map[node]int) int {
	currentNode := node{0, 0}
	direction := north
	infectionCount := 0

	for i := 0; i < 10000000; i++ {
		// Check the current node
		switch infections[currentNode] {
		case clean:
			// Turn left and weaken the node
			direction = (direction + 3) % 4
			infections[currentNode] = weakened
		case weakened:
			// Continue direction and infect the node
			// Increment the number of infected nodes
			infectionCount++
			infections[currentNode] = infected
		case infected:
			// Turn right and flag the node
			direction = (direction + 1) % 4
			infections[currentNode] = flagged
		case flagged:
			// Turn around and clean thenode
			direction = (direction + 2) % 4
			infections[currentNode] = clean
		}
		// Move the current node
		switch direction {
		case north:
			currentNode = node{currentNode.x, currentNode.y - 1}
		case east:
			currentNode = node{currentNode.x + 1, currentNode.y}
		case south:
			currentNode = node{currentNode.x, currentNode.y + 1}
		case west:
			currentNode = node{currentNode.x - 1, currentNode.y}
		}

	}
	return infectionCount
}

func getInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Build a 2D slice of all characters in the grid
	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, ""))
	}
	return lines
}

func getInfectionMap(lines [][]string) (infections map[node]int) {
	// Use a map of nodes to int to map coordinates of infected areas
	infections = make(map[node]int)
	sizeLines := len(lines)
	for x := 0; x < sizeLines; x++ {
		for y := 0; y < sizeLines; y++ {
			// node with x and y is infected add it to the map
			if lines[x][y] == "#" {
				node := node{y - sizeLines/2, x - sizeLines/2}
				infections[node] = infected
			}
		}
	}

	return
}
