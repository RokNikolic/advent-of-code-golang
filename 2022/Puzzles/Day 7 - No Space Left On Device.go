package year2022

import (
	"fmt"
	"os"
	"strings"
)

func day7part1(puzzleInput string) int {
	return 0
}

func day7part2(puzzleInput string) int {
	return 0
}

func Day7() {
	puzzleRead, err := os.ReadFile("2022/Input/day7.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		puzzleString := string(puzzleRead)
		puzzleCleaned := strings.Replace(puzzleString, "\r", "", -1)

		fmt.Printf("Part 1 result is: %v\n", day7part1(puzzleCleaned))
		fmt.Printf("Part 2 result is: %v\n", day7part2(puzzleCleaned))
	}
}
