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
	file, err := os.Open("2024/12/star/input.txt")
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
		areaName := input[allPoints[0].y][allPoints[0].x]
		sides := findSides(allPoints, areaName)
		area := len(allPointsByArea[currentArea])
		result += sides * area
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

func findSides(allAreaPoints []Point, areaName string) (sides int) {
	sides = 4
	if len(allAreaPoints) <= 2 {
		return
	}

	sides = 0
	//Build connecting map
	for _, point := range allAreaPoints {
		hitUpLeftCorner := 1
		hitUpRightCorner := 1
		hitDownRightCorner := 1
		hitDownLeftCorner := 1
		for _, innerPoint := range allAreaPoints {
			if point == innerPoint {
				continue
			}
			if innerPoint.x == point.x && innerPoint.y == point.y-1 { //up
				hitUpLeftCorner = 0
				if !slices.Contains(allAreaPoints, Point{x: point.x + 1, y: point.y}) && !slices.Contains(allAreaPoints, Point{x: point.x + 1, y: point.y - 1}) {
					hitUpRightCorner = 0
				}
			}
			if innerPoint.x == point.x+1 && innerPoint.y == point.y { //right
				hitUpRightCorner = 0
				if !slices.Contains(allAreaPoints, Point{x: point.x, y: point.y + 1}) && !slices.Contains(allAreaPoints, Point{x: point.x + 1, y: point.y + 1}) {
					hitDownRightCorner = 0
				}
			}
			if innerPoint.x == point.x && innerPoint.y == point.y+1 { //down
				hitDownRightCorner = 0
				if !slices.Contains(allAreaPoints, Point{x: point.x - 1, y: point.y}) && !slices.Contains(allAreaPoints, Point{x: point.x - 1, y: point.y + 1}) {
					hitDownLeftCorner = 0
				}
			}
			if innerPoint.x == point.x-1 && innerPoint.y == point.y { //left
				hitDownLeftCorner = 0
				if !slices.Contains(allAreaPoints, Point{x: point.x, y: point.y - 1}) && !slices.Contains(allAreaPoints, Point{x: point.x - 1, y: point.y - 1}) {
					hitUpLeftCorner = 0
				}
			}
		}
		sides += hitUpLeftCorner + hitUpRightCorner + hitDownRightCorner + hitDownLeftCorner
	}

	return
}

func toString(d int) string {
	return strconv.FormatInt(int64(d), 10)
}
