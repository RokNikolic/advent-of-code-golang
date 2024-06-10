package day7

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

func Day7() {
	puzzleRead, err := os.ReadFile("../Input/day7.txt")
	if err != nil {
		fmt.Println(err)
	}
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 7 Part 1 result is: %v\n", part1(puzzleCleaned))
	fmt.Printf("Day 7 Part 2 result is: %v\n", part2(puzzleCleaned))
}
