package day13

import (
	. "AdventofCode2024/utiles"
	"fmt"
)

func Part1() {
	defer Timer("Part 1")()
	fmt.Println("part 1")
	behavior := parser("day13/input.txt")
	Part1result := 0
	Part2result := 0
	part1 := 0
	part2 := 10000000000000
	for _, b := range behavior {
		temp1 := FindWaytoWin(b, part1)
		temp2 := FindWaytoWin(b, part2)
		if temp1 != -1 {
			Part1result += temp1
		}
		if temp2 != -1 {
			Part2result += temp2
		}
	}
	fmt.Println("Part1", Part1result)
	fmt.Println("Part2", Part2result)
}
func FindWaytoWin(behavior buttonBehavior, part int) int {
	behavior.prize.X += part
	behavior.prize.Y += part

	ax, ay := behavior.A.X, behavior.A.Y
	bx, by := behavior.B.X, behavior.B.Y
	tx, ty := behavior.prize.X, behavior.prize.Y

	det := ax*by - ay*bx
	if det == 0 {
		return -1
	}

	m := float64(tx*by-ty*bx) / float64(det)
	n := float64(ax*ty-ay*tx) / float64(det)

	if m != float64(int64(m)) || n != float64(int64(n)) {
		return -1
	}

	cost := int(m)*3 + int(n)
	return cost
}
func addPosition(Pos1 Position, Pos2 Position) Position {
	Pos1.X += Pos2.X
	Pos1.Y += Pos2.Y
	return Pos1

}
func ClickOnButton(inc Position, clicks int) Position {
	value := Position{0, 0}
	value.X += inc.X * clicks
	value.Y += inc.Y * clicks
	return value

}
