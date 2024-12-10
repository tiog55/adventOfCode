package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/10/input.txt")
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
				findings := foundTrailHeads(input, x, y, 1)
				uniqueFindings := make([]Point, 0)
				seen := make(map[Point]bool)
				for _, finding := range findings {
					if !seen[finding] {
						seen[finding] = true
						uniqueFindings = append(uniqueFindings, finding)
					}
				}
				fmt.Println("Unique ", uniqueFindings, " Count ", len(uniqueFindings))
				result += len(uniqueFindings)
			}
		}
	}

	fmt.Println("The result ist ", result)
}

func foundTrailHeads(input [][]string, x, y, nextNumber int) (findings []Point) {
	if y > 0 && toInt(input[y-1][x]) == nextNumber { //up
		if nextNumber == 9 {
			findings = append(findings, Point{x, y - 1})
		} else {
			findings = append(findings, foundTrailHeads(input, x, y-1, nextNumber+1)...)
		}
	}
	if x < len(input[0])-1 && toInt(input[y][x+1]) == nextNumber { //right
		if nextNumber == 9 {
			findings = append(findings, Point{x + 1, y})
		} else {
			findings = append(findings, foundTrailHeads(input, x+1, y, nextNumber+1)...)
		}
	}
	if y < len(input)-1 && toInt(input[y+1][x]) == nextNumber { //down
		if nextNumber == 9 {
			findings = append(findings, Point{x, y + 1})
		} else {
			findings = append(findings, foundTrailHeads(input, x, y+1, nextNumber+1)...)
		}
	}
	if x > 0 && toInt(input[y][x-1]) == nextNumber { //left
		if nextNumber == 9 {
			findings = append(findings, Point{x - 1, y})
		} else {
			findings = append(findings, foundTrailHeads(input, x-1, y, nextNumber+1)...)
		}
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
