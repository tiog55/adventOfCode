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
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/3/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()

	}

	// At first we use everything before first `don't()`
	StartingRegex, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\).*?(don't\(\))`)
	startingResult := StartingRegex.FindString(input)
	result += multiplyEachPair(startingResult)

	// Afterward, we only process everything between 'do()' and 'don't()'
	EnabledRegex, _ := regexp.Compile(`do\(\).*?mul\(([0-9]+),([0-9]+)\).*?don't\(\)`)
	enabledResult := EnabledRegex.FindAllString(input, -1)
	for _, enabled := range enabledResult {
		result += multiplyEachPair(enabled)
	}
	// Finally, we need to include last where no 'don't' is included
	FinaleRegex, _ := regexp.Compile(`do\(\).*?mul\(([0-9]+),([0-9]+)\).*?`)
	doIndices := FinaleRegex.FindAllIndex([]byte(input), -1)
	finalInput := input[doIndices[len(doIndices)-1][0]:]
	result += multiplyEachPair(finalInput)

	fmt.Println("The result ist ", result)
}

func multiplyEachPair(input string) int {
	regex, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	startingResults := regex.FindAllString(input, -1)
	return multiplyString(startingResults)
}

func multiplyString(input []string) int {
	var result int
	for _, mul := range input {
		innerRegex, _ := regexp.Compile(`([0-9]+)`)
		values := innerRegex.FindAllString(mul, -1)
		first := -1
		for _, value := range values {
			if first == -1 {
				first = toInt(value)
				continue
			}
			second := toInt(value)
			result += first * second
		}
	}
	return result
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
