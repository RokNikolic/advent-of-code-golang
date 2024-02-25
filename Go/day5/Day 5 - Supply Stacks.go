package day5

import (
	"fmt"
	"os"
	"strings"
)

func part1(puzzleInput string) int {
	return 0
}

func part2(puzzleInput string) int {
	return 0
}

func Day5() {
	puzzleRead, _ := os.ReadFile("Input/day3.txt")
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 3 Part 1 result is: %v\n", part1(puzzleCleaned))
	fmt.Printf("Day 3 Part 2 result is: %v\n", part2(puzzleCleaned))
}
