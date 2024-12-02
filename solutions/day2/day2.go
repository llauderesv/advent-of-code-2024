package day2

import (
	"strings"

	"github.com/llauderesv/aoc-2024/utils"
)

func PartOne() int {
	input := utils.ReadInput("inputs/day2.txt")
	var count int
	for _, v := range strings.Split(input, "\n") {
		if v != "" {
			count++
			arr := utils.ConvertArrayStringToInt(strings.Split(v, " "))
			if !isValid(arr) {
				count--
			}
		}
	}

	return count
}

func PartTwo() int {
	input := utils.ReadInput("inputs/day2.txt")
	var count int
	for _, v := range strings.Split(input, "\n") {
		if v != "" {
			count++
			arr := utils.ConvertArrayStringToInt(strings.Split(v, " "))
			if !isValidWrapper(arr) {
				count--
			}
		}
	}

	return count
}

func isValid(arr []int) bool {
	j := arr[0]
	tmp := 0
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		d := j - v

		if i > 1 {
			if tmp < 0 {
				if d > 0 {
					return false
				}
			} else if tmp > 0 {
				if d < 0 {
					return false
				}
			}
		}

		tmp = d

		// make positive
		if d < 0 {
			d = -d
		}

		if d > 3 || d < 1 {
			return false
		}

		j = v
	}

	return true
}

func isValidWrapper(arr []int) bool {
	s := len(arr)
	for i := 0; i < s; i++ {

		tmpArr := make([]int, 0)
		for j := 0; j < s; j++ {
			if j != i {
				tmpArr = append(tmpArr, arr[j])
			}
		}

		if isValid(tmpArr) {
			return true
		}
	}

	return false
}
