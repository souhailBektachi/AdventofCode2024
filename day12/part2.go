package day12

import (
	. "AdventofCode2024/utiles"
	"fmt"
	"sort"
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
	Sorted := SortPostion()
	PositionByY := GroupPositionsByY(Sorted)
	for j, line := range PositionByY {
		for i, pos := range line {
			if i == 0 || i == len(line)-1 {
				if j > 0 && len(PositionByY[j-1]) > i && i == len(line)-1 {
					temp := Position{pos.X, pos.Y - 1}
					if temp == PositionByY[j-1][i] {
						continue
					}
				}
				if j < len(PositionByY)-1 && len(PositionByY[j+1]) > i && i == len(line)-1 {
					temp := Position{pos.X, pos.Y + 1}
					if temp == PositionByY[j+1][i] {
						continue
					}
				}
				sides++
				if _, exists := Positions[Position{pos.X + 1, pos.Y}]; !exists {
					sides++

				}
				if _, exists := Positions[Position{pos.X - 1, pos.Y}]; !exists {
					sides++
				}

			}
		}
		fmt.Println(sides)
	}
	fmt.Println(sides)
	return sides
}

func SortPostion() []Position {
	positions := make([]Position, 0, len(Positions))
	for pos := range Positions {
		positions = append(positions, pos)
	}
	sort.Slice(positions, func(i, j int) bool {
		if positions[i].Y == positions[j].Y {
			return positions[i].X < positions[j].X
		}
		return positions[i].Y < positions[j].Y
	})
	return positions
}
func GroupPositionsByY(sorted []Position) [][]Position {
	if len(sorted) == 0 {
		return nil
	}

	result := [][]Position{}
	currentY := sorted[0].Y
	currentGroup := []Position{sorted[0]}

	for i := 1; i < len(sorted); i++ {
		if sorted[i].Y == currentY {
			currentGroup = append(currentGroup, sorted[i])
		} else {
			result = append(result, currentGroup)
			currentGroup = []Position{sorted[i]}
			currentY = sorted[i].Y
		}
	}
	result = append(result, currentGroup)

	return result
}
