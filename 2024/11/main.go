package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

func main() {
	overallStart := time.Now()
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var input []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		for _, entry := range strings.Split(row, " ") {
			input = append(input, toInt(entry))
		}
	}

	for i := 0; i < 25; i++ {
		input = observe(input)
	}

	result = len(input)
	fmt.Println("The result ist ", result)
	elapsed := time.Since(overallStart)
	log.Println("Overall took", elapsed)
}

func observe(input []int64) (output []int64) {
	output = make([]int64, 0)
	for _, entry := range input {
		if entry == 0 { //zero
			output = append(output, 1)
		} else if len(toString(entry))%2 == 0 { //even number
			stringEntry := toString(entry)
			middle := len(stringEntry) / 2
			firstEntry := stringEntry[:middle]
			secondEntry := stringEntry[middle:]
			output = append(output, toInt(firstEntry))
			output = append(output, toInt(secondEntry))
		} else { // everything else
			output = append(output, entry*2024)
		}
	}
	return
}

func toInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}

func toString(d int64) string {
	return strconv.FormatInt(d, 10)
}
