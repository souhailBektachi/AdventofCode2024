package day14

import (
	. "AdventofCode2024/utiles"
	"fmt"
)

func Part2v2() {
	space := Position{101, 103}
	RobotsPos := make(map[Position]bool)
	second := 0
	avgX, avgY := 0, 0
	var_x, var_y := 0, 0
	for true {
		getNextPos(RobotsPos, space, second)
		for Robot := range RobotsPos {
			avgX += Robot.X
			avgY += Robot.Y
		}
		avgX /= len(posVel)
		avgY /= len(posVel)
		for Robot := range RobotsPos {

			var_x += (Robot.X - avgX) * (Robot.X - avgX)
			var_y += (Robot.Y - avgY) * (Robot.Y - avgY)
		}
		var_x /= len(posVel)
		var_y /= len(posVel)

		if var_x < 400 && var_y < 400 {
			fmt.Println("tree", second)
			break

		}

		second++
	}

}
