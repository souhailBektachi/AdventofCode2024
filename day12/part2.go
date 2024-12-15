package day12

import (
	. "AdventofCode2024/utiles"
	"fmt"
)

func Part2() {
	fmt.Println(traverseGarde2())
}

func traverseGarde2() int {
	result := 0
	for y, line := range garden {
		for x := 0; x < len(line); x++ {
			plant := garden[y][x]

			if plant == "." {
				continue
			}
			ClearAll()
			SearchForPlant(plant, Position{x, y})
			sides := CalculateSides()

			result += sides * plantCount
			plantCount = 0
			per = 0
			clearPosition()
		}
	}
	return result

}

func CalculateSides() int {
	sides := 0
	for pos := range Positions {

		sides += CheckSequare(pos)
	}
	fmt.Println(sides)
	fmt.Println("____________________")
	return sides
}
func CheckSequare(pos Position) int {
	sides := 0

	_, exist := Positions[Position{pos.X + 1, pos.Y}]
	if !exist {
		_, exist1 := Positions[Position{pos.X, pos.Y - 1}]
		_, exist2 := Positions[Position{pos.X, pos.Y + 1}]
		_, exist3 := Positions[Position{pos.X + 1, pos.Y - 1}]
		_, exist4 := Positions[Position{pos.X + 1, pos.Y + 1}]
		if !exist1 && !exist2 && !exist3 && !exist4 {
			sides += 2
		} else {
			if (!exist1 && !exist3) || (!exist2 && !exist4) {
				sides++
			}
		}
		_, exist1 = Positions[Position{pos.X + 1, pos.Y - 1}]
		_, exist2 = Positions[Position{pos.X + 1, pos.Y + 1}]
		if exist1 && exist2 {

			sides += 2

		} else if exist1 || exist2 {

			sides++
			if pos.X == 2 && pos.Y == 2 {
				fmt.Println(sides)
			}

		}

	}

	_, exist = Positions[Position{pos.X - 1, pos.Y}]
	if !exist {
		_, exist1 := Positions[Position{pos.X, pos.Y - 1}]
		_, exist2 := Positions[Position{pos.X, pos.Y + 1}]
		_, exist3 := Positions[Position{pos.X - 1, pos.Y - 1}]
		_, exist4 := Positions[Position{pos.X - 1, pos.Y + 1}]
		if !exist1 && !exist2 && !exist3 && !exist4 {
			sides += 2
		} else {
			if (!exist1 && !exist3) || (!exist2 && !exist4) {
				sides++

			}
		}
		_, exist1 = Positions[Position{pos.X - 1, pos.Y - 1}]
		_, exist2 = Positions[Position{pos.X - 1, pos.Y + 1}]
		if exist1 && exist2 {
			sides += 2

		} else if exist1 || exist2 {
			sides++
		}
	}

	// fmt.Println(pos, sides)

	return sides
}
