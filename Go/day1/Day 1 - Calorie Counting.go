package day1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func partBoth(puzzleInput string, part int) int {
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

func Day1() {
	puzzleRead, err := os.ReadFile("Input/day1.txt")
	if err != nil {
		fmt.Println(err)
	}
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 1 Part 1 result is: %v\n", partBoth(puzzleCleaned, 1))
	fmt.Printf("Day 1 Part 2 result is: %v\n", partBoth(puzzleCleaned, 2))
}
