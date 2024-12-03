package day1

import (
	"AdventofCode2024/utiles"
	"fmt"
	"sort"
)

func Part1() {

	scanner, err, file := utiles.ReturnScanner("day1/input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	leftArray, rightArray := parser(scanner)
	sort.Ints(leftArray)
	sort.Ints(rightArray)

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
