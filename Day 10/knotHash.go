package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	// --------------------------------------------------
	rope := makeRope(256)
	currentPosition := 0
	skipSize := 0
	// Get all length values from file
	lengths := getLengths()

	// Perfrom 1 round of knot hashing
	for _, length := range lengths {
		// Get the start and end positions of the sub rope
		start := (currentPosition) % len(rope)
		end := (start + length - 1) % len(rope)
		// Reverse the sub rope
		reverse(rope, start, end, length)
		// Update the current position and skipSize
		currentPosition += length + skipSize
		skipSize++
	}

	fmt.Println("Part 1:", rope[0]*rope[1])

	// Part 2
	// --------------------------------------------------
	hashRope := makeRope(256)
	// Get all length values from file
	lengths = getLengthsPart2()
	currentPosition = 0
	skipSize = 0
	numberRounds := 64
	// Perfrom 64 rounds of the knot hash
	for i := 0; i < numberRounds; i++ {
		for _, length := range lengths {
			start := (currentPosition) % len(hashRope)
			end := (currentPosition + length - 1) % len(hashRope)
			reverse(hashRope, start, end, length)
			currentPosition += length + skipSize
			skipSize++
		}
	}
	hash := convertToHash(hashRope)
	hex := convertToHex(hash)
	fmt.Println("Part 2:", hex)
}

// Function builds an incrementing list of integers of length size
func makeRope(size int) []int {
	rope := make([]int, size)
	for i := range rope {
		rope[i] = i
	}
	return rope
}

// Function reads the input file and creates a slice of integers with all lengths to be moved
func getLengths() []int {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Remove the newline and divide into a slice of strings
	arr := strings.Split(strings.Trim(string(bytes), "\n"), ",")
	lengths := make([]int, len(arr))
	// Convert each string into an integer
	for i, v := range arr {
		lengths[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
	}
	return lengths
}

func getLengthsPart2() []int {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Append the reuired numbers to the []byte
	lengths := append(file, 17, 31, 73, 47, 23)
	intLengths := make([]int, len(lengths))
	// Type cast the bytes to integers
	for i, element := range lengths {
		intLengths[i] = int(element)
	}
	return intLengths
}

// Function reverses a sub rope of length starting at from and ending at to
func reverse(rope []int, from int, to int, length int) {
	// Make a copy of rope to preserve values
	copyMoves := make([]int, len(rope))
	copy(copyMoves, rope)
	// We need to make length/2 number of moves to reverse a list of length
	// Each move swap 2 indexes of length
	for i := 0; i < length/2; i++ {
		// Calculate the indexes
		// index1 moves forward 1 each time, index2 moves backwards 1 each time
		// We use the mod function to move around the slice in a circular way
		index1 := mod(from+i, len(rope))
		index2 := mod(to-i, len(rope))
		// Swap values in index1 with index2 and vice versa
		rope[index1], rope[index2] = copyMoves[index2], copyMoves[index1]
	}
}

// Function helps find indexes around the circular slice
func mod(number int, lengthRope int) int {
	// If the number is negative then we have to move back to the end of the slice from the beginning
	if number < 0 {
		return mod(lengthRope+number, lengthRope)
	}
	// We use remainder division to loop around the slice in a circular way when
	return number % lengthRope
}

// Convert the sparse hash of numbers from 0 to 255 into a dense hash of only 16 numbers
// We'll use the bitwise XOR `^` operator to create the sparse hash
func convertToHash(list []int) []int {
	denseHash := make([]int, len(list)/16)
	for i := 0; i < len(list); i += 16 {
		// Use ^ to combine 16 values
		for j := 0; j < 16; j++ {
			denseHash[i/16] ^= list[i+j]
		}
	}
	return denseHash
}

// Function converts the row to hex
func convertToHex(list []int) string {
	hexString := ""
	for _, element := range list {
		// Print each value in the dense hash in hexadecimal notation
		// %x indicates hexadeciaml
		// zerofill to two places
		hexString += fmt.Sprintf("%.02x", element)
	}
	return hexString
}
