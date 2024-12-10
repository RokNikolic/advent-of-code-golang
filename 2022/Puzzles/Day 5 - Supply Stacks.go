package year2022

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func moveStacks(stacks [][]string, instructions []int, oneByOne bool) [][]string {
	amount := instructions[0]
	source := instructions[1] - 1
	destination := instructions[2] - 1

	movingSlice := stacks[source][len(stacks[source])-amount:]
	if oneByOne {
		slices.Reverse(movingSlice)
	}
	stacks[source] = stacks[source][:len(stacks[source])-amount]
	stacks[destination] = append(stacks[destination], movingSlice...)
	return stacks
}

func day5part1(puzzleInput string) string {
	chunks := strings.Split(puzzleInput, "\n\n")
	pallets := chunks[0]
	instructions := chunks[1]

	palletLines := strings.Split(pallets, "\n")
	palletLines = palletLines[:len(palletLines)-1]
	var stacks [][]string
	for _, line := range palletLines {
		var lineItems []string
		for i := 0; i <= len(palletLines[len(palletLines)-1])/4; i++ {
			line += "                                        "
			lineItem := string(line[i*4+1])
			lineItems = append(lineItems, lineItem)
		}
		stacks = append(stacks, lineItems)
	}

	transposedStacks := transpose(stacks)
	var finalStacks [][]string
	for _, stack := range transposedStacks {
		var reversedStack []string
		for _, element := range stack {
			element = strings.Replace(element, " ", "0", -1)
			if element != "0" {
				reversedStack = append(reversedStack, element)
			}
		}
		slices.Reverse(reversedStack)
		finalStacks = append(finalStacks, reversedStack)
	}

	instructionLines := strings.Split(instructions, "\n")
	for _, line := range instructionLines {
		pattern := regexp.MustCompile("\\d+")
		foundNumbers := pattern.FindAllString(line, -1)

		var instructionNumbers []int
		for _, number := range foundNumbers {
			intNum, _ := strconv.Atoi(number)
			instructionNumbers = append(instructionNumbers, intNum)
		}

		finalStacks = moveStacks(finalStacks, instructionNumbers, true)
	}

	outputString := ""
	for _, stack := range finalStacks {
		topItem := stack[len(stack)-1]
		outputString += topItem
	}

	return outputString
}

func day5part2(puzzleInput string) string {
	chunks := strings.Split(puzzleInput, "\n\n")
	pallets := chunks[0]
	instructions := chunks[1]

	palletLines := strings.Split(pallets, "\n")
	palletLines = palletLines[:len(palletLines)-1]
	var stacks [][]string
	for _, line := range palletLines {
		var lineItems []string
		for i := 0; i <= len(palletLines[len(palletLines)-1])/4; i++ {
			line += "                                        "
			lineItem := string(line[i*4+1])
			lineItems = append(lineItems, lineItem)
		}
		stacks = append(stacks, lineItems)
	}

	transposedStacks := transpose(stacks)
	var finalStacks [][]string
	for _, stack := range transposedStacks {
		var reversedStack []string
		for _, element := range stack {
			element = strings.Replace(element, " ", "0", -1)
			if element != "0" {
				reversedStack = append(reversedStack, element)
			}
		}
		slices.Reverse(reversedStack)
		finalStacks = append(finalStacks, reversedStack)
	}

	instructionLines := strings.Split(instructions, "\n")
	for _, line := range instructionLines {
		pattern := regexp.MustCompile("\\d+")
		foundNumbers := pattern.FindAllString(line, -1)

		var instructionNumbers []int
		for _, number := range foundNumbers {
			intNum, _ := strconv.Atoi(number)
			instructionNumbers = append(instructionNumbers, intNum)
		}

		finalStacks = moveStacks(finalStacks, instructionNumbers, false)

	}

	outputString := ""
	for _, stack := range finalStacks {
		topItem := stack[len(stack)-1]
		outputString += topItem
	}

	return outputString
}

func Day5() {
	puzzleRead, err := os.ReadFile("2022/Input/day5.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		puzzleString := string(puzzleRead)
		puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

		fmt.Printf("Part 1 result is: %v\n", day5part1(puzzleCleaned))
		fmt.Printf("Part 2 result is: %v\n", day5part2(puzzleCleaned))
	}
}
