package year2022

import (
	"fmt"
	"os"
	"strings"
)

func slidingWindow(wholeSlice string, windowSize int) []string {
	var sliceOfWindows []string
	for i := 0; i < len(wholeSlice)-windowSize+1; i++ {
		currentWindow := wholeSlice[i : i+windowSize]
		sliceOfWindows = append(sliceOfWindows, currentWindow)
	}
	return sliceOfWindows
}

func day6part1(puzzleInput string) int {
	windowSize := 4
	windows := slidingWindow(puzzleInput, windowSize)
	for i, window := range windows {
		windowSet := make(map[int32]bool)
		for _, char := range window {
			windowSet[char] = true
		}
		if len(windowSet) == windowSize {
			return i + windowSize
		}
	}
	return -1
}

func day6part2(puzzleInput string) int {
	windowSize := 14
	windows := slidingWindow(puzzleInput, windowSize)
	for i, window := range windows {
		windowSet := make(map[int32]bool)
		for _, char := range window {
			windowSet[char] = true
		}
		if len(windowSet) == windowSize {
			return i + windowSize
		}
	}
	return -1
}

func Day6() {
	puzzleRead, err := os.ReadFile("2022/Input/day6.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		puzzleString := string(puzzleRead)
		puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

		fmt.Printf("Part 1 result is: %v\n", day6part1(puzzleCleaned))
		fmt.Printf("Part 2 result is: %v\n", day6part2(puzzleCleaned))
	}
}
