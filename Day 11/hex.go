package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pmcxs/hexgrid"
)

func main() {
	inputs := getInput()

	q := 0
	r := 0
	maxDistance := 0

	// Use the hexgrid package to set the origin hex tile
	origin := hexgrid.NewHex(0, 0)
	// Get direction from the childs path (inputs)
	// Use Axial coordinates to increment q and r based on the direction
	// Axial coordinates have two axes q,r that are 60° or 120° apart.
	for _, direction := range inputs {
		switch direction {
		case "n":
			q--
		case "ne":
			r++
			q--
		case "se":
			r++
		case "s":
			q++
		case "sw":
			r--
			q++
		case "nw":
			r--
		}
		// Add a new hex along the path
		currentHex := hexgrid.NewHex(q, r)
		// Each time calculate distance from origin and current hex. Keep track of the maximum distance for part 2
		currentDistance := hexgrid.HexDistance(origin, currentHex)
		if currentDistance > maxDistance {
			maxDistance = currentDistance
		}
	}
	// The child is the final position
	hexagonB := hexgrid.NewHex(q, r)
	// The distance between two hexes is the length of the line between them.
	// This is the distance between the origin and the child
	distance := hexgrid.HexDistance(origin, hexagonB)
	fmt.Println("Part 1:", distance)
	fmt.Println("Part 2:", maxDistance)
}

// Function builds a slice of input directions from a comma separated file
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
