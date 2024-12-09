package day9

import (
	"AdventofCode2024/utiles"
	"strconv"
	"strings"
)

var diskmap = parser("day9/input.txt")

func parser(path string) []int {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	var StringDiskMap []string
	for i := 0; scanner.Scan(); i++ {
		StringDiskMap = strings.Split(scanner.Text(), "")

	}
	var IntDiskMap = make([]int, len(StringDiskMap))
	for i, elem := range StringDiskMap {
		temp, _ := strconv.Atoi(elem)
		IntDiskMap[i] = temp
	}
	return IntDiskMap
}
