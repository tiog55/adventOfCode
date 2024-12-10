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
	file, err := os.Open("2024/9/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int
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

	currentFreeSpaceCandidate := ""
	currentFreeSpaceCandidateCount := 0
	for rightIndex := len(formattedRow) - 1; rightIndex > 0; rightIndex-- {
		if currentFreeSpaceCandidate == "" && formattedRow[rightIndex] != "." {
			currentFreeSpaceCandidate = formattedRow[rightIndex]
			currentFreeSpaceCandidateCount++
		} else if currentFreeSpaceCandidate == formattedRow[rightIndex] {
			currentFreeSpaceCandidateCount++
		} else if currentFreeSpaceCandidate != "" {
			freeSpaceStartIndex := -1
			freeSpace := 0
			for leftIndex := 0; leftIndex < rightIndex; leftIndex++ {
				if freeSpaceStartIndex == -1 && formattedRow[leftIndex] == "." {
					freeSpaceStartIndex = leftIndex
					freeSpace++
				} else if formattedRow[leftIndex] == "." {
					freeSpace++
				} else if freeSpaceStartIndex != -1 && formattedRow[leftIndex] != "." {
					if freeSpace >= currentFreeSpaceCandidateCount {
						//shift
						for i := 0; i < currentFreeSpaceCandidateCount; i++ {
							formattedRow[freeSpaceStartIndex+i] = currentFreeSpaceCandidate
							formattedRow[rightIndex+1+i] = "."
						}
						break
					} else {
						freeSpaceStartIndex = -1
						freeSpace = 0
					}
				}
			}
			if formattedRow[rightIndex] != "." {
				currentFreeSpaceCandidate = formattedRow[rightIndex]
				currentFreeSpaceCandidateCount = 1
			} else {
				currentFreeSpaceCandidate = ""
				currentFreeSpaceCandidateCount = 0
			}
		}
	}

	for rowNumber, entry := range formattedRow {
		if entry != "." {
			result += rowNumber * toInt(entry)
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
