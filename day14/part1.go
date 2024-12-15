package day14

import (
	. "AdventofCode2024/utiles"
	"fmt"
)

var posVel = parser("day14/input.txt")

func Part1() {

	result := 1
	space := Position{101, 103}
	var quadrant [4]int
	for _, Robot := range posVel {
		newPos := MoveRobot(Robot, 100, space)

		if newPos.X == space.X/2 || newPos.Y == space.Y/2 {
			continue

		}
		q := 0
		if newPos.X >= space.X/2 {
			q += 1
		}
		if newPos.Y >= space.Y/2 {
			q += 2
		}
		quadrant[q]++
	}
	for _, num := range quadrant {
		result *= num
	}
	fmt.Println(result)

}

func MoveRobot(RobotPosVel PosVel, seconds int, space Position) Position {
	NewX := (RobotPosVel.Pos.X + (RobotPosVel.Velocity.X * seconds)) % space.X
	NewY := (RobotPosVel.Pos.Y + (RobotPosVel.Velocity.Y * seconds)) % space.Y
	if NewX < 0 {
		NewX = space.X + NewX
	}
	if NewY < 0 {
		NewY = space.Y + NewY
	}
	return Position{NewX, NewY}
}
