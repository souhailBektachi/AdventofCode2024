package day6

import "fmt"

func Part2() {

	i, j := FindGuard()
	moveGuard(i, j)
	result := PutLoopObstacle(position{i, j})
	fmt.Println(result)

}
func CheckForLoops(index int) bool {
	IPos := IndexofPositions[index-1]
	currentPos, direction := IPos.POS, IPos.Dir
	var isEnd bool
	PositionMap := make(map[position]int)
	for {
		temppos := currentPos
		for {
			currentPos, isEnd, direction = moveOnepos(currentPos.i, currentPos.j, direction)

			if currentPos != temppos {
				break

			}
		}
		PositionMap[currentPos]++
		if PositionMap[currentPos] > 5 {
			return true
		}
		if isEnd {
			return false
		}

	}
}

func PutLoopObstacle(InitPos position) int {
	var temp indexedPosition
	counter := 0
	for index, Iposition := range IndexofPositions {

		if index == 1 {
			continue
		}
		temp = Iposition
		Map[temp.POS.i][temp.POS.j] = "#"

		if CheckForLoops(index) {
			counter++
		}
		Map[temp.POS.i][temp.POS.j] = "."

	}
	return counter
}
