package main

import (
	"fmt"
	"os"
	"strings"
)

func stringIntersection(strings ...string) {

}

func stringIntersectionBasic(string1 string, string2 string) {
	workingSet := make(map[int32]bool)
	for _, char := range string1 {
		workingSet[char] = true
	}
	intersectionSet := make(map[int32]bool)
	for _, char := range string2 {
		if _, found := workingSet[char]; found {
			intersectionSet[char] = true
		}
	}
}

func commonLetter(string1 string, string2 string) []int {
	var commonLetters []int
	for _, letter1 := range string1 {
		for _, letter2 := range string2 {
			if letter1 == letter2 {
				commonLetters = append(commonLetters, int(letter1))
			}
		}
	}
	return commonLetters
}

func getPriority(asciiChar int) int {
	if asciiChar > 96 {
		return asciiChar - 96
	} else {
		return asciiChar - 64 + 26
	}
}

func day3Part1(puzzleInput string) int {
	lines := strings.Split(puzzleInput, "\n")
	totalSum := 0
	for _, line := range lines {
		middlePoint := len(line) / 2
		first, second := line[:middlePoint], line[middlePoint:]
		asciiLetters := commonLetter(first, second)
		fmt.Println(asciiLetters)
		totalSum += getPriority(asciiLetters[0])
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
