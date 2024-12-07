package day5

import "fmt"

func Parte2() {
	uncorrectLines := getTheUncorrects()
	result := 0
	fixUncorectLines(uncorrectLines)
	for _, uncorecct := range uncorrectLines {
		result += findmiddlePage(uncorecct)
	}

	fmt.Println(result)
}
func getTheUncorrects() [][]int {
	uncorrectLines := make([][]int, 0)

outerLoop:

	for _, line := range Update {
		for i, num := range line {
			if CheckIfUncorect(line, i, num) {
				uncorrectLines = append(uncorrectLines, line)
				continue outerLoop
			}

		}
	}
	return uncorrectLines

}
func CheckIfUncorect(line []int, index int, num int) bool {
	size := len(line)
	for i := index + 1; i < size; i++ {

		if _, exist := pageOrdering[num][line[i]]; exist {
			return true
		}
	}
	return false
}

func fixUncorectLines(uncorrectLines [][]int) {
	for i, line := range uncorrectLines {
		size := len(line)
		swapped := true

		for swapped {
			swapped = false
			for j := 0; j < size-1; j++ {
				if _, exists := pageOrdering[line[j+1]][line[j]]; exists {
					line[j], line[j+1] = line[j+1], line[j]
					swapped = true
				}
			}
		}
		uncorrectLines[i] = line
	}
}
