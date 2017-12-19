package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// Type coord hold the x and y coordinate values
type coord struct {
	x int
	y int
}

func main() {
	lines := getInput()
	array, currentPos, ymax, xmax := createGrid(lines)
	solution, numbersteps := traverse(currentPos, ymax, xmax, array)
	fmt.Println(solution, numbersteps)
}

func traverse(currentPos coord, ymax int, xmax int, array [][]string) (string, int) {
	// xdir and ydir track the direction of movement. We start moving down
	xdir := 0
	ydir := 1
	// Soltuion will hold a string of visited letters
	solution := ""
	// visited tracks if coordinates have been visited
	visited := make(map[coord]bool)
	// The number of steps needed for part 2
	numbersteps := 0
	// While we are still within the grid
	for currentPos.y < ymax && currentPos.y >= 0 && currentPos.x < xmax && currentPos.x >= 0 {
		// Get the current character from the grid using the current coordinates
		currentChar := array[currentPos.y][currentPos.x]
		// If we've found the exit then break from the loop
		if currentChar == " " {
			break
		}
		// Set the character to visited
		visited[currentPos] = true
		// Check which direction to move, if | or - continue in the current direction
		switch currentChar {
		case "|":
			currentPos.x += xdir
			currentPos.y += ydir
		case "-":
			currentPos.x += xdir
			currentPos.y += ydir
		// A corner needs to check surroundings
		case "+":
			var next coord
		Loop:

			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					// Create the next coordinate
					next = coord{currentPos.x + x, currentPos.y + y}
					// If the next is out of the grid skip it
					if next.y >= ymax || next.y < 0 || next.x >= xmax || next.x < 0 {
						continue
					}
					// If the next is a diagonal skip it
					if math.Abs(float64(x))+math.Abs(float64(y)) == 2 {
						continue
					}
					// If the next hasn't been visited and is not an empty space move to it and break the loop
					if visited[next] == false && array[next.y][next.x] != " " && array[next.y][next.x] != "" {
						next = next
						xdir = x
						ydir = y
						currentPos.x += xdir
						currentPos.y += ydir
						break Loop
					}
				}
			}
		// If it wasn't a direction it must be a letter so add the letter to the solution and continue in the direction
		default:
			solution += currentChar
			currentPos.x += xdir
			currentPos.y += ydir
		}
		numbersteps++
	}
	return solution, numbersteps
}

// Functions builds a 2D grid from a slice of strings
func createGrid(lines []string) ([][]string, coord, int, int) {
	// grid will hold all the characters in the maze
	grid := make([][]string, len(lines))
	currentPos := coord{0, 0}
	ymax := len(lines)
	xmax := len(lines)
	// Set each grid space to a char from the input lines
	for y, line := range lines {
		grid[y] = make([]string, len(lines))
		for x, char := range line {
			grid[y][x] = string(char)
			// Get the starting coordinates by finding | in the first row
			if y == 0 && string(char) == "|" {
				currentPos.x = x
			}
		}
	}
	return grid, currentPos, ymax, xmax
}

// Function builds a slice of strings
// Each string is a row in the pipeline maze
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
