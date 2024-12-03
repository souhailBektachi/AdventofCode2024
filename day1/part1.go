package day1

import (
	"AdventofCode2024/utiles"
	"bufio"
	"fmt"
	"sort"
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
	sort.Ints(leftArray)
	sort.Ints(rightArray)
	return leftArray, rightArray
}
func Part1() {

	scanner, err := utiles.ReturnScanner("day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	leftArray, rightArray := parser(scanner)

	result := 0
	temp := 0
	for index := range leftArray {
		temp = leftArray[index] - rightArray[index]
		if temp < 0 {
			temp = temp * -1
		}
		result += temp

	}
	fmt.Println(result)

}
