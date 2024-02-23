package main

import (
	"fmt"
	"os"
	"strings"
)

func stringIntersection(strings ...string) []int {
	firstString := strings[0]
	strings = strings[1:]
	workingSet := make(map[int32]bool)
	for _, char := range firstString {
		workingSet[char] = true
	}
	for _, workString := range strings {
		intersectionSet := make(map[int32]bool)
		for _, char := range workString {
			if _, found := workingSet[char]; found {
				intersectionSet[char] = true
			}
		}
		workingSet = intersectionSet
	}
	var outputSlice []int
	for key := range workingSet {
		outputSlice = append(outputSlice, int(key))
	}
	return outputSlice
}

func getPriority(asciiChar int) int {
	if asciiChar > 96 {
		return asciiChar - 96
	} else {
		return asciiChar - 64 + 26
	}
}

func day3Part1(puzzleInput string) int {
	lines := strings.Split(puzzleInput, "\n")
	totalSum := 0
	for _, line := range lines {
		middlePoint := len(line) / 2
		firstHalf, secondHalf := line[:middlePoint], line[middlePoint:]
		asciiLetters := stringIntersection(firstHalf, secondHalf)
		totalSum += getPriority(asciiLetters[0])
	}
	return totalSum
}

func day3Part2(puzzleInput string) int {
	lines := strings.Split(puzzleInput, "\n")
	totalSum := 0
	for i := 0; i < len(lines)/3; i++ {
		iter := i * 3
		first, second, third := lines[iter], lines[iter+1], lines[iter+2]
		asciiLetters := stringIntersection(first, second, third)
		totalSum += getPriority(asciiLetters[0])
	}
	return totalSum
}

func day3() {
	puzzleRead, _ := os.ReadFile("Input/day3.txt")
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 3 Part 1 result is: %v\n", day3Part1(puzzleCleaned))
	fmt.Printf("Day 3 Part 2 result is: %v\n", day3Part2(puzzleCleaned))
}
