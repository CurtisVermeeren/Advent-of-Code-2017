package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pattern [][]bool

// A rule has an input pattern which results in the output patter
type rule struct {
	input  pattern
	output pattern
}

func main() {
	// Get the input lines and build the rulebook of patterns
	lines := getInput()
	rulebook := buildRulebook(lines)

	// Use a 3D slice, to make a slice of patterns
	art := [][][]bool{{{false, true, false}, {false, false, true}, {true, true, true}}}

	// Perform 5 iterations of the art generation
	for i := 0; i < 5; i++ {
		art = next(rulebook, art)
	}

	part1 := count(art)
	fmt.Println("Part 1:", part1)

	// Perform 18 iterations of the art generation
	for i := 5; i < 18; i++ {
		art = next(rulebook, art)
	}
	part2 := count(art)
	fmt.Println("Part 2:", part2)
}

// Function counts the number of true values in the art. Returns that value
func count(art [][][]bool) int {
	count := 0
	// Iterate through each value and check if true
	for _, p := range art {
		for _, row := range p {
			for _, b := range row {
				if b {
					count++
				}
			}
		}
	}
	return count
}

func next(rulebook map[string][][]bool, in [][][]bool) (out [][][]bool) {
	for _, pat := range in {
		// If the length of a pattern is 4 (4x4) divide it into 4 parts of 2 (2x2)
		if len(pat) == 4 {
			// Build the sub pattern
			p1 := make([][]bool, 2)
			p1[0] = pat[0][:2]
			p1[1] = pat[1][:2]
			// Get the output for the part from the rulebook
			p1 = rulebook[encode(p1)]

			p2 := make([][]bool, 2)
			p2[0] = pat[0][2:]
			p2[1] = pat[1][2:]
			p2 = rulebook[encode(p2)]

			p3 := make([][]bool, 2)
			p3[0] = pat[2][:2]
			p3[1] = pat[3][:2]
			p3 = rulebook[encode(p3)]

			p4 := make([][]bool, 2)
			p4[0] = pat[2][2:]
			p4[1] = pat[3][2:]
			p4 = rulebook[encode(p4)]

			// 4 2x2 have a 6x6 output
			combined := make([][]bool, 6)
			for i := 0; i < 3; i++ {
				combined[i] = p1[i]
				combined[i] = append(combined[i], p2[i]...)
			}
			for i := 3; i < 6; i++ {
				combined[i] = p3[i-3]
				combined[i] = append(combined[i], p4[i-3]...)
			}
			out = append(out, combined)
			// If the length of a pattern is 6 (6x6) divide it into 9 parts of 2 (2x2)
		} else if len(pat) == 6 {
			for i := 0; i < 3; i++ {
				// Build the sub pattern
				p := make([][]bool, 2)
				p[0] = pat[i*2][:2]
				p[1] = pat[i*2+1][:2]
				// Get the output from the rulebook and append to out
				// The 2x2 becomes a 3x3
				out = append(out, rulebook[encode(p)])

				p = make([][]bool, 2)
				p[0] = pat[i*2][2:4]
				p[1] = pat[i*2+1][2:4]
				out = append(out, rulebook[encode(p)])

				p = make([][]bool, 2)
				p[0] = pat[i*2][4:]
				p[1] = pat[i*2+1][4:]
				out = append(out, rulebook[encode(p)])
			}
		} else {
			// If the length of the pattern is not 4 or 6 is is 3 so get the 4x4 from the rulebook and append it
			out = append(out, rulebook[encode(pat)])
		}
	}
	return
}

// Function "encodes" a 2D bool slice by converting it's values to a string
func encode(in [][]bool) (out string) {
	for _, row := range in {
		for _, b := range row {
			// Append a 1 or 0 if the value is true or false
			if b {
				out += "1"
			} else {
				out += "0"
			}
		}
	}
	return
}

// Function rotates a 2D slice n times
func rotate(in [][]bool, n int) (out [][]bool) {
	// If n == 0 keep the original
	if n == 0 {
		out = in

		// If n == 1 do 1 clockwise rotation
	} else if n == 1 {
		out = make([][]bool, len(in))
		for i := range in {
			out[i] = make([]bool, len(in))
		}
		// Rotate a 2x2
		if len(in) == 2 {
			out[0][0] = in[1][0]
			out[0][1] = in[0][0]
			out[1][0] = in[1][1]
			out[1][1] = in[0][1]
			// Rotate a 3x3
		} else {
			out[0][0] = in[2][0]
			out[0][1] = in[1][0]
			out[0][2] = in[0][0]
			out[1][0] = in[2][1]
			out[1][1] = in[1][1]
			out[1][2] = in[0][1]
			out[2][0] = in[2][2]
			out[2][1] = in[1][2]
			out[2][2] = in[0][2]
		}
		// If more than 1 clockwise rotation is wanted perform them recursively
	} else {
		out = rotate(rotate(in, n-1), 1)
	}
	return
}

// Function flips a 2D bool slice
func flip(in [][]bool) (out [][]bool) {
	out = make([][]bool, len(in))
	// Make a copy to be altered
	for i := range in {
		out[i] = make([]bool, len(in))
		copy(out[i], in[i])
	}
	// Flip all values of input
	for i := range in {
		out[i][0], out[i][len(out)-1] = out[i][len(out)-1], out[i][0]
	}
	return
}

// Function builds a map of rules from a slice of strings
func buildRulebook(lines []string) map[string][][]bool {
	rules := make([]rule, len(lines))
	for idx, line := range lines {
		// If the line has more than 20 characters it is a 3x3 going to a 4x4
		if len(line) > 20 {
			rules[idx] = newRule(3, 4)
			// Build the rules for the input and output
			parseRule(&rules[idx].input, line[:11])
			parseRule(&rules[idx].output, line[15:])
			// If the line has less than 20 characters it is a 2x2 going to a 3x3
		} else {
			// Build the rules for the input and output
			rules[idx] = newRule(2, 3)
			parseRule(&rules[idx].input, line[:5])
			parseRule(&rules[idx].output, line[9:])
		}
	}

	// Create a map to hold all possible inputs and their outputs
	rulebook := make(map[string][][]bool)
	for _, r := range rules {
		// Build all the combinations of input by rotating or flipping the input
		for i := 0; i < 4; i++ {
			// "encode" the input rule as a string to use as a unique key to it's output
			rulebook[encode(rotate(r.input, i))] = r.output
			rulebook[encode(flip(rotate(r.input, i)))] = r.output
		}
	}
	return rulebook
}

// Function converts a string into pattern object
// Finds '#' in the string and sets the corresponding [][]bool to true
func parseRule(r *pattern, s string) {
	i := 0
	for row := range *r {
		for col := range (*r)[row] {
			if s[i] == '#' {
				(*r)[row][col] = true
			}
			i++
		}
		i++
	}
}

// Function creates a new rule object
// allocates space for the input and output based on the size of the rule
func newRule(in, out int) rule {
	var r rule
	r.input = make([][]bool, in)
	for j := range r.input {
		r.input[j] = make([]bool, in)
	}
	r.output = make([][]bool, out)

	for j := range r.output {
		r.output[j] = make([]bool, out)
	}
	return r
}

// Function builds a slice of strings from an input file
func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	inputs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}
	return inputs
}
