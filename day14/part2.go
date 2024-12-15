package day14

import (
	. "AdventofCode2024/utiles"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func Part2() {
	space := Position{101, 103}
	RobotsPos := make(map[Position]bool)
	second := 434

	for true {
		img := image.NewRGBA(image.Rect(0, 0, space.X, space.Y))
		background := color.RGBA{0, 0, 0, 255}
		robotColor := color.RGBA{255, 255, 255, 255}

		for y := 0; y < space.Y; y++ {
			for x := 0; x < space.X; x++ {
				img.Set(x, y, background)
			}
		}

		getNextPos(RobotsPos, space, second)

		for pos := range RobotsPos {
			img.Set(pos.X, pos.Y, robotColor)
		}

		fileName := fmt.Sprintf("day14/images/frame_%d.png", second)
		os.MkdirAll("day14/images", os.ModePerm)
		f, _ := os.Create(fileName)
		png.Encode(f, img)
		f.Close()

		second++
	}

}
func getNextPos(RobotsPos map[Position]bool, space Position, second int) {
	clear(RobotsPos)
	for _, Robot := range posVel {
		RobotsPos[MoveRobot(Robot, second, space)] = true
	}
}
