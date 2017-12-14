package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	security := buildSecuritySystem()
	part1, _ := severity(security, 0)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", delayEscape(security))
}

// Function builds the security system from the input file
// security system is a map of int to int
// The key is the depth of the security and the value is the range of the security system at that depth
func buildSecuritySystem() map[int]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	regValues := regexp.MustCompile("[0-9]+")
	security := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		values := regValues.FindAllString(line, -1)
		depth, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}

		rangeVal, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}

		security[depth] = rangeVal

	}
	return security
}

// Function calculates the severity of being caught at time picoseconds
// returns an int: The severity of being caught at the current time
// returns a bool: If we were caught in depth 0
func severity(security map[int]int, time int) (int, bool) {
	sum := 0
	caughtZero := false

	for depth, rangeVal := range security {
		// We add the time to the depth level
		// This tells us how much time has passed at each depth
		// If we leave at time 0 then at depth 0 the time will be 0 + 0, at depth 1 the time will be 1 + 0 and so on.
		// We use remainder division on the total time using  (range - 1 * 2) to determine where the security scanner is at that time
		// We multiply by two because the security scanner travels both directions in the range and it travels at 1 range per 1 time
		// If the remainder is 0 then the scanner is in range place 0 and we are caught
		// If caught we add the severity calculation to the total sum
		if (depth+time)%((rangeVal-1)*2) == 0 {
			sum += depth * rangeVal
			// If the depth 0 is caught return true that we were caught in depth 0
			if depth == 0 {
				caughtZero = true
			}
		}
	}
	return sum, caughtZero
}

// Function searches for a time when we can escape without being caught
func delayEscape(security map[int]int) int {
	// Keep incrementing the delay value and checking the severity
	for delay := 0; ; delay++ {
		// Check the severity of being caught at the current time
		sever, caughtZero := severity(security, delay)
		// If the severity of being caught is 0 then we check if we were caught in depthZero
		// We have to check depth zero because depth 0 * range  = 0 even though we may have been caught.
		if sever == 0 && !caughtZero {
			// If we have escaped successfully return the time we can escape
			return delay
		}
	}
	return 0
}
