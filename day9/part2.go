package day9

import (
	"AdventofCode2024/utiles"
	"fmt"
)

func Part2() {
	defer utiles.Timer("Part2")()
	diskmap := ConvertToFileLayout(diskmap)
	fillGaps(diskmap)
	fmt.Println(CalculateSum(diskmap))

}
func CalculateSum(diskmap []int) int {
	id := 0
	result := 0
	for _, elem := range diskmap {
		if elem != -1 {
			result += elem * id
		}
		id++
	}
	return result
}

func fillGaps(diskmap []int) {
	seconPointer := len(diskmap) - 1
	noMovePossible := false

	for {
		if seconPointer < 1 || noMovePossible {
			break
		}
		lastBlock, count := findLastBloc(diskmap, seconPointer)
		gap := findGap(diskmap, count, lastBlock)
		if gap != -1 {
			for i := 0; i < count; i++ {
				diskmap[gap+i], diskmap[lastBlock+i] = diskmap[lastBlock+i], diskmap[gap+i]
			}
			noMovePossible = false
		} else {
			if seconPointer == lastBlock {
				noMovePossible = true
			}
		}
		seconPointer = lastBlock - 1
	}
}

func findGap(diskmap []int, size int, lenght int) int {
	counter := 0

	for i := 0; i < lenght; i++ {
		if diskmap[i] == -1 {
			counter++

		} else {
			counter = 0
		}
		if counter == size {
			return i - size + 1
		}
	}
	return -1
}
func findLastBloc(diskmap []int, lastIndex int) (int, int) {
	var index int
	for i := lastIndex; i >= 0; i-- {
		if diskmap[i] != -1 {
			index = i
			break
		}
	}

	value := diskmap[index]
	count := 0

	for index >= 0 && diskmap[index] == value {
		count++
		index--
	}
	return index + 1, count
}
