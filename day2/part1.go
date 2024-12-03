package day2

import (
	"AdventofCode2024/utiles"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {

	scanner, err, file := utiles.ReturnScanner("day2/input.txt")
	defer file.Close()
	if err != nil {
		panic(err)

	}
	Levels := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		Levels = append(Levels, strings.Split(line, " "))

	}
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
	inc := isincreasing(level)
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

func isincreasing(level []string) bool {
	current, _ := strconv.Atoi(level[0])
	temp, _ := strconv.Atoi(level[1])
	if current > temp {
		return false
	}
	return true
}
