package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1(puzzleInput string) int {
	cleanedInput := strings.Replace(puzzleInput, "\r", "", -1)
	chunks := strings.Split(cleanedInput, "\n\n")

	maxSum := 0
	for _, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		sum := 0
		for _, line := range lines {
			intValue, _ := strconv.Atoi(line)
			sum += intValue
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func day1() {
	puzzleRead, _ := os.ReadFile("Input/day1.txt")
	puzzleString := string(puzzleRead)

	start := time.Now()
	result := part1(puzzleString)
	duration := time.Since(start)
	fmt.Printf("Part 1 result is: %v, computed in: %v \n", result, duration)
}
