package day1

import (
	"slices"
	"strconv"
	"strings"

	"github.com/llauderesv/aoc-2024/utils"
)

func PartOne() int {
	input := utils.ReadInput("inputs/day1.txt")

	arr1, arr2 := NewArr(input)

	ans := 0
	for i, v := range arr1 {
		diff := v - arr2[i]
		if diff < 0 {
			diff = -diff
		}

		ans += diff
	}

	return ans
}

func PartTwo() int {
	input := utils.ReadInput("inputs/day1.txt")

	arr1, arr2 := NewArr(input)
	total := 0
	for _, v := range arr1 {
		c := 0
		for _, v2 := range arr2 {
			if v == v2 {
				c++
			}
		}

		total += v * c
	}

	return total
}

func NewArr(input string) ([]int, []int) {

	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		// Split line into fields (columns)
		fields := strings.Fields(line)
		if i, err := strconv.Atoi(fields[0]); err == nil {
			arr1 = append(arr1, i)
		}
		if i, err := strconv.Atoi(fields[1]); err == nil {
			arr2 = append(arr2, i)
		}
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	return arr1, arr2
}
