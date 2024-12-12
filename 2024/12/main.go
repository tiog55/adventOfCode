package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	file, err := os.Open("2024/12/input.txt")
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

	allPointsByArea := make(map[string][]Point, 0)
	alreadyTraversed := make([]Point, 0)
	for y, row := range input {
		for x, _ := range row {
			currentArea := input[y][x]
			currentPoint := Point{x: x, y: y}
			if !slices.Contains(alreadyTraversed, currentPoint) {
				allConnectedPoints := findConnectedPoints(input, currentPoint, currentArea, &alreadyTraversed)
				allPointsByArea[currentArea+toString(x)+toString(y)] = allConnectedPoints
			}
		}
	}

	for currentArea, allPoints := range allPointsByArea {
		perimeter := findPerimeter(allPoints)
		area := len(allPointsByArea[currentArea])
		result += perimeter * area
	}

	fmt.Println("The result ist ", result)
	elapsed := time.Since(overallStart)
	log.Println("Overall took", elapsed)
}

func findConnectedPoints(input [][]string, point Point, area string, alreadyTraversed *[]Point) []Point {
	connectedPoints := make([]Point, 0)
	connectedPoints = append(connectedPoints, point)
	*alreadyTraversed = append(*alreadyTraversed, point)
	if point.y > 0 && input[point.y-1][point.x] == area && !slices.Contains(*alreadyTraversed, Point{x: point.x, y: point.y - 1}) { //up
		connectedPoints = append(connectedPoints, findConnectedPoints(input, Point{x: point.x, y: point.y - 1}, area, alreadyTraversed)...)
	}
	if point.x < len(input[0])-1 && input[point.y][point.x+1] == area && !slices.Contains(*alreadyTraversed, Point{x: point.x + 1, y: point.y}) { //right
		connectedPoints = append(connectedPoints, findConnectedPoints(input, Point{x: point.x + 1, y: point.y}, area, alreadyTraversed)...)
	}
	if point.y < len(input)-1 && input[point.y+1][point.x] == area && !slices.Contains(*alreadyTraversed, Point{x: point.x, y: point.y + 1}) { //down
		connectedPoints = append(connectedPoints, findConnectedPoints(input, Point{x: point.x, y: point.y + 1}, area, alreadyTraversed)...)
	}
	if point.x > 0 && input[point.y][point.x-1] == area && !slices.Contains(*alreadyTraversed, Point{x: point.x - 1, y: point.y}) { //left
		connectedPoints = append(connectedPoints, findConnectedPoints(input, Point{x: point.x - 1, y: point.y}, area, alreadyTraversed)...)
	}

	return connectedPoints

}

func findPerimeter(allAreaPoints []Point) (perimeter int) {
	for _, point := range allAreaPoints {
		startingPerimeter := 4
		if slices.Contains(allAreaPoints, Point{x: point.x, y: point.y - 1}) { //up
			startingPerimeter--
		}
		if slices.Contains(allAreaPoints, Point{x: point.x + 1, y: point.y}) { //right
			startingPerimeter--
		}
		if slices.Contains(allAreaPoints, Point{x: point.x, y: point.y + 1}) { //down
			startingPerimeter--
		}
		if slices.Contains(allAreaPoints, Point{x: point.x - 1, y: point.y}) { //left
			startingPerimeter--
		}
		perimeter += startingPerimeter
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

func toString(d int) string {
	return strconv.FormatInt(int64(d), 10)
}
