package day5

import "fmt"

func Parte1() {

	correctLines := make([][]int, 0)
outerLoop:
	for _, line := range Update {
		for i, num := range line {
			if i == len(line)-1 {
				break
			}
			if !CheckIfCorrect(line, i, num) {
				continue outerLoop
			}

		}
		correctLines = append(correctLines, line)
	}
	result := 0
	for _, correct := range correctLines {
		result += findmiddlePage(correct)
	}
	fmt.Println(result)
}

func CheckIfCorrect(line []int, index int, num int) bool {

	for i := index + 1; i < len(line); i++ {
		if i == index {
			continue
		}

		if _, exist := pageOrdering[num][line[i]]; exist {
			return false
		}
	}
	return true
}
