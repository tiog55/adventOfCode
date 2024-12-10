package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/10/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var input [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		input = append(input, strings.Split(row, ""))
	}

	for y, row := range input {
		for x, entry := range row {
			if entry == "0" {
				result += foundTrailHeads(input, x, y, 1)
			}
		}
	}

	fmt.Println("The result ist ", result)
}

func foundTrailHeads(input [][]string, x, y, nextNumber int) (findings int) {
	if nextNumber == 10 {
		//fmt.Println("Found ", x, y)
		return 1
	}

	if y > 0 && toInt(input[y-1][x]) == nextNumber { //up
		findings += foundTrailHeads(input, x, y-1, nextNumber+1)
	}
	if x < len(input[0])-1 && toInt(input[y][x+1]) == nextNumber { //right
		findings += foundTrailHeads(input, x+1, y, nextNumber+1)
	}
	if y < len(input)-1 && toInt(input[y+1][x]) == nextNumber { //down
		findings += foundTrailHeads(input, x, y+1, nextNumber+1)
	}
	if x > 0 && toInt(input[y][x-1]) == nextNumber { //left
		findings += foundTrailHeads(input, x-1, y, nextNumber+1)
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
