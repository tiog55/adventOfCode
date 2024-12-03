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
	file, err := os.Open("2024/3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		regex, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
		results := regex.FindAllString(row, -1)
		for _, mul := range results {
			innerRegex, _ := regexp.Compile(`([0-9]+)`)
			values := innerRegex.FindAllString(mul, -1)
			first := -1
			for _, value := range values {
				if first == -1 {
					first = toInt(value)
					continue
				}
				second := toInt(value)
				result += (first * second)
			}
			first = -1
		}
	}

	fmt.Println("The result ist ", result)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
