package year2022

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day1part1(puzzleInput string) int {
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
	return allSums[0]
}

func day1part2(puzzleInput string) int {
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
	lastFew := allSums[0:3]
	for _, value := range lastFew {
		maxSum += value
	}

	return maxSum
}

func Day1() {
	puzzleRead, err := os.ReadFile("2022/Input/day1.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		puzzleString := string(puzzleRead)
		puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

		fmt.Printf("Part 1 result is: %v\n", day1part1(puzzleCleaned))
		fmt.Printf("Part 2 result is: %v\n", day1part2(puzzleCleaned))
	}
}
