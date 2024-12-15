package day12

import (
	. "AdventofCode2024/utiles"
	"fmt"
)

var garden = parser("day12/input.txt")
var per int

func Part1() {
	fmt.Println("part1")
	result := traverseGarde()
	fmt.Println(result)
}
func traverseGarde() int {
	result := 0
	for y, line := range garden {
		for x := 0; x < len(line); x++ {
			plant := garden[y][x]

			if plant == "." {
				continue
			}
			ClearAll()
			SearchForPlant(plant, Position{x, y})
			result += per * plantCount
			plantCount = 0
			per = 0
			clearPosition()
		}
	}
	return result

}
func ClearAll() {
	clear(Positions)

}

var plantCount int
var Positions = make(map[Position]bool)

func SearchForPlant(Plant string, Pos Position) {
	_, exist := Positions[Pos]
	if exist {
		return
	}
	if garden[Pos.Y][Pos.X] != Plant {
		per++
		return
	}
	Positions[Pos] = true
	plantCount++

	directions := []Position{
		{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1},
	}

	for _, dir := range directions {
		newPos := Position{X: Pos.X + dir.X, Y: Pos.Y + dir.Y}
		if newPos.X >= 0 && newPos.X < len(garden[0]) && newPos.Y >= 0 && newPos.Y < len(garden) {
			SearchForPlant(Plant, newPos)
		} else {
			per++
		}
	}
}

func clearPosition() {
	for pos, _ := range Positions {
		garden[pos.Y][pos.X] = "."
	}
}
