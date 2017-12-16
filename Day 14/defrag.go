package main

import "fmt"

func main() {
	input := "uugsqrei"
	grid := make([][]bool, 128)
	visited := 0
	for i := 0; i < 128; i++ {
		grid[i] = make([]bool, 128)
		visited += processGridRow(grid[i], fmt.Sprintf("%s-%d", input, i))
	}
	fmt.Println("Part 1:", visited)
	fmt.Println("Part 2:", countRegions(grid))
}

func processGridRow(row []bool, lengths string) int {
	binaryString := knotHash(lengths)
	count := 0
	for charNum, char := range binaryString {
		if string(char) == "1" {
			row[charNum] = true
			count++
		}
	}
	return count
}

func knotHash(input string) string {
	lengths := []byte(input)
	hashRope := makeRope(256)
	lengths = append(lengths, 17, 31, 73, 47, 23)
	currentPosition := 0
	skipSize := 0
	numberRounds := 64
	// Perfrom 64 rounds of the knot hash
	for i := 0; i < numberRounds; i++ {
		for _, byteLength := range lengths {
			length := int(byteLength)
			start := (currentPosition) % len(hashRope)
			end := (currentPosition + length - 1) % len(hashRope)
			reverse(hashRope, start, end, length)
			currentPosition += length + skipSize
			skipSize++
		}
	}
	hash := convertToHash(hashRope)
	binary := convertToBinary(hash)
	return binary
}

// Function builds an incrementing list of integers of length size
func makeRope(size int) []byte {
	rope := make([]byte, size)
	for i := range rope {
		rope[i] = byte(i)
	}
	return rope
}

// Function reverses a sub rope of length starting at from and ending at to
func reverse(rope []byte, from int, to int, length int) {
	// Make a copy of rope to preserve values
	copyMoves := make([]byte, len(rope))
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
func convertToHash(list []byte) []byte {
	denseHash := make([]byte, len(list)/16)
	for i := 0; i < len(list); i += 16 {
		// Use ^ to combine 16 values
		for j := 0; j < 16; j++ {
			denseHash[i/16] ^= list[i+j]
		}
	}
	return denseHash
}

// Function converts the row to hex
func convertToBinary(list []byte) string {
	binaryString := ""
	for _, element := range list {
		// Print each value in the dense hash in hexadecimal notation
		// %x indicates hexadeciaml
		// zerofill to two places
		binaryString += fmt.Sprintf("%08b", element)
	}
	return binaryString
}

func countRegions(grid [][]bool) int {
	count := 0
	for i, row := range grid {
		for j, used := range row {
			if used {
				visit(i, j, grid)
				count++
			}
		}
	}
	return count
}

func visit(i, j int, grid [][]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || !grid[i][j] {
		return
	}
	grid[i][j] = false
	visit(i+1, j, grid)
	visit(i-1, j, grid)
	visit(i, j+1, grid)
	visit(i, j-1, grid)
}
