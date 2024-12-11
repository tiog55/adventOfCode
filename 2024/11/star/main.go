package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	overallStart := time.Now()
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/11/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		for _, entry := range strings.Split(row, " ") {
			input = append(input, toInt(entry))
		}
	}

	stones := make(map[int]int)
	for _, entry := range input {
		if _, ok := stones[entry]; !ok {
			stones[entry] = 1
		} else {
			stones[entry]++
		}
	}

	for i := 0; i < 75; i++ {
		var keys []int
		for k := range stones {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		addToStones := make(map[int]int)
		for _, key := range keys {
			for _, singleStoneResult := range observe(key) {
				if _, ok := addToStones[singleStoneResult]; ok {
					addToStones[singleStoneResult] += stones[key]
				} else {
					addToStones[singleStoneResult] = stones[key]
				}
			}
		}
		for key, value := range addToStones {
			if _, ok := stones[value]; ok {
				stones[key] += value
			} else {
				stones[key] = value
			}
		}
		stones = addToStones
	}

	for _, count := range stones {
		result += count
	}

	fmt.Println("The result ist ", result)
	elapsed := time.Since(overallStart)
	log.Println("Overall took", elapsed)
}

func observe(input int) (output []int) {
	output = make([]int, 0)
	if input == 0 { //zero
		output = append(output, 1)
	} else if len(toString(input))%2 == 0 { //even number
		stringEntry := toString(input)
		middle := len(stringEntry) / 2
		firstEntry := stringEntry[:middle]
		secondEntry := stringEntry[middle:]
		output = append(output, toInt(firstEntry))
		output = append(output, toInt(secondEntry))
	} else { // everything else
		output = append(output, input*2024)
	}
	return
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}

func toString(d int) string {
	return strconv.FormatInt(int64(d), 10)
}
