package day2

import (
	"fmt"
	"strconv"
)

func Part2() {
	Levels := parser("day2/input.txt")
	safes := 0
	for _, level := range Levels {
		if traverse(level) {
			safes++
			continue
		}
	}
	fmt.Println(safes)
}
func traverse(level []string) bool {
	for i := 0; i < len(level); i++ {
		if isSafe2(level, i) {
			return true
		}
	}
	return false

}
func isSafe2(level []string, unsafe int) bool {
	current, _ := strconv.Atoi(level[0])
	next, _ := strconv.Atoi(level[1])

	inc, safe := isincreasing2(level)
	if !safe {
		return false
	}
	i := 1
	if unsafe == 0 {
		current, _ = strconv.Atoi(level[1])
		i = 2
	}

	for i < len(level) {
		if unsafe == i {
			i++
			continue
		}
		next, _ = strconv.Atoi(level[i])
		Diffrence := next - current
		if (inc && next > current && Diffrence <= 3) || (!inc && current > next && Diffrence >= -3) {
			current = next
			i++
			continue
		} else {

			return false
		}

	}

	return true

}

func isincreasing2(level []string) (bool, bool) {
	current, _ := strconv.Atoi(level[0])
	inc, dec := 0, 0

	for i := 1; i < len(level); i++ {
		next, _ := strconv.Atoi(level[i])
		Diffrence := next - current

		if Diffrence > 0 {
			inc++
		} else {
			dec++
		}
		current = next
	}
	if inc == dec {
		return false, false
	}
	return inc > dec, true

}
