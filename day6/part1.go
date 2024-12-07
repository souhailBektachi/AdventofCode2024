package day6

import "fmt"

func Part1() {

	i, j := FindGuard()
	fmt.Println(moveGuard(i, j))
}
func moveOnepos(i, j int, direction direction) (position, bool, direction) {

	switch direction {
	case up:
		if i-1 != -1 {
			if Map[i-1][j] == "#" {
				direction = right

			} else {
				i--
			}

		} else {
			return position{-1, -1}, true, direction
		}

	case left:
		if j-1 > -1 {
			if Map[i][j-1] == "#" {
				direction = up

			} else {
				j--
			}

		} else {
			return position{-1, -1}, true, direction
		}

	case right:
		if j+1 != len(Map[i]) {
			if Map[i][j+1] == "#" {
				direction = down

			} else {
				j++
			}

		} else {
			return position{-1, -1}, true, direction
		}
	case down:
		if i+1 != len(Map) {
			if Map[i+1][j] == "#" {
				direction = left

			} else {
				i++
			}

		} else {
			return position{-1, -1}, true, direction
		}

	}
	return position{i, j}, false, direction

}
