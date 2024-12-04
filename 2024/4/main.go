package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/4/input.txt")
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
			if string(v) == "X" {
				fmt.Println(x, y, "", "[M A S]")
				result += isFound(horizontal, x, y, "", []string{"M", "A", "S"}, "")
			}
		}
	}

	fmt.Println("The result ist ", result)
}

func isFound(input []string, x int, y int, direction string, target []string, iteration string) (hits int) {
	iteration += "   "
	if len(target) == 0 {
		return 1
	}
	// up
	if (direction == "" || direction == "up") && y > 0 && string(input[y-1][x]) == target[0] {
		fmt.Println(iteration, x, y-1, "up", target[1:])
		hits += isFound(input, x, y-1, "up", target[1:], iteration)
	}

	// up right
	if (direction == "" || direction == "up right") && x < len(input[x])-1 && y > 0 && string(input[y-1][x+1]) == target[0] {
		fmt.Println(iteration, x+1, y-1, "up right", target[1:])
		hits += isFound(input, x+1, y-1, "up right", target[1:], iteration)
	}

	// right
	if (direction == "" || direction == "right") && x < len(input[x])-1 && string(input[y][x+1]) == target[0] {
		fmt.Println(iteration, x+1, y, "right", target[1:])
		hits += isFound(input, x+1, y, "right", target[1:], iteration)
	}

	// right down
	if (direction == "" || direction == "right down") && x < len(input[x])-1 && y < len(input)-1 && string(input[y+1][x+1]) == target[0] {
		fmt.Println(iteration, x+1, y+1, "right down", target[1:])
		hits += isFound(input, x+1, y+1, "right down", target[1:], iteration)
	}

	// down
	if (direction == "" || direction == "down") && y < len(input)-1 && string(input[y+1][x]) == target[0] {
		fmt.Println(iteration, x, y+1, "down", target[1:])
		hits += isFound(input, x, y+1, "down", target[1:], iteration)
	}

	// down left
	if (direction == "" || direction == "down left") && x > 0 && y < len(input)-1 && string(input[y+1][x-1]) == target[0] {
		fmt.Println(iteration, x-1, y+1, "down left", target[1:])
		hits += isFound(input, x-1, y+1, "down left", target[1:], iteration)
	}

	// left
	if (direction == "" || direction == "left") && x > 0 && string(input[y][x-1]) == target[0] {
		fmt.Println(iteration, x-1, y, "left", target[1:])
		hits += isFound(input, x-1, y, "left", target[1:], iteration)
	}

	// left up
	if (direction == "" || direction == "left up") && x > 0 && y > 0 && string(input[y-1][x-1]) == target[0] {
		fmt.Println(iteration, x-1, y-1, "left up", target[1:])
		hits += isFound(input, x-1, y-1, "left up", target[1:], iteration)
	}

	if hits == 0 && len(target) <= 1 {
		fmt.Println("No hit")
	}

	return
}
