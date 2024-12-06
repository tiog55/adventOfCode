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
	file, err := os.Open("2024/6/input.txt")
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
		fmt.Println(row)
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

	result = findExit(input, direction, startX, startY)
	fmt.Println()
	for _, row := range input {
		fmt.Println(row)
	}
	fmt.Println("The result ist ", result)
}

func findExit(input [][]string, direction string, currentX, currentY int) (steps int) {
	if direction == "up" {
		if currentY == 0 {
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
