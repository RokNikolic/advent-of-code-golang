package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day1PartBoth(puzzleInput string, part int) int {
	chunks := strings.Split(puzzleInput, "\n\n")

	var allSums []int
	for _, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		sum := 0
		for _, line := range lines {
			intValue, _ := strconv.Atoi(line)
			sum += intValue
		}
		allSums = append(allSums, sum)
	}
	slices.Sort(allSums)
	slices.Reverse(allSums)

	var maxSum int

	if part == 1 {
		maxSum = allSums[0]
	} else {
		lastFew := allSums[0:3]
		for _, value := range lastFew {
			maxSum += value
		}
	}

	return maxSum
}

func day1() {
	puzzleRead, _ := os.ReadFile("Input/day1.txt")
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 1 Part 1 result is: %v\n", day1PartBoth(puzzleCleaned, 1))
	fmt.Printf("Day 1 Part 2 result is: %v\n", day1PartBoth(puzzleCleaned, 2))
}
