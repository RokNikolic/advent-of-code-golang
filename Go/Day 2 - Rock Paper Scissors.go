package main

import (
	"fmt"
	"os"
	"strings"
)

func day2Part1(puzzleInput string) int {
	replace1 := strings.Replace(puzzleInput, "X", "A", -1)
	replace2 := strings.Replace(replace1, "Y", "B", -1)
	replace3 := strings.Replace(replace2, "Z", "C", -1)
	lines := strings.Split(replace3, "\n")

	totalStore := 0
	for _, line := range lines {
		roundScore := 0
		shapes := strings.Split(line, " ")
		opponentsShape, yourShape := shapes[0], shapes[1]

		switch yourShape {
		case "A":
			roundScore += 1
		case "B":
			roundScore += 2
		case "C":
			roundScore += 3
		}

		if opponentsShape == yourShape {
			roundScore += 3
		} else if (yourShape == "A" && opponentsShape == "C") || (yourShape == "B" && opponentsShape == "A") || (yourShape == "C" && opponentsShape == "B") {
			roundScore += 6
		}
		totalStore += roundScore
	}

	return totalStore
}

func day2() {
	puzzleRead, _ := os.ReadFile("Input/day2.txt")
	puzzleString := string(puzzleRead)
	cleanedPuzzleString := strings.Replace(puzzleString, "\r", "", -1)

	result1 := day2Part1(cleanedPuzzleString)
	fmt.Printf("Day2 Part 1 result is: %v\n", result1)

	//result2 := day2Part1(puzzleString, 2)
	//fmt.Printf("Part 2 result is: %v\n", result2)
}
