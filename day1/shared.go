package day1

import (
	"bufio"
	"strconv"
	"strings"
)

func parser(scanner *bufio.Scanner) ([]int, []int) {
	leftArray, rightArray := make([]int, 0), make([]int, 0)
	for scanner.Scan() {
		noSpaces := strings.Split(scanner.Text(), "   ")
		righNum, _ := strconv.Atoi(noSpaces[1])
		leftNum, _ := strconv.Atoi(noSpaces[0])
		leftArray = append(leftArray, leftNum)
		rightArray = append(rightArray, righNum)

	}
	return leftArray, rightArray
}
