package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file2, err2 := os.Open("input.txt")
	if err != nil {
		log.Fatal(err2)
	}
	defer file2.Close()

	fmt.Println(solvePart1(file))
	fmt.Println(solvePart2(file2))
}

func solvePart1(f *os.File) int {
	// Open input file to read
	scanner := bufio.NewScanner(f)

	totalValid := 0

	// Read the file line by line and convert each passphrase into []string using strings.Fields
	for scanner.Scan() {
		passPhrase := strings.Fields(scanner.Text())
		// Use a map to track words in the passphrase
		check := make(map[string]bool)
		isVald := true
		// Attempt to add each passphrase word to the map. If it is found then the passphrase is invalid
		for _, value := range passPhrase {
			if _, ok := check[value]; !ok {
				check[value] = true
			} else {
				isVald = false
				break
			}
		}
		// If each key in the map was unique then the passphrase is valid.
		if isVald {
			totalValid++
		}
	}
	return totalValid
}

func solvePart2(f *os.File) int {
	// Open input file to read
	scanner := bufio.NewScanner(f)

	totalValid := 0

	for scanner.Scan() {
		passPhrase := strings.Fields(scanner.Text())
		// Use a map to track words in the passphrase
		check := make(map[string]bool)
		isVald := true
		// Attempt to add each passphrase word to the map. If it is found then the passphrase is invalid
		for _, value := range passPhrase {
			value = sortStringByCharacter(value)
			if _, ok := check[value]; !ok {
				check[value] = true
			} else {
				isVald = false
				break
			}
		}
		// If each key in the map was unique then the passphrase is valid.
		if isVald {
			totalValid++
		}
	}
	return totalValid
}

// Function converts a string to a slice of runes
func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// Function sorts a string by character returning a sorted string
func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
