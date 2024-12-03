package day2

import (
	"fmt"
	"strconv"
)

func Part1() {
	Levels := parser("day2/input.txt")
	safes := 0
	for _, level := range Levels {
		if isSafe(level) {
			safes++
		}
	}
	fmt.Println(safes)
}

func isSafe(level []string) bool {
	current, _ := strconv.Atoi(level[0])
	inc := isincreasing(level, 0, 1)
	for i := 1; i < len(level); i++ {
		next, _ := strconv.Atoi(level[i])
		Diffrence := next - current

		if (inc && Diffrence <= 3 && Diffrence > 0) || (!inc && Diffrence >= -3 && Diffrence < 0) {
			current = next
			continue
		} else {
			return false
		}

	}
	return true
}

func isincreasing(level []string, ex1 int, ex2 int) bool {
	current, _ := strconv.Atoi(level[ex1])
	temp, _ := strconv.Atoi(level[ex2])
	if current > temp {
		return false
	}
	return true
}
