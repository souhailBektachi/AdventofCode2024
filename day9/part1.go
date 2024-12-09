package day9

import (
	"fmt"
)

func Part1() {
	diskmap := ConvertToFileLayout(diskmap)
	removeGaps(diskmap)
	fmt.Println(calculateCheckSum(diskmap))

}
func calculateCheckSum(diskmap []int) int {
	result := 0
	for index, elem := range diskmap {
		if elem == -1 {
			return result
		}

		result += index * elem
	}
	return 0
}
func ConvertToFileLayout(diskmap []int) []int {
	result := make([]int, 0)
	id := 0
	for index, elem := range diskmap {
		if index%2 == 0 {
			for i := 0; i < elem; i++ {
				result = append(result, id)
			}
			id++
			continue
		}
		for i := 0; i < elem; i++ {
			result = append(result, -1)
		}
		continue

	}
	return result
}
func removeGaps(diskmap []int) {

	secondPointer := findLast(diskmap, len(diskmap)-1)
	for firstPointer, elem := range diskmap {
		if firstPointer == secondPointer {
			break

		}
		if elem != -1 {
			continue
		}
		diskmap[firstPointer], diskmap[secondPointer] = diskmap[secondPointer], diskmap[firstPointer]
		secondPointer = findLast(diskmap, secondPointer)
	}

}

func findLast(diskmap []int, current int) int {
	for i := current; i > 0; i-- {
		if diskmap[i] != -1 {
			return i
		}
	}
	return -1
}
