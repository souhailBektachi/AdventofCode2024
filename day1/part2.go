package day1

import (
	"AdventofCode2024/utiles"
	"fmt"
)

func Part2() {
	scanner, err, file := utiles.ReturnScanner("day1/input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	leftArray, rightArray := parser(scanner)
	apparancesMap := make(map[int]int)
	for _, num := range rightArray {
		apparancesMap[num]++

	}
	result := 0
	for _, num := range leftArray {
		if _, exist := apparancesMap[num]; exist {
			result += num * apparancesMap[num]
		}
	}
	fmt.Println(result)

}
