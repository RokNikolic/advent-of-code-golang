package main

import (
	"fmt"
	"os"
	"strings"
)

func commonLetter(string1 string, string2 string) int {
	for _, letter1 := range string1 {
		for _, letter2 := range string2 {
			if letter1 == letter2 {
				return int(letter1)
			}
		}
	}
	return -1
}

func day3Part1(puzzleInput string) int {
	lines := strings.Split(puzzleInput, "\n")
	totalSum := 0
	for _, line := range lines {
		middlePoint := len(line) / 2
		first, second := line[:middlePoint], line[middlePoint:]
		asciiLetter := commonLetter(first, second)
		if asciiLetter > 96 {
			totalSum += asciiLetter - 96
		} else {
			totalSum += asciiLetter - 64 + 26
		}
	}
	return totalSum
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
