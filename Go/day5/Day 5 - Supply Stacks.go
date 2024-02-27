package day5

import (
	"fmt"
	"os"
	"slices"
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

func part1(puzzleInput string) int {
	chunks := strings.Split(puzzleInput, "\n\n")
	lines := strings.Split(chunks[0], "\n")
	lines = lines[:len(lines)-1]

	var stacks [][]string
	for _, line := range lines {
		var lineItems []string
		for i := 0; i <= len(lines[len(lines)-1])/4; i++ {
			line += "              "
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
	fmt.Println(finalStacks)
	return 0
}

func part2(puzzleInput string) int {
	return 0
}

func Day5() {
	puzzleRead, err := os.ReadFile("Input/day5.txt")
	if err != nil {
		fmt.Println(err)
	}
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 5 Part 1 result is: %v\n", part1(puzzleCleaned))
	fmt.Printf("Day 5 Part 2 result is: %v\n", part2(puzzleCleaned))
}
