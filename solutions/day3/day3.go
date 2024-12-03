package day3

import (
	"regexp"
	"strconv"

	"github.com/llauderesv/aoc-2024/utils"
)

func Mul(a, b int) int {
	return a * b
}

// RegEx?
// Parse all the characters that has mul()
func PartOne() int {
	input := utils.ReadInput("inputs/day3.txt")
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	re2 := regexp.MustCompile(`\d+`)

	arr := re.FindAllStringSubmatch(input, -1)
	ans := 0
	for _, v := range arr {
		nums := re2.FindAllStringSubmatch(v[0], -1)
		a, _ := strconv.Atoi(nums[0][0])
		b, _ := strconv.Atoi(nums[1][0])
		ans += Mul(a, b)
	}

	return ans
}

func PartTwo() int {
	input := utils.ReadInput("inputs/day3.txt")
	r := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	r2 := regexp.MustCompile(`\d+`)

	ans := 0

	arr := r.FindAllStringSubmatch(input, -1)
	enabled := true
	for _, str := range arr {
		if str[0] == "don't()" {
			enabled = false
			continue
		} else if str[0] == "do()" {
			enabled = true
			continue
		}

		if !enabled {
			continue
		}

		nums := r2.FindAllStringSubmatch(str[0], -1)
		a, _ := strconv.Atoi(nums[0][0])
		b, _ := strconv.Atoi(nums[1][0])
		ans += Mul(a, b)
	}

	return ans
}
