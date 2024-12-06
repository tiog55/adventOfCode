package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/6/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var input [][]string
	var startX int
	var startY int
	var direction string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		input = append(input, row)
		if slices.Index(row, "^") != -1 {
			{
				direction = "up"
				startX = slices.Index(row, "^")
				startY = len(input) - 1
			}
		} else if slices.Index(row, ">") != -1 {
			{
				direction = "right"
				startX = slices.Index(row, ">")
				startY = len(input) - 1
			}
		} else if slices.Index(row, "<") != -1 {
			{
				direction = "left"
				startX = slices.Index(row, "<")
				startY = len(input) - 1
			}
		} else if slices.Index(row, "v") != -1 {
			{
				direction = "down"
				startX = slices.Index(row, "v")
				startY = len(input) - 1
			}
		}
	}
	fmt.Println("StartingPosition: ", startX, startY)
	fmt.Println("StartingDirection: ", direction)
	findExit(input, direction, startX, startY)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[j][i] == "X" {
				input[j][i] = "#" // mark as visited
				if !isExitable(input, direction, startX, startY, 0) {
					result++
				}
				input[j][i] = "X" // reset
			}
		}
	}

	fmt.Println("The result ist ", result)
}

func isExitable(input [][]string, direction string, currentX, currentY, currentSteps int) (exitable bool) {
	if currentSteps > len(input)*len(input[0]) {
		return false
	}
	currentSteps++
	if direction == "up" {
		if currentY == 0 {
			exitable = true
			return
		} else if currentY > 0 && input[currentY-1][currentX] == "#" {
			exitable = isExitable(input, "right", currentX, currentY, currentSteps)
		} else {
			currentY--
			exitable = isExitable(input, "up", currentX, currentY, currentSteps)
		}
	} else if direction == "right" {
		if currentX == len(input[0])-1 {
			exitable = true
			return
		} else if currentX < len(input[0])-1 && input[currentY][currentX+1] == "#" {
			exitable = isExitable(input, "down", currentX, currentY, currentSteps)
		} else {
			currentX++
			exitable = isExitable(input, "right", currentX, currentY, currentSteps)
		}

	} else if direction == "down" {
		if currentY == len(input)-1 {
			exitable = true
			return
		} else if currentY < len(input)-1 && input[currentY+1][currentX] == "#" {
			exitable = isExitable(input, "left", currentX, currentY, currentSteps)
		} else {
			currentY++
			exitable = isExitable(input, "down", currentX, currentY, currentSteps)
		}
	} else if direction == "left" {
		if currentX == 0 {
			exitable = true
			return
		} else if currentX > 0 && input[currentY][currentX-1] == "#" {
			exitable = isExitable(input, "up", currentX, currentY, currentSteps)
		} else {
			currentX--
			exitable = isExitable(input, "left", currentX, currentY, currentSteps)
		}
	} else {
		panic("Invalid direction")
	}

	return
}

func findExit(input [][]string, direction string, currentX, currentY int) (steps int) {
	if direction == "up" {
		if currentY == 0 {
			input[currentY][currentX] = "X"
			steps++
			return
		} else if currentY > 0 && input[currentY-1][currentX] == "#" {
			steps += findExit(input, "right", currentX, currentY)
		} else {
			if input[currentY-1][currentX] != "X" {
				steps++
			}
			input[currentY][currentX] = "X"
			currentY--
			steps += findExit(input, "up", currentX, currentY)
		}
	} else if direction == "right" {
		if currentX == len(input[0])-1 {
			input[currentY][currentX] = "X"
			steps++
			return
		} else if currentX < len(input[0])-1 && input[currentY][currentX+1] == "#" {
			steps += findExit(input, "down", currentX, currentY)
		} else {
			if input[currentY][currentX+1] != "X" {
				steps++
			}
			input[currentY][currentX] = "X"
			currentX++
			steps += findExit(input, "right", currentX, currentY)
		}

	} else if direction == "down" {
		if currentY == len(input)-1 {
			input[currentY][currentX] = "X"
			steps++
			return
		} else if currentY < len(input)-1 && input[currentY+1][currentX] == "#" {
			steps += findExit(input, "left", currentX, currentY)
		} else {
			if input[currentY+1][currentX] != "X" {
				steps++
			}
			input[currentY][currentX] = "X"
			currentY++
			steps += findExit(input, "down", currentX, currentY)
		}
	} else if direction == "left" {
		if currentX == 0 {
			input[currentY][currentX] = "X"
			steps++
			return
		} else if currentX > 0 && input[currentY][currentX-1] == "#" {
			steps += findExit(input, "up", currentX, currentY)
		} else {
			if input[currentY][currentX-1] != "X" {
				steps++
			}
			input[currentY][currentX] = "X"
			currentX--
			steps += findExit(input, "left", currentX, currentY)
		}
	} else {
		panic("Invalid direction")
	}

	return
}
