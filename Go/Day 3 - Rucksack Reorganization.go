package main

import (
	"fmt"
	"os"
	"strings"
)

func day3Part1(puzzleInput string) int {
	//lines := strings.Split(puzzleInput, "\n")

	return 0
}

func day3Part2(puzzleInput string) int {
	//lines := strings.Split(puzzleInput, "\n")

	return 0
}

func day3() {
	puzzleRead, _ := os.ReadFile("Input/day3.txt")
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 3 Part 1 result is: %v\n", day3Part1(puzzleCleaned))
	fmt.Printf("Day 3 Part 2 result is: %v\n", day3Part2(puzzleCleaned))
}
