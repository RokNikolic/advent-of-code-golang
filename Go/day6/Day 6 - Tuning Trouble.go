package day6

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

func partBoth(puzzleInput string, windowSize int) int {
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
	puzzleRead, err := os.ReadFile("Input/day6.txt")
	if err != nil {
		fmt.Println(err)
	}
	puzzleString := string(puzzleRead)
	puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

	fmt.Printf("Day 6 Part 1 result is: %v\n", partBoth(puzzleCleaned, 4))
	fmt.Printf("Day 6 Part 2 result is: %v\n", partBoth(puzzleCleaned, 14))
}
