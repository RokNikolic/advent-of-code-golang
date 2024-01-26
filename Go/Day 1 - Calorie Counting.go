package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func part1(puzzleInput string) int {
	cleanedInput := strings.Replace(puzzleInput, "\r", "", -1)
	chunks := strings.Split(cleanedInput, "\n\n")
	fmt.Println(chunks[0])

	return 0
}

func day1() {
	puzzleRead, _ := os.ReadFile("Input/day1.txt")
	puzzleString := string(puzzleRead)

	start := time.Now()
	result := part1(puzzleString)
	duration := time.Since(start)
	fmt.Printf("Part 1 result is: %v, computed in: %v \n", result, duration)
}
