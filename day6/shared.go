package day6

import (
	"AdventofCode2024/utiles"
	"strings"
)

var Map = parser("day6/input.txt")

type direction int

const (
	left direction = iota
	right
	up
	down
)

type position struct {
	i int
	j int
}
type indexedPosition struct {
	POS position
	Dir direction
}

var PossibleGuard = map[string]direction{"<": left, ">": right, "^": up, "v": down}
var PossitionsVistied = make(map[position]bool)
var IndexofPositions = make(map[int]indexedPosition)

func parser(path string) [][]string {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	Map := make([][]string, 0)
	for scanner.Scan() {

		line := strings.Split(scanner.Text(), "")
		Map = append(Map, line)
	}
	return Map
}
func FindGuard() (int, int) {
	for i, line := range Map {
		for j, elem := range line {
			if isGuard(elem) {
				return i, j
			}
		}

	}
	return -1, -1
}
func isGuard(Pguard string) bool {
	if _, exist := PossibleGuard[Pguard]; exist {
		return true

	}
	return false
}

func moveGuard(i, j int) int {
	direction := PossibleGuard[Map[i][j]]
	index := 0
	PossitionsVistied[position{i, j}] = true
	IndexofPositions[index] = indexedPosition{position{i, j}, direction}
	index++

	for i < len(Map) && j < len(Map[i]) {
		temppos, isEnd, Newdir := moveOnepos(i, j, direction)
		direction = Newdir

		if isEnd {
			break
		}

		i, j = temppos.i, temppos.j

		if _, exist := PossitionsVistied[position{i, j}]; !exist {
			IndexofPositions[index] = indexedPosition{position{i, j}, direction}
			index++
		}
		PossitionsVistied[position{i, j}] = true
	}
	return len(PossitionsVistied)

}
