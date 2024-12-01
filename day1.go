package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]int, []int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		tuple := strings.Fields(line)
		if i, err := strconv.Atoi(tuple[0]); err == nil {
			arr1 = append(arr1, i)
		}
		if i, err := strconv.Atoi(tuple[1]); err == nil {
			arr2 = append(arr2, i)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return nil, nil, err
	}

	return arr1, arr2, nil
}

func partOne(arr1, arr2 *[]int) {
	total := 0
	for i, v := range *arr1 {
		diff := v - (*arr2)[i]
		if diff < 0 {
			diff = -diff // Convert to positive if negative
		}

		total += diff
	}

	fmt.Printf("Answer: %d\n", total)
}

func partTwo(arr1, arr2 *[]int) {
	total := 0
	for _, v := range *arr1 {
		c := 0
		for _, v2 := range *arr2 {
			if v == v2 {
				c++
			}
		}

		total += v * c
	}

	fmt.Printf("Answer: %d\n", total)
}
