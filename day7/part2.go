package day7

import (
	"fmt"
	"strconv"
)

func Part2() {
	equation := parser("day7/input.txt")
	result := 0
	for value, nums := range equation {
		if checkIfEvaluted2(value, nums) {
			result += value
		}
	}
	fmt.Println(result)

}

func checkIfEvaluted2(value int, nums []int) bool {
	combinations := generateCombination2(len(nums))
	var evalNums func([]int, []rune) int
	evalNums = func(nums []int, operators []rune) int {
		result := nums[0]
		for i := 0; i < len(operators); i++ {
			switch operators[i] {
			case '+':
				result += nums[i+1]
			case '*':
				result *= nums[i+1]
			case '|':
				StringNuM1 := strconv.Itoa(result)
				StringNuM2 := strconv.Itoa(nums[i+1])
				combined := StringNuM1 + StringNuM2
				result, _ = strconv.Atoi(combined)

			}
		}
		return result
	}

	for _, comb := range combinations {
		result := evalNums(nums, comb)
		if value == result {
			return true
		}
	}
	return false
}

var combinationsCache map[int][][]rune = make(map[int][][]rune)
var ops = []rune{'+', '*', '|'}

func generateCombination2(size int) [][]rune {

	if size == 1 {
		return [][]rune{{'+', '*', '|'}}
	}
	if cached, ok := combinationsCache[size]; ok {
		return cached
	}

	total := 1
	for i := 0; i < size-1; i++ {
		total *= 3
	}
	result := make([][]rune, total)

	for i := 0; i < total; i++ {
		combo := make([]rune, size-1)
		temp := i
		for j := 0; j < size-1; j++ {
			combo[j] = ops[temp%3]
			temp /= 3
		}
		result[i] = combo
	}
	combinationsCache[size] = result
	return result
}
