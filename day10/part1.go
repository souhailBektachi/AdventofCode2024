package day10

import (
	. "AdventofCode2024/utiles"
	"fmt"
)

type Part int

const (
	part1 Part = iota
	part2
)

var NinesFound map[Position]bool = make(map[Position]bool)

func Part1() {
	IntTopoMap := parser("day10/input.txt")
	fmt.Println(findZero(IntTopoMap, part1))

}
func Part2() {
	clear(NinesFound)
	IntTopoMap := parser("day10/input.txt")
	fmt.Println(findZero(IntTopoMap, part2))

}
