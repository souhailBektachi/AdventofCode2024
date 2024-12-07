package day7

import "fmt"

func Part1() {
	equation := parser("day7/input.txt")
	result := 0
	for value, nums := range equation {
		if checkIfEvaluted(value, nums) {
			result += value
		}
	}
	fmt.Println(result)

}

func checkIfEvaluted(value int, nums []int) bool {
	combinations := generateCombination(len(nums))

	var evalNums func([]int, []rune) int
	evalNums = func(nums []int, operators []rune) int {
		result := nums[0]
		for i := 0; i < len(operators); i++ {
			switch operators[i] {
			case '+':
				result += nums[i+1]
			case '*':
				result *= nums[i+1]
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

func evaluate(op rune, nums [2]int) int {
	switch op {
	case '+':
		return nums[0] + nums[1]
	case '*':
		return nums[0] * nums[1]
	}
	return -1
}

func generateCombination(size int) [][]rune {
	if size == 1 {
		return [][]rune{{'+'}, {'*'}}
	}

	ops := []rune{'+', '*'}
	total := 1 << (size - 1)
	result := make([][]rune, total)

	for i := 0; i < total; i++ {
		combo := make([]rune, size-1)
		for j := 0; j < size-1; j++ {
			if (i & (1 << j)) != 0 {
				combo[j] = ops[1]
			} else {
				combo[j] = ops[0]
			}
		}
		result[i] = combo
	}
	return result
}
