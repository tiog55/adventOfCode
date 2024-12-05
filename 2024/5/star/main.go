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
	file, err := os.Open("2024/5/star/input.txt")
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
		if isInvalid(print, dependsOn) {
			isValid := false
			currentPrint := print
			for !isValid {
				correctedOrder, valid := buildCorrectOrder(currentPrint, dependsOn)
				if valid {
					isValid = true
					result += toInt(correctedOrder[len(correctedOrder)/2])
				} else {
					currentPrint = strings.Join(correctedOrder, ",")
				}
			}
		}
	}

	fmt.Println("The result ist ", result)
}

func isInvalid(print string, dependsOn map[string][]string) bool {
	singlePrint := strings.Split(print, ",")
	for i, printDigit := range singlePrint {
		dependentPrints := dependsOn[printDigit]
		if len(singlePrint) > i+1 {
			restOfPrints := singlePrint[i+1:]
			for _, dependent := range dependentPrints {
				if slices.Contains(restOfPrints, dependent) {
					return true
				}
			}
		}
	}
	return false
}

func buildCorrectOrder(print string, dependsOn map[string][]string) ([]string, bool) {
	var correctOrder []string
	singlePrint := strings.Split(print, ",")
	for i, printDigit := range singlePrint {
		dependentPrints := dependsOn[printDigit]
		if len(singlePrint) > i+1 {
			restOfPrints := singlePrint[i+1:]
			dependFound := false
			for _, dependent := range dependentPrints {
				if slices.Contains(restOfPrints, dependent) {
					dependFound = true
					correctOrder = append(correctOrder, dependent)
					correctOrder = append(correctOrder, printDigit)
					index := slices.Index(restOfPrints, dependent)
					for i, rest := range restOfPrints {
						if i == index {
							continue
						}
						correctOrder = append(correctOrder, rest)
					}
					return correctOrder, !isInvalid(strings.Join(correctOrder, ","), dependsOn)
				}
			}
			if !dependFound {
				correctOrder = append(correctOrder, printDigit)
			}
		}
	}
	return correctOrder, !isInvalid(strings.Join(correctOrder, ","), dependsOn)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
