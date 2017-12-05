package main

import (
	"fmt"
	"math"
)

func main() {
	target := 312051
	solvePart1(target)
	solvePart2(target)
}

// Function to solve part 1 of Advent of code day 3
func solvePart1(target int) {
	sidelength := lengthOfSideContaining(target)
	midpoints := midPointsForSideLength(sidelength)
	// stepsToRingFromCenter is the number of steps from each midpoint to the center value
	// because this is the most direct route we need to find the midpoint closest to the target value
	stepsToRingFromCenter := (sidelength - 1) / 2
	result := make([]float64, 4)
	// Calculate the number of steps from target to each midpoint
	// Add this to the stepsToRingFromCenter value to determine the total number of steps from target to midpoint to center
	for index, value := range midpoints {
		result[index] = float64(stepsToRingFromCenter) + math.Abs(float64(target)-float64(value))
	}
	// Find the smallest number of steps calculated above.
	// The smallest result is the minimum number of steps to center
	smallest := result[0]
	for _, value := range result {
		if value < smallest {
			smallest = value
		}
	}
	fmt.Println("The smallest number of steps to reach target: ", target, " is ", smallest)
}

// Function to solve part 2 of Advent of Code day 3
func solvePart2(target int) {
	// Build a grid with size to fit the target value.
	sideLength := lengthOfSideContaining(target) + 1
	grid := make([][]int, sideLength)
	for i := range grid {
		grid[i] = make([]int, sideLength)
	}
	// Start in the middle of the grid with value 1
	x := int(math.Ceil(float64(sideLength) / 2))
	y := int(math.Ceil(float64(sideLength) / 2))
	value := 1
	direction := "E"
	grid[x][y] = 1
	x++
	// While the value is less than the target continue to build the spiral grid
	for value <= target {
		// Calculate the value and set it
		// Check the direction for which way to build in the grid and increment x and y accordingly
		value = calculateValue(grid, x, y)
		grid[x][y] = value
		direction = calculateDirection(grid, x, y, direction)
		switch direction {
		case "E":
			x++
		case "N":
			y--
		case "W":
			x--
		case "S":
			y++
		}
	}
	// Print the first value larget than target
	fmt.Println("The next value after target: ", target, " is ", value)
}

// Calculate the value at (x,y) by adding together all values surrounding it
func calculateValue(grid [][]int, x, y int) int {
	value := grid[x+1][y] + grid[x+1][y-1] + grid[x][y-1] + grid[x-1][y-1] + grid[x-1][y] + grid[x-1][y+1] + grid[x][y+1] + grid[x+1][y+1]
	return value
}

// Calculate the direction of travel
func calculateDirection(grid [][]int, x, y int, dir string) string {
	// Find the corners of the build grid by looking for empty values
	// Gone as far east as needed (There are empty spaces to move up) we can go north.
	// Use the same concept for the other directions to create the spiral
	if dir == "E" && grid[x][y-1] == 0 {
		return "N"
	} else if dir == "N" && grid[x-1][y] == 0 {
		return "W"
	} else if dir == "W" && grid[x][y+1] == 0 {
		return "S"
	} else if dir == "S" && grid[x+1][y] == 0 {
		return "E"
	}
	// If not a corner carry on moving in the same direction
	return dir
}

// Function returns the length of the sides of the box containing the target value
/*
If the target is t in the following example

n n n n n
n n n n n
n n n n n
n n n n n
n n n t n

Then the value returned is 5 because t is in a row of 5 values.
*/
func lengthOfSideContaining(target int) int {
	target64 := float64(target)
	length := math.Ceil(math.Sqrt(target64))
	if int(length)%2 == 0 {
		return int(length + 1)
	}
	return int(length)
}

// Function returns a slice of all mindpoints in the rows/columns target could be in
/*
If the target is t then the input lenght is 5

n n a n n
n n n n n
a n n n a
n n n n n
n n a t n

The function will then return all 4 midpoints denoted a in the diagram
*/
func midPointsForSideLength(length int) []int {
	sides := make([]int, 4)
	highestOnSide := length * length
	offset := ((length - 1) / 2)
	for index, value := range sides {
		value = highestOnSide - (offset + (index * (length - 1)))
		sides[index] = value
	}
	return sides
}
