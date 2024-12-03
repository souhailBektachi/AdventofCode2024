package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func Part1() {
	multipliers := parser("day3/example.txt")
	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`

	re := regexp.MustCompile(pattern)
	var matches [][]string

	matches = re.FindAllStringSubmatch(multipliers, -1)
	if matches == nil {
		panic("No match found")
	}

	result := 0
	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		result += first * second

	}
	fmt.Println(result)
}
