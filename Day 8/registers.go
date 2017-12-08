package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var checkKey = ""
var maxValue = 0

func main() {

	registers, expressions := getData()
	calculateExpressions(registers, expressions)
	fmt.Println(findLargest(registers))
	fmt.Println(maxValue)
}

func getData() (map[string]int, []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var expressions []string

	registers := make(map[string]int)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		checkKey = line[0]

		// Add each value to register
		registers[line[0]] = 0

		// Add the expression to the expressions list
		expressions = append(expressions, scanner.Text())
	}

	return registers, expressions
}

func calculateExpressions(registers map[string]int, expressions []string) {
	for _, stringExp := range expressions {
		expression := strings.Fields(stringExp)

		left := registers[expression[4]]
		right, err := strconv.Atoi(expression[6])
		if err != nil {
			log.Fatal(err)
		}
		check := expression[5]
		checkedValue := validate(check, left, right)
		if checkedValue {
			if expression[1] == "inc" {
				incrementValue, err := strconv.Atoi(expression[2])
				if err != nil {
					log.Fatal(err)
				}
				registers[expression[0]] += incrementValue
				if registers[expression[0]] > maxValue {
					maxValue = registers[expression[0]]
				}
			} else {
				decrementValue, err := strconv.Atoi(expression[2])
				if err != nil {
					log.Fatal(err)
				}
				registers[expression[0]] -= decrementValue
				if registers[expression[0]] > maxValue {
					maxValue = registers[expression[0]]
				}
			}
		}
	}
}

func validate(check string, left, right int) bool {
	switch {
	case check == "<":
		return left < right
	case check == ">":
		return left > right
	case check == "==":
		return left == right
	case check == "<=":
		return left <= right
	case check == ">=":
		return left >= right
	case check == "!=":
		return left != right
	}
	return false
}

func findLargest(registers map[string]int) int {
	largest := registers[checkKey]
	for _, value := range registers {
		if value > largest {
			largest = value
		}
	}
	return largest
}
