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
	// If the node has less than 3 children then we have reached the endpoint.
	// The 2 children cannot be unbalanced as the problem states exactly 1 node has the wrong weight.
	// If 1 child of 2 is wrong the other may be wrong too either could be higher or lower to balance
	for n, _ := solvePart1(in); n != nil && len(n.children) > 2; {
		w := n.children[0].totalWeight
		w2 := n.children[1].totalWeight
		// If child 3 is the same as child 2 then we know that the weight (w2) is correct so we make w = w2
		// If the weight of child 3 is not the same as the weight of child 2 then 1 of child 2 or 3 is wrong so we know the weight of child 1 (w) is correct
		if n.children[2].totalWeight == w2 {
			w = w2
		}
		var next *program
		// Loop through all children of the node
		for _, n := range n.children {
			// If the weight of the child is NOT the correct determined weight
			if n.totalWeight != w {
				// We know this node is wrong so we continue deeper in the tower to fix it
				next = n
				// Calculate the weight of the wrong child by the difference between it's current total weight and the correct total weight
				fixedWeight = n.weight - (n.totalWeight - w)
			}
		}
		// If no total weights were wrong in this layer then next == nil, so next = n == nil and the loop ends because n != nil computes false
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
		// For every node
		// Get the node's parent, add the weight of the node to the parent, then the parent become the node, repeat with the new node
		// Do this for every node and all weights will be added
		for k := n.parent; k != nil; k = k.parent {
			k.totalWeight += n.weight
		}
	}
	return in
}
