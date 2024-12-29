package day4

import (
	"strings"

	"github.com/llauderesv/aoc-2024/utils"
)

func PartOne() int {
	input := utils.ReadInput("inputs/day4.txt")
	grid := strings.Split(input, "\n")
	word := "XMAS"

	countWordOccurrences := func(grid []string, word string) int {
		rows, cols := len(grid), len(grid[0])
		wordLength := len(word)
		count := 0

		// Define all 8 possible directions
		directions := [][2]int{
			{0, 1},   // Horizontal right
			{0, -1},  // Horizontal left
			{1, 0},   // Vertical down
			{-1, 0},  // Vertical up
			{1, 1},   // Diagonal down-right
			{-1, -1}, // Diagonal up-left
			{1, -1},  // Diagonal down-left
			{-1, 1},  // Diagonal up-right
		}

		// Helper function to check if the word exists in a given direction
		checkDirection := func(x, y, dx, dy int) bool {
			for i := 0; i < wordLength; i++ {
				nx, ny := x+i*dx, y+i*dy
				if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != word[i] {
					return false
				}
			}
			return true
		}

		// Iterate through every cell in the grid
		for x := 0; x < rows; x++ {
			for y := 0; y < cols; y++ {
				// Check all directions from the current cell
				for _, d := range directions {
					dx, dy := d[0], d[1]
					if checkDirection(x, y, dx, dy) {
						count++
					}
				}
			}
		}

		return count
	}

	ans := countWordOccurrences(grid, word)

	return ans
}

func PartTwo() int {
	input := utils.ReadInput("inputs/day4.txt")
	grid := strings.Split(input, "\n")

	allLines := []string{}
	for _, line := range grid {
		if len(line) > 0 {
			allLines = append(allLines, line)
		}
	}

	width, height := len(allLines[0]), len(allLines)

	checkXmasAt := func(row, col int) {
		if allLines[row][col] != 'A' {
			return
		}
	}

	symbolAt := func(row, col int) rune {
		if row < 0 {
			return '0'
		}
		if col < 0 {
			return ""
		}

		if row >= height {
			return ""
		}
		if col >= width {
			return ""
		}

		return rune(allLines[row][col])
	}

	return 0
}
