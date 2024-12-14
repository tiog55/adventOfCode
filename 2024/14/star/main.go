package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Robot struct {
	x, y   int
	vX, vY int
	fX, fY int
}

func main() {
	overallStart := time.Now()
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/14/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	seconds := 10_000
	fieldX := 101
	fieldY := 103
	var input []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		robot := Robot{}
		firstSplit := strings.Split(scanner.Text(), " ")
		robotPosition := strings.Split(firstSplit[0][2:], ",")
		robotVelocity := strings.Split(firstSplit[1][2:], ",")
		robot.vX = toInt(robotVelocity[0])
		robot.vY = toInt(robotVelocity[1])
		robot.x = toInt(robotPosition[0])
		robot.y = toInt(robotPosition[1])
		input = append(input, robot)
	}

	findings := make(map[int][]string)

	for ; result < seconds; result++ {
		output := step(input, result, fieldX, fieldY)

		singleFinding := make(map[int]string)
		for y := 0; y < fieldY; y++ {
			singleFinding[y] = strings.Repeat("_", fieldX)
		}
		found := false
		for _, robot := range output {
			singleFinding[robot.fY] = singleFinding[robot.fY][:robot.fX] + "X" + singleFinding[robot.fY][robot.fX+1:]
			if strings.Contains(singleFinding[robot.fY], "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX") {
				found = true
			}
		}
		if found {
			break
		}

		findings[result] = maps.Values(singleFinding)
	}

	fmt.Println("The result ist ", result)
	elapsed := time.Since(overallStart)
	log.Println("Overall took", elapsed)
}

func step(input []Robot, second, fieldX, fieldY int) (output []Robot) {
	output = make([]Robot, 0)
	for _, robot := range input {
		xMoved := (robot.vX * second) % fieldX
		xMovedActual := xMoved + robot.x
		if xMovedActual != 0 {
			if xMoved < 0 { //left
				if xMovedActual > 0 { //fit into line
					robot.fX = xMovedActual
				} else { //overlaps
					robot.fX = fieldX + xMovedActual
				}
			} else { //right
				if xMovedActual < fieldX {
					robot.fX = xMovedActual
				} else { //overlaps
					robot.fX = xMovedActual - fieldX
				}
			}
		}

		yMoved := (robot.vY * second) % fieldY
		yMovedActual := yMoved + robot.y
		if yMovedActual != 0 {
			if yMoved < 0 { //up
				if yMovedActual > 0 { //fit into line
					robot.fY = yMovedActual
				} else { //overlaps
					robot.fY = fieldY + yMovedActual
				}
			} else { //down
				if yMovedActual < fieldY {
					robot.fY = yMovedActual
				} else { //overlaps
					robot.fY = yMovedActual - fieldY
				}
			}
		}
		output = append(output, robot)
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
