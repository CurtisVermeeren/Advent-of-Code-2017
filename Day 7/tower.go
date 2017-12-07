package main

import (
	"bufio"
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

func solvePart1(programs map[string]*program) (*program, error) {
	for _, prog := range programs {
		for prog.parent != nil {
			prog = prog.parent
		}
		return prog, nil
	}
	panic("empty map!")
}

func solvePart2(in map[string]*program) int {
	var fixedWeight int
	for n, _ := solvePart1(in); n != nil && len(n.children) > 2; {
		w := n.children[0].totalWeight
		w2 := n.children[1].totalWeight
		if n.children[2].totalWeight == w2 {
			w = w2
		}
		var next *program
		for _, n := range n.children {
			if n.totalWeight != w {
				next = n
				fixedWeight = n.weight - (n.totalWeight - w)
			}
		}
		n = next
	}
	return fixedWeight
}

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

	regNames := regexp.MustCompile("[a-z]+")
	regWeight := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		line := scanner.Text()
		names := regNames.FindAllString(line, -1)
		weight, err := strconv.Atoi(regWeight.FindString(line))
		if err != nil {
			log.Fatal(err)
		}
		prog := &program{name: names[0], weight: weight, totalWeight: weight}
		in[names[0]] = prog
		children[prog] = names[1:]
	}

	for prog, kids := range children {
		prog.children = make([]*program, len(kids))
		for index, child := range kids {
			prog.children[index] = in[child]
		}
	}

	for _, prog := range in {
		for _, child := range prog.children {
			child.parent = prog
		}
	}

	//calculate total weights
	for _, n := range in {
		for k := n.parent; k != nil; k = k.parent {
			k.totalWeight += n.weight
		}
	}
	return in
}
