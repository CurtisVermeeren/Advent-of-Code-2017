package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type program struct {
	id          string
	connections []string
}

var visited map[string]bool

func main() {
	visited = make(map[string]bool)
	tree := getInput()
	visitConnections(tree, tree["0"])
	fmt.Println("Part 1:", len(visited))
	visited = make(map[string]bool)
	part2 := numberOfGroups(tree)
	fmt.Println("Part 2", part2)
}

func getInput() map[string]*program {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	regID := regexp.MustCompile("[0-9]+")

	allPrograms := make(map[string]*program)

	for scanner.Scan() {
		line := scanner.Text()
		names := regID.FindAllString(line, -1)
		p := &program{id: names[0], connections: names[1:]}
		allPrograms[names[0]] = p
	}
	return allPrograms
}

func visitConnections(tree map[string]*program, currentProgram *program) {
	// If already visited skip this program
	if visited[currentProgram.id] == true {
		return
	}

	visited[currentProgram.id] = true

	for _, connect := range currentProgram.connections {
		visitConnections(tree, tree[connect])
	}
}

func numberOfGroups(tree map[string]*program) int {
	numGroups := 0
	for key, program := range tree {
		if visited[key] == true {
			continue
		}
		visitConnections(tree, program)
		numGroups++
	}
	return numGroups
}
