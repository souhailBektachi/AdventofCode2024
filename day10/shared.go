package day10

import (
	"AdventofCode2024/utiles"
	. "AdventofCode2024/utiles"
	"strconv"
	"strings"
)

func parser(path string) [][]int {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	var StrTopoMap [][]string
	for i := 0; scanner.Scan(); i++ {
		StrTopoMap = append(StrTopoMap, strings.Split(scanner.Text(), ""))

	}
	var IntTopoMap [][]int
	for _, row := range StrTopoMap {
		var intRow []int
		for _, val := range row {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			intRow = append(intRow, intVal)
		}
		IntTopoMap = append(IntTopoMap, intRow)
	}
	return IntTopoMap
}

func findZero(IntTopoMap [][]int, part Part) int {
	score := 0
	for x, line := range IntTopoMap {

		for y, elem := range line {
			if elem == 0 {
				if part == part1 {
					clear(NinesFound)
				}
				find9Ways(IntTopoMap, Position{x, y}, &score, part)

			}
		}
	}
	return score
}

func find9Ways(IntTopoMap [][]int, Pos Position, score *int, part Part) {
	if IntTopoMap[Pos.X][Pos.Y] == 9 {
		if part == part1 {
			if _, exist := NinesFound[Pos]; exist {
				return
			} else {
				NinesFound[Pos] = true
			}
		}
		(*score) += 1
		return
	}

	Possibleways := findPossibleway(IntTopoMap, Pos)
	for _, pos := range Possibleways {

		find9Ways(IntTopoMap, pos, score, part)

	}

}

func findPossibleway(IntTopoMap [][]int, Pos Position) []Position {
	result := make([]Position, 0)
	for i := Pos.X - 1; i <= Pos.X+1; i++ {
		for j := Pos.Y - 1; j <= Pos.Y+1; j++ {
			if i == Pos.X && j == Pos.Y {
				continue
			}
			if i != Pos.X && j != Pos.Y {
				continue
			}
			if i >= len(IntTopoMap) || i < 0 || j >= len(IntTopoMap[0]) || j < 0 {
				continue
			} else {
				if IntTopoMap[i][j] == IntTopoMap[Pos.X][Pos.Y]+1 {
					result = append(result, Position{i, j})
				}
			}
		}
	}

	return result
}
