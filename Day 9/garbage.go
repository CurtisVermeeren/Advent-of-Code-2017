package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	// Read the file into a slice of bytes
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	// Convert the slice of bytes to a string
	stream := string(bytes)
	// Create a regular expression "!." using MustCompile
	// Any ! followed by any character is removed (replaced by "") in the string
	stream = regexp.MustCompile("!.").ReplaceAllString(stream, "")

	// Garbage will track the amount of garbage removed
	garbage := 0

	// Use a Regexp to get all "Garbage"
	arr := regexp.MustCompile("<.*?>").FindAllStringSubmatch(stream, -1)
	for index := range arr {
		// subtract 2 for the < and > surrounding the garbage characters.
		garbage += len(arr[index][0]) - 2
	}

	// Find all "Garbage" any characters (if any) between a < and >
	// remove all garbage by replacing with "" in the string
	stream = regexp.MustCompile("<.*?>").ReplaceAllString(stream, "")
	// Break the string into a slice of substrings
	// Because we use "" as the seperation string each substring is 1 character
	chars := strings.Split(stream, "")
	// Track the current level (number of nested groups) and the current socre
	level, score := 1, 0
	for _, val := range chars {
		switch val {
		case "{":
			// Every time a { is encountered increment another level of groups and add it to the score
			score += level
			level++
		case "}":
			// Every time a { is encounterd decement the level of groups
			level--
		}
	}
	fmt.Println(score)
	fmt.Println(garbage)
}
