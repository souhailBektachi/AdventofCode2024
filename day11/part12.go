package day11

import (
	"AdventofCode2024/utiles"
	"fmt"
	"strconv"
)

var limit = 75

func Part12() {
	defer utiles.Timer("test")()
	StonesLine := parser("day11/input.txt")
	score := 0
	for _, elem := range StonesLine {
		score += Blinc(0, elem)
	}
	fmt.Println("score", score)

}

var Stones map[string]int = make(map[string]int)

func Blinc(index int, stone string) int {
	key := fmt.Sprintf("%d-%s", index, stone)
	if val, exist := Stones[key]; exist {
		return val
	}

	if index == limit {
		Stones[key] = 1
		return 1
	}

	index++
	var result int

	if stone == "0" {
		result = Blinc(index, "1")
	} else if len(stone)%2 == 0 {
		stone1, stone2 := split(stone)
		result = Blinc(index, stone1) + Blinc(index, stone2)
	} else {
		result = Blinc(index, multiplyBy2024(stone))
	}

	Stones[key] = result
	return result
}

func clearStones() {
	Stones = make(map[string]int)
}

func removeZero(value string) string {
	var result string
	flag := false
	for _, char := range value {
		if char != '0' {
			flag = true
		}
		if flag {
			result += string(char)
		}
	}
	if len(result) == 0 {
		return "0"
	} else {
		return result
	}
}
func split(stone string) (string, string) {
	stone1 := stone[:len(stone)/2]
	stone2 := stone[len(stone)/2:]
	if stone2[0] == '0' {
		stone2 = removeZero(stone2)
	}

	return stone1, stone2

}

func zeroToOne(stone string) string {
	return "1"

}
func multiplyBy2024(stone string) string {
	temp, _ := strconv.ParseUint(stone, 10, 64)

	temp = temp * 2024

	result := strconv.FormatUint(temp, 10)
	return result
}
