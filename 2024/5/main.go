package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("2024/5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var rules []string
	dependsOn := make(map[string][]string)
	var prints []string
	rulesInput := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		if rulesInput && row != "" {
			order := strings.Split(row, "|")
			dependsOn[order[1]] = append(dependsOn[order[1]], order[0])
			rules = append(rules, row)
		} else if rulesInput && strings.Replace(row, " ", "", -1) == "" {
			rulesInput = false
		} else if !rulesInput {
			prints = append(prints, row)
		}

	}

	for _, print := range prints {
		singlePrint := strings.Split(print, ",")
		validPrint := true
		for i, printDigit := range singlePrint {
			if !validPrint {
				break
			}
			dependentPrints := dependsOn[printDigit]
			restOfPrints := singlePrint[i+1:]
			for _, dependent := range dependentPrints {
				if slices.Contains(restOfPrints, dependent) {
					validPrint = false
					break
				}
			}
		}
		if validPrint {
			result += toInt(singlePrint[len(singlePrint)/2])
		}
	}

	fmt.Println("The result ist ", result)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
