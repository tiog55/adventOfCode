package main

import (
	"bufio"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Button struct {
	xAddition, yAddition int
	costs                int
}

type Price struct {
	x, y int
}

type Machine struct {
	price   Price
	buttonA Button
	buttonB Button
}

func main() {
	overallStart := time.Now()
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var parsedInput []Machine
	machine := Machine{}
	for i, row := range input {
		firstSplit := strings.Split(row, ": ")
		if strings.HasPrefix(row, "Button A") {
			secondSplit := strings.Split(firstSplit[1], ", ")
			machine.buttonA = Button{costs: 3, xAddition: toInt(strings.Split(secondSplit[0], "+")[1]), yAddition: toInt(strings.Split(secondSplit[1], "+")[1])}
		} else if strings.HasPrefix(row, "Button B") {
			secondSplit := strings.Split(firstSplit[1], ", ")
			machine.buttonB = Button{costs: 1, xAddition: toInt(strings.Split(secondSplit[0], "+")[1]), yAddition: toInt(strings.Split(secondSplit[1], "+")[1])}
		} else if strings.HasPrefix(row, "Prize") {
			secondSplit := strings.Split(firstSplit[1], ", ")
			machine.price = Price{x: toInt(strings.Split(secondSplit[0], "=")[1]), y: toInt(strings.Split(secondSplit[1], "=")[1])}
		} else if len(row) == 0 {
			parsedInput = append(parsedInput, machine)
			machine = Machine{}
		}
		if i == len(input)-1 {
			parsedInput = append(parsedInput, machine)
			machine = Machine{}
		}
	}

	for i, m := range parsedInput {
		prize := isSolvable(m)
		price := m.buttonA.costs*prize.x + m.buttonB.costs*prize.y
		if prize.x == 0 && prize.y == 0 {
			fmt.Println("Machine", i, "is not solvable")
		} else {
			fmt.Println("Machine", i, "is solvable by", prize.x, prize.y, "with a price of ", price)
		}
		result += price
	}
	fmt.Println("The result ist ", result)
	elapsed := time.Since(overallStart)
	log.Println("Overall took", elapsed)
}

func isSolvable(machine Machine) (price Price) {
	buttonA := machine.buttonA
	buttonB := machine.buttonB
	prize := machine.price
	A := mat.NewDense(2, 2, []float64{float64(buttonA.xAddition), float64(buttonB.xAddition), float64(buttonA.yAddition), float64(buttonB.yAddition)})
	b := mat.NewVecDense(2, []float64{float64(prize.x), float64(prize.y)})

	var sol mat.VecDense
	if err := sol.SolveVec(A, b); err != nil {
		fmt.Println(err)
		return
	}
	solutionData := sol.RawVector().Data
	buttonAPushes := solutionData[0]
	buttonBPushes := solutionData[1]
	if buttonA.xAddition*int(buttonAPushes)+buttonB.xAddition*int(buttonBPushes) == prize.x &&
		buttonA.yAddition*int(buttonAPushes)+buttonB.yAddition*int(buttonBPushes) == prize.y {
		price = Price{x: int(buttonAPushes), y: int(buttonBPushes)}
	} else if buttonA.xAddition*int(math.Floor(buttonAPushes))+buttonB.xAddition*int(math.Ceil(buttonBPushes)) == prize.x &&
		buttonA.yAddition*int(math.Floor(buttonAPushes))+buttonB.yAddition*int(math.Ceil(buttonBPushes)) == prize.y {
		price = Price{x: int(math.Floor(buttonAPushes)), y: int(math.Ceil(buttonBPushes))}
	} else if buttonA.xAddition*int(math.Ceil(buttonAPushes))+buttonB.xAddition*int(math.Floor(buttonBPushes)) == prize.x &&
		buttonA.yAddition*int(math.Ceil(buttonAPushes))+buttonB.yAddition*int(math.Floor(buttonBPushes)) == prize.y {
		price = Price{x: int(math.Ceil(buttonAPushes)), y: int(math.Floor(buttonBPushes))}
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

func toString(d int) string {
	return strconv.FormatInt(int64(d), 10)
}
