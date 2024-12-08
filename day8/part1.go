package day8

import "fmt"

func Part1() {

	FreqCor := getFreqCord(FrequenciScan)
	for _, cord := range FreqCor {
		countAntiNodes(cord, len(FrequenciScan), len(FrequenciScan[0]))
	}
	fmt.Println(len(AntinNodeMap))
}

var AntinNodeMap = make(map[position]bool)

func countAntiNodes(Positions []position, height, width int) {
	for i, pos1 := range Positions {
		for j := i + 1; j < len(Positions); j++ {
			pos2 := Positions[j]
			// if !isInline(pos1, pos2) {
			// 	continue
			// }
			AntiNod1 := AddPos(pos1, SubPos(pos1, pos2))
			AntiNod2 := AddPos(pos2, SubPos(pos2, pos1))
			// if (AntiNod1.x == 1 && AntiNod1.y == 3) || (AntiNod2.x == 1 && AntiNod2.y == 3) {
			// 	fmt.Println(pos1, pos2)

			// }
			isValid(AntiNod1, height, width)

			isValid(AntiNod2, height, width)

		}

	}

}
