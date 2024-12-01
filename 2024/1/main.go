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
	file, err := os.Open("2024/1/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var iter int
	var first []string
	occurencesInSecond := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		first = append(first, words[0])
		val := occurencesInSecond[words[1]]
		occurencesInSecond[words[1]] = val + 1
		iter++
	}

	var result int
	slices.Sort(first)

	for i := 0; i < iter; i++ {
		if val, ok := occurencesInSecond[first[i]]; ok {
			f := toInt(first[i])
			singleResult := f * val
			fmt.Println("first: ", f, " occurences: ", val, " result: ", singleResult)
			result += singleResult
		}
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
