package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		firstSpllit := strings.Split(row, ": ")
		SingleResult := toInt(firstSpllit[0])
		operations := []int{}
		for _, operation := range strings.Split(firstSpllit[1], " ") {
			operations = append(operations, toInt(operation))
		}
		possibleOperations := int(math.Pow(2, float64(len(operations)-1)))
		for i := 0; i < possibleOperations; i++ {
			operators := fmt.Sprintf("%b", i)
			if len(operators) < len(operations)-1 {
				operators = strings.Repeat("0", len(operations)-len(operators)-1) + operators
			}
			if solvable(operations, SingleResult, strings.Split(operators, "")) {
				result += SingleResult
				break
			}
		}
	}

	fmt.Println("The result ist ", result)
}

func solvable(input []int, result int, operations []string) bool {
	r := 0
	for i, operation := range operations {
		if r > result {
			break
		}
		if r == 0 {
			r = input[0]
		}
		if operation == "0" {
			r = r + input[i+1]
		} else if operation == "1" {
			r = r * input[i+1]
		} else {
			panic("Not implemented")
		}
	}
	return r == result
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
