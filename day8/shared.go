package day8

import (
	"AdventofCode2024/utiles"
	"strings"
)

var FrequenciScan = parser("day8/input.txt")
var FreqCor = getFreqCord(FrequenciScan)

type position struct {
	x, y int
}

func getFreqCord(FrequenciScan [][]string) map[string][]position {
	FreqCor := make(map[string][]position)
	for x, line := range FrequenciScan {
		for y, freq := range line {
			if freq == "." {
				continue
			}
			if _, exist := FreqCor[freq]; !exist {
				FreqCor[freq] = make([]position, 0)
			}
			FreqCor[freq] = append(FreqCor[freq], position{x, y})

		}
	}
	return FreqCor
}
func parser(path string) [][]string {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	Scanne := make([][]string, 0)

	for i := 0; scanner.Scan(); i++ {
		Scanne = append(Scanne, strings.Split(scanner.Text(), ""))

	}
	return Scanne
}

func isValid(Pos position, height, width int) bool {
	if Pos.x < height && Pos.y < width && Pos.x > -1 && Pos.y > -1 {
		AntinNodeMap2[Pos] = true

		return true
	}
	return false

}
func SubPos(Position1, Position2 position) position {
	var result position
	result.x = Position1.x - Position2.x
	result.y = Position1.y - Position2.y
	return result
}
func AddPos(Pos1, Pos2 position) position {
	var result position
	result.x = Pos1.x + Pos2.x
	result.y = Pos1.y + Pos2.y
	return result
}
