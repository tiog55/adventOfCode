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
	file, err := os.Open("2024/9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int64
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	row := strings.Split(scanner.Text(), "")
	rowNumber := 0
	formattedRow := make([]string, 0)
	for index, entry := range row {
		dataEntry := index%2 == 0
		var newEntry string
		if dataEntry {
			newEntry = toString(rowNumber)
			rowNumber++
		} else {
			newEntry = "."
		}
		for i := 0; i < toInt(entry); i++ {
			formattedRow = append(formattedRow, newEntry)
		}
	}

	leftIndex := 0
	rightIndex := len(formattedRow) - 1
	for leftIndex < rightIndex {
		if formattedRow[leftIndex] != "." {
			leftIndex++
		} else if formattedRow[rightIndex] == "." {
			rightIndex--
		} else {
			leftEntry := formattedRow[leftIndex]
			rightEntry := formattedRow[rightIndex]
			formattedRow[leftIndex] = rightEntry
			formattedRow[rightIndex] = leftEntry
		}
	}

	for rowNumber, entry := range formattedRow {
		if entry != "." {
			result += int64(rowNumber * toInt(entry))
		}
	}

	fmt.Println("The result ist ", result)
}

func toString(d int) string {
	return strconv.FormatInt(int64(d), 10)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
