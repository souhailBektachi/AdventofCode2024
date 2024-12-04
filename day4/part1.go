package day4

import "fmt"

var XMAS []rune = []rune{'M', 'A', 'S'}

func Part1() {
	Words := parser("day4/input.txt")

	counter := 0
	for line := range Words {
		for j, word := range Words[line] {
			if word == 'X' {
				if traverseHorizontal(Words, false, line, j) {
					counter++
				}
				if traverseHorizontal(Words, true, line, j) {
					counter++
				}
				if traverseVertical(Words, false, line, j) {
					counter++
				}
				if traverseVertical(Words, true, line, j) {
					counter++
				}
				if traverseDiagonal(Words, [2]bool{true, false}, line, j) {
					counter++
				}
				if traverseDiagonal(Words, [2]bool{false, false}, line, j) {
					counter++
				}
				if traverseDiagonal(Words, [2]bool{false, true}, line, j) {
					counter++
				}
				if traverseDiagonal(Words, [2]bool{true, true}, line, j) {
					counter++
				}

			}
		}
	}
	fmt.Println(counter)
}

func traverseHorizontal(Words [][]rune, reverse bool, i int, j int) bool {
	if !incrementIJ(&j, len(Words[i]), reverse) {
		return false
	}

	for _, word := range XMAS {
		if Words[i][j] == word {

			if word == 'S' {
				return true
			}
			if !reverse {
				if j+1 < len(Words[i]) {
					j++
				} else {
					return false
				}
			} else {
				if j-1 >= 0 {
					j--

				} else {
					return false
				}

			}
			continue
		} else {
			return false
		}

	}
	return false

}
func incrementIJ(x *int, size int, reverse bool) bool {
	if !reverse {
		if (*x)+1 < size {
			(*x)++
		} else {

			return false
		}
	} else {
		if reverse && (*x)-1 >= 0 {
			(*x)--
		} else {

			return false
		}
	}
	return true

}
func incrementIJDIAG(i *int, j *int, size [2]int, reverse [2]bool) bool {

	if !(incrementIJ(i, size[0], reverse[0]) && incrementIJ(j, size[1], reverse[1])) {
		return false
	}

	return true

}
func traverseVertical(Words [][]rune, reverse bool, i int, j int) bool {
	if !incrementIJ(&i, len(Words), reverse) {
		return false
	}
	for _, word := range XMAS {
		if Words[i][j] == word {

			if word == 'S' {
				return true
			}
			if !reverse {
				if i+1 < len(Words) {
					i++
				} else {
					return false
				}
			} else {
				if i-1 >= 0 {
					i--

				} else {
					return false
				}

			}
			continue
		} else {
			return false
		}

	}
	return false
}

func traverseDiagonal(Words [][]rune, reverse [2]bool, i int, j int) bool {
	if !incrementIJDIAG(&i, &j, [2]int{len(Words), len(Words[i])}, reverse) {
		return false
	}
	for _, word := range XMAS {
		if Words[i][j] == word {

			if word == 'S' {

				return true
			}

			if !incrementIJDIAG(&i, &j, [2]int{len(Words), len(Words[i])}, reverse) {
				return false
			}
		} else {
			return false
		}

	}
	return false
}
