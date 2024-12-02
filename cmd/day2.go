package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertArrayStringToInt(strs []string) []int {
	var ints []int
	for _, str := range strs {
		if num, err := strconv.Atoi(str); err == nil {
			ints = append(ints, num)
		}
	}
	return ints
}

// Part one
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

// Part two
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

func parseInput() {
	scanner := bufio.NewScanner(os.Stdin)
	record := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		arr := convertArrayStringToInt(strings.Fields(line))
		record++
		if !isValidWrapper(arr) {
			record--
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	fmt.Printf("Total Safe Record: %d", record)
}
