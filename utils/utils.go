package utils

import (
	"io/ioutil"
	"log"
	"strconv"
)

func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	return string(data)
}

func ConvertArrayStringToInt(strs []string) []int {
	var ints []int
	for _, str := range strs {
		if num, err := strconv.Atoi(str); err == nil {
			ints = append(ints, num)
		}
	}
	return ints
}
