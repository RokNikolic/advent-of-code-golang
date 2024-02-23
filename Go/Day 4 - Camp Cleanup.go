package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rangeOverlap(range1 []int, range2 []int) []int {
	var overlapSlice []int
	for i := range1[0]; i <= range1[1]; i++ {
		if i >= range2[0] && i <= range2[1] {
			overlapSlice = append(overlapSlice, i)
		}
	}
	return overlapSlice
}

func stringToIntRange(inputRange []string) []int {
	var outputRange []int
	for _, element := range inputRange {
		intElement, _ := strconv.Atoi(element)
		outputRange = append(outputRange, intElement)
	}
	return outputRange
}

func day4PartBoth(puzzleInput string, part int) int {
	lines := strings.Split(puzzleInput, "\n")
	totalSum := 0
	for _, line := range lines {
		ranges := strings.Split(line, ",")
		firstRange, secondRange := strings.Split(ranges[0], "-"), strings.Split(ranges[1], "-")
		intRange1, intRange2 := stringToIntRange(firstRange), stringToIntRange(secondRange)

		firstLength, secondLength := (intRange1[1]+1)-intRange1[0], (intRange2[1]+1)-intRange2[0]
		smallestRange := min(firstLength, secondLength)
		overlap := rangeOverlap(intRange1, intRange2)

		if smallestRange == len(overlap) && part == 1 {
			totalSum += 1
		}
		if len(overlap) > 0 && part == 2 {
			totalSum += 1
		}
	}
	return totalSum
}

func day4() {
	puzzleRead, _ := os.ReadFile("Input/day4.txt")
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 4 Part 1 result is: %v\n", day4PartBoth(puzzleCleaned, 1))
	fmt.Printf("Day 4 Part 2 result is: %v\n", day4PartBoth(puzzleCleaned, 2))
}
