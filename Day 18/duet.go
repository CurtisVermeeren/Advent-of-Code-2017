package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	inputs := getInput()
	solvePart1(inputs)

	queue1 := make(chan int, 100)
	queue2 := make(chan int, 100)

	waitGroup.Add(1)
	go solvePart2(inputs, 0, queue1, queue2)
	waitGroup.Add(1)
	go solvePart2(inputs, 1, queue1, queue2)
	waitGroup.Wait()
}

// Function creates a slice of string instructions given an input file
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

// Function solves part 1 given a slice of instructions
func solvePart1(inputs []string) {
	registers := make(map[string]int)
	soundFrequency := 0
	// Perform instructions while they remain
	for i := 0; i < len(inputs); i++ {
		// Split the instruction into parts for parsing
		input := inputs[i]
		parts := strings.Split(input, " ")
		// Use a switch to detect the instruction type
		switch parts[0] {
		// snd X : Play the sound of value X
		case "snd":
			if isNumber(parts[1]) {
				fmt.Println(toNumber(parts[1]))
				soundFrequency = toNumber(parts[1])
			} else {
				soundFrequency = registers[parts[1]]
			}
		// set X Y : Set register X to value Y
		case "set":
			if isNumber(parts[2]) {
				registers[parts[1]] = toNumber(parts[2])
			} else {
				registers[parts[1]] = registers[parts[2]]
			}
		// add X Y : Increase register X by value Y
		case "add":
			if isNumber(parts[2]) {
				registers[parts[1]] += toNumber(parts[2])
			} else {
				registers[parts[1]] += registers[parts[2]]
			}
		// mul X Y : Set X to value of X * Y
		case "mul":
			if isNumber(parts[2]) {
				registers[parts[1]] = registers[parts[1]] * toNumber(parts[2])
			} else {
				registers[parts[1]] = registers[parts[1]] * registers[parts[2]]
			}
		// mod X Y : Set X to value of X % Y
		case "mod":
			if isNumber(parts[2]) {
				registers[parts[1]] %= toNumber(parts[2])
			} else {
				registers[parts[1]] %= registers[parts[2]]
			}
		// rcv X : get the last played frequency if X != 0
		case "rcv":
			if isNumber(parts[1]) {
				if toNumber(parts[1]) != 0 {
					fmt.Println(soundFrequency)
				}
			} else {
				if registers[parts[1]] != 0 {
					fmt.Println(soundFrequency)

				}
			}
			// If the last sound Frequency is output we can exit as we have the answer to part 1
			i = len(inputs)
		// jgz X Y : jumps the instructions forward Y steps if X > 0
		case "jgz":
			if isNumber(parts[1]) {
				if toNumber(parts[1]) > 0 {
					if isNumber(parts[2]) {
						i += toNumber(parts[2]) - 1
					} else {
						i += registers[parts[2]] - 1
					}
				}
			} else {
				if registers[parts[1]] > 0 {
					if isNumber(parts[2]) {
						i += toNumber(parts[2]) - 1
					} else {
						i += registers[parts[2]] - 1
					}
				}
			}
		}
	}
}

// Function is a program in part 2
// has an id of 0 or 1, and two channels used for communication between programs
func solvePart2(inputs []string, id int, program1Chan chan int, program2Chan chan int) {
	registers := make(map[string]int)

	// Track the number of sends by each program
	numberSend := 0

	// Set register p equal to the id according to the problem
	registers["p"] = id
	fmt.Printf("Program %v starting\n", id)

	// Perform instructions while they remain
	for i := 0; i < len(inputs); i++ {
		// Split the instruction into parts for parsing
		input := inputs[i]
		parts := strings.Split(input, " ")
		// Use a switch to detect the instruction type
		switch parts[0] {

		// snd X : Sends the value of X to the other programs
		case "snd":
			// Get the value of X
			value := 0
			if isNumber(parts[1]) {
				value = toNumber(parts[1])
			} else {
				value = registers[parts[1]]
			}
			// Pass it to the other programs channel
			if id == 0 {
				program1Chan <- value
			} else {
				program2Chan <- value
			}
			// Increment the number of sends
			numberSend++

		// set X Y : Set register X to value Y
		case "set":
			if isNumber(parts[2]) {
				registers[parts[1]] = toNumber(parts[2])
			} else {
				registers[parts[1]] = registers[parts[2]]
			}
		// add X Y : Increase register X by value Y
		case "add":
			if isNumber(parts[2]) {
				registers[parts[1]] += toNumber(parts[2])
			} else {
				registers[parts[1]] += registers[parts[2]]
			}
		// mul X Y : Set X to value of X * Y
		case "mul":
			if isNumber(parts[2]) {
				registers[parts[1]] = registers[parts[1]] * toNumber(parts[2])
			} else {
				registers[parts[1]] = registers[parts[1]] * registers[parts[2]]
			}
		// mod X Y : Set X to value of X % Y
		case "mod":
			if isNumber(parts[2]) {
				registers[parts[1]] %= toNumber(parts[2])
			} else {
				registers[parts[1]] %= registers[parts[2]]
			}
		// rcv X : Receives the next value and stores it in register X
		case "rcv":
			var receive int
			var timeout bool
			// Attempt to get the value from the other program
			if id == 0 {
				receive, timeout = getFromChannelWithTimeout(program2Chan)
			} else if id == 1 {
				receive, timeout = getFromChannelWithTimeout(program1Chan)
			}

			// If there was a timeout then there is a deadlock so end the program
			if timeout {
				// Print the number of values sent by this program in its operation
				fmt.Printf("Program %v send %v messages\n", id, numberSend)
				waitGroup.Done()
				return
			}
			// If no deadlock then set X to the value received from the other program
			registers[parts[1]] = receive

		// jgz X Y : jumps the instructions forward Y steps if X > 0
		case "jgz":
			if isNumber(parts[1]) {
				if toNumber(parts[1]) > 0 {
					if isNumber(parts[2]) {
						i += toNumber(parts[2]) - 1
					} else {
						i += registers[parts[2]] - 1
					}
				}
			} else {
				if registers[parts[1]] > 0 {
					if isNumber(parts[2]) {
						i += toNumber(parts[2]) - 1
					} else {
						i += registers[parts[2]] - 1
					}
				}
			}
		}

	}
	waitGroup.Done()
}

// Function converts a string to an integer with reusable error checking
func toNumber(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

// Returns true is str is number, false otherwise
func isNumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

// Function attempts to get a value from the channel
func getFromChannelWithTimeout(channel chan int) (int, bool) {
	var receive int
	timeout := make(chan bool, 1)
	// Set a timeout of 1 second
	// After this time send a timeout (true) to the timeout channel
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()
	// Select waits for 1 of 2 operations
	select {
	// Either we receive a value from the channel
	case receive = <-channel:
		return receive, false
	// Or we receive a timeout
	case <-timeout:
		fmt.Println("Timeout")
		return receive, true
	}

}
