package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/4/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var horizontal []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		horizontal = append(horizontal, row)
	}
	for _, value := range horizontal {
		for _, v := range value {
			fmt.Print(string(v), " ")
		}
		fmt.Println("")
	}
	for y, value := range horizontal {
		for x, v := range value {
			if string(v) == "A" {
				result += isFound(horizontal, x, y)
			}
		}
	}

	fmt.Println("The result ist ", result)
}

func isFound(input []string, x int, y int) (hits int) {
	rDiagonal := make([]string, 0)
	lDiagonal := make([]string, 0)

	// up right
	if x < len(input[x])-1 && y > 0 {
		lDiagonal = append(lDiagonal, string(input[y-1][x+1]))
	}

	// down left
	if x > 0 && y < len(input)-1 {
		lDiagonal = append(lDiagonal, string(input[y+1][x-1]))
	}

	// right down
	if x < len(input[x])-1 && y < len(input)-1 {
		rDiagonal = append(rDiagonal, string(input[y+1][x+1]))
	}

	// left up
	if x > 0 && y > 0 {
		rDiagonal = append(rDiagonal, string(input[y-1][x-1]))
	}

	if len(rDiagonal) == 2 && len(lDiagonal) == 2 {
		if slices.Contains(rDiagonal, "M") && slices.Contains(rDiagonal, "S") {
			if slices.Contains(lDiagonal, "M") && slices.Contains(lDiagonal, "S") {
				hits++
			}
		}
	}

	return
}
