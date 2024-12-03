package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func Part2() {
	multipliers := parser("day3/input.txt")
	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`

	re := regexp.MustCompile(pattern)
	var matches [][]string

	matches = re.FindAllStringSubmatch(multipliers, -1)
	if matches == nil {
		panic("No match found")
	}

	result := 0
	doFlag := true
	for _, match := range matches {
		if match[0] == "do()" {
			doFlag = true
			continue
		}
		if match[0] == "don't()" {
			doFlag = false
			continue
		}

		if doFlag {
			first, _ := strconv.Atoi(match[1])
			second, _ := strconv.Atoi(match[2])
			result += first * second
		}

	}
	fmt.Println(result)
}
