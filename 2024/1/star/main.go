package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var iter int
	var first []string
	var second []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		first = append(first, words[0])
		second = append(second, words[1])
		iter++
	}

	var result int
	slices.Sort(first)
	slices.Sort(second)

	for i := 0; i < iter; i++ {
		f := toInt(first[i])
		s := toInt(second[i])
		singleResult := diff(f, s)
		fmt.Println("first: ", f, " second: ", s, " result: ", singleResult)
		result += singleResult
	}
	fmt.Println("There were ", iter, " entries in file")
	fmt.Println("The final result is ", result)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
