package day4

import "fmt"

var X_MAS map[string]string = map[string]string{
	"M.M": "S.S",
	"M.S": "M.S",
	"S.S": "M.M",
	"S.M": "S.M",
}

func Part2() {
	Words := parser("day4/input.txt")

	counter := 0
	for line := range Words {

		for j, word := range Words[line] {
			if word == 'A' {
				if traverseX_MAS(Words, line, j) {
					counter++
				}

			}
		}
	}
	fmt.Println(counter)

}
func traverseX_MAS(Words [][]rune, i int, j int) bool {
	WordsSize := len(Words)
	lineSize := len(Words[i])
	if !(incrementIJ(&i, WordsSize, true) && incrementIJ(&j, lineSize, true)) {
		return false
	}
	firstLine := string(Words[i][j]) + "."
	for k := 0; k < 2; k++ {
		if !incrementIJ(&j, lineSize, false) {
			return false
		}
	}
	firstLine += string(Words[i][j])
	for k := 0; k < 2; k++ {
		if !incrementIJ(&i, WordsSize, false) {
			return false
		}

	}
	j -= 2
	secondLine := string(Words[i][j]) + "."
	for k := 0; k < 2; k++ {
		if !incrementIJ(&j, lineSize, false) {
			return false
		}
	}
	secondLine += string(Words[i][j])
	if X_MAS[firstLine] == secondLine {
		return true
	} else {
		return false
	}

}
