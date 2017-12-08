package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// A program type represents a node in the tower structure
type program struct {
	name        string
	children    []*program
	parent      *program
	weight      int
	totalWeight int
}

func main() {
	in := buildTower("input.txt")
	rootName, err := solvePart1(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rootName.name)
	fmt.Println(solvePart2(in))
}

// Solve part 1 by finding the root of the tree (The program with no parent)
func solvePart1(programs map[string]*program) (*program, error) {
	// Get a random program from the tree
	for _, prog := range programs {
		// While the program has a parent
		for prog.parent != nil {
			// Check the parent for a parent
			prog = prog.parent
		}
		return prog, nil
	}
	return nil, errors.New("Root not found")
}

func solvePart2(in map[string]*program) int {
	var fixedWeight int
	// Start at root get weights of each child
	for n, _ := solvePart1(in); n != nil && len(n.children) > 2; {
		w := n.children[0].totalWeight
		w2 := n.children[1].totalWeight
		// Check weights of children and balance if needed
		if n.children[2].totalWeight == w2 {
			w = w2
		}
		var next *program
		// Balance all children
		for _, n := range n.children {
			// If a subtree isn't balanced use it as the next to balance
			if n.totalWeight != w {
				next = n
				fixedWeight = n.weight - (n.totalWeight - w)
			}
		}
		n = next
	}
	return fixedWeight
}

// Function takes the filepath of the input and returns a map of all programs
func buildTower(filepath string) map[string]*program {
	// Open input file to read
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	in := make(map[string]*program)
	children := make(map[*program][]string)

	// Regular expressions to get the names of the program and children and the weight of the program
	regNames := regexp.MustCompile("[a-z]+")
	regWeight := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		line := scanner.Text()
		// FindAllString returns all strings that match the expression in a slice
		names := regNames.FindAllString(line, -1)
		// Find the integer string for the weight and convert to integer
		weight, err := strconv.Atoi(regWeight.FindString(line))
		if err != nil {
			log.Fatal(err)
		}
		// Create and add the program to the map of all programs.
		prog := &program{name: names[0], weight: weight, totalWeight: weight}
		in[names[0]] = prog
		// Map the children of the program
		children[prog] = names[1:]
	}

	// Add the children to each program
	for prog, kids := range children {
		prog.children = make([]*program, len(kids))
		for index, child := range kids {
			prog.children[index] = in[child]
		}
	}

	// For all programs get all of the children and add the program as a parent of it's children
	for _, prog := range in {
		for _, child := range prog.children {
			child.parent = prog
		}
	}

	// Calculate total weights
	for _, n := range in {
		for k := n.parent; k != nil; k = k.parent {
			k.totalWeight += n.weight
		}
	}
	return in
}
