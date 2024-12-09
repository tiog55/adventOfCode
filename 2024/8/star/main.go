package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Antenna struct {
	x, y int
}

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/8/star/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var field [][]string
	antennas := make(map[string][]Antenna)
	var index int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		field = append(field, row)
		fmt.Println(row)
		for i, cell := range row {
			if cell != "." {
				antennas[cell] = append(antennas[cell], Antenna{x: i, y: index})
			}
		}
		index++
	}

	findings := make([]Antenna, 0)
	for _, currentAntennas := range antennas {
		for _, antenna := range currentAntennas {
			r, f, updatedField := findAntinodes(field, antenna, currentAntennas, findings)
			result += r
			findings = f
			field = updatedField
		}
	}
	fmt.Println("The result ist ", result)
}

func findAntinodes(field [][]string, currentAntenna Antenna, antennas []Antenna, previousFindings []Antenna) (hits int, findings []Antenna, updatedField [][]string) {
	currentX := currentAntenna.x
	currentY := currentAntenna.y
	updatedField = field
	findings = previousFindings
	for _, antenna := range antennas {
		diagX := int(math.Abs(float64(currentX - antenna.x)))
		diagY := int(math.Abs(float64(currentY - antenna.y)))
		// left up
		if antenna.x < currentX && antenna.y < currentY {
			antinodeX := currentX - diagX
			antinodeY := currentY - diagY
			inField := true
			for inField {
				if antinodeX >= 0 && antinodeY >= 0 {
					if field[antinodeY][antinodeX] != "#" && !alreadyFound(findings, antinodeX, antinodeY) {
						findings = append(findings, Antenna{x: antinodeX, y: antinodeY})
						hits++
					}
					if field[antinodeY][antinodeX] == "." {
						updatedField[antinodeY][antinodeX] = "#"
					}
				} else {
					inField = false
				}
				antinodeX = antinodeX - diagX
				antinodeY = antinodeY - diagY
			}
		} else if antenna.x > currentX && antenna.y < currentY {
			antinodeX := currentX + diagX
			antinodeY := currentY - diagY
			inField := true
			for inField {
				if antinodeX < len(field[0]) && antinodeY >= 0 {
					if field[antinodeY][antinodeX] != "#" && !alreadyFound(findings, antinodeX, antinodeY) {
						findings = append(findings, Antenna{x: antinodeX, y: antinodeY})
						hits++
					}
					if field[antinodeY][antinodeX] == "." {
						updatedField[antinodeY][antinodeX] = "#"
					}
				} else {
					inField = false
				}
				antinodeX = antinodeX + diagX
				antinodeY = antinodeY - diagY
			}
		} else if antenna.x > currentX && antenna.y > currentY {
			antinodeX := currentX + diagX
			antinodeY := currentY + diagY
			inField := true
			for inField {
				if antinodeX < len(field[0]) && antinodeY < len(field) {
					if field[antinodeY][antinodeX] != "#" && !alreadyFound(findings, antinodeX, antinodeY) {
						findings = append(findings, Antenna{x: antinodeX, y: antinodeY})
						hits++
					}
					if field[antinodeY][antinodeX] == "." {
						updatedField[antinodeY][antinodeX] = "#"
					}
				} else {
					inField = false
				}
				antinodeX = antinodeX + diagX
				antinodeY = antinodeY + diagY
			}
		} else if antenna.x < currentX && antenna.y > currentY {
			antinodeX := currentX - diagX
			antinodeY := currentY + diagY
			inField := true
			for inField {
				if antinodeX >= 0 && antinodeY < len(field) {
					if field[antinodeY][antinodeX] != "#" && !alreadyFound(findings, antinodeX, antinodeY) {
						findings = append(findings, Antenna{x: antinodeX, y: antinodeY})
						hits++
					}
					if field[antinodeY][antinodeX] == "." {
						updatedField[antinodeY][antinodeX] = "#"
					}
				} else {
					inField = false
				}
				antinodeX = antinodeX - diagX
				antinodeY = antinodeY + diagY
			}
		}
	}

	return
}

func alreadyFound(findings []Antenna, x, y int) (found bool) {
	for _, f := range findings {
		if f.x == x && f.y == y {
			return true
		}
	}
	return
}
