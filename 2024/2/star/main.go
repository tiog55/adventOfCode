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
	file, err := os.Open("2024/2/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var iter int
	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isSafe(strings.Fields(scanner.Text()), false) {
			result++
		}
		iter++
	}

	fmt.Println(result, " of ", iter, " are valid.")
}

func isSafe(row []string, alreadyExcluded bool) bool {
	// direction == 1 increasing | direction == 0 decreasing | direction == -1 not set
	decidedDirection := -1
	lastEntry := -1
	for _, entry := range row {
		currentEntry := toInt(entry)
		// handle first entry
		if lastEntry == -1 {
			lastEntry = currentEntry
			continue
		}
		// handle direction decision
		if decidedDirection == -1 {
			decidedDirection = decideDirection(lastEntry, currentEntry)
		}
		if !logicallyIncreasing(lastEntry, currentEntry, decidedDirection) && !logicallyDecreasing(lastEntry, currentEntry, decidedDirection) {
			// if not already excluded, try by removing any entry
			if !alreadyExcluded {
				for i, _ := range row {
					if isSafe(removeIndexFromRow(row, i), true) {
						return true
					}
				}
			}
			return false
		}
		lastEntry = currentEntry
	}
	return true
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}

func decideDirection(lastEntry int, currentEntry int) int {
	if logicallyIncreasing(lastEntry, currentEntry, 1) {
		return 1
	} else {
		return 0
	}
}

func logicallyIncreasing(lastEntry int, currentEntry int, direction int) bool {
	return lastEntry < currentEntry && lastEntry+3 >= currentEntry && direction == 1
}

func logicallyDecreasing(lastEntry int, currentEntry int, direction int) bool {
	return lastEntry > currentEntry && lastEntry-3 <= currentEntry && direction == 0
}

func removeIndexFromRow(row []string, index int) []string {
	var newRow []string
	for i, entry := range row {
		if i != index {
			newRow = append(newRow, entry)
		}
	}
	return newRow
}
