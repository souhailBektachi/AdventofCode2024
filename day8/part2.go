package day8

import "fmt"

func Part2() {
	for _, cord := range FreqCor {
		countAntiNodes2(cord, len(FrequenciScan), len(FrequenciScan[0]))
	}
	result := len(AntinNodeMap2)
	fmt.Println(result)
}

var AntinNodeMap2 = make(map[position]bool)

func countAntiNodes2(Positions []position, height, width int) {
	for i, pos1 := range Positions {
		for j := i + 1; j < len(Positions); j++ {
			pos2 := Positions[j]
			// if !isInline(pos1, pos2) {
			// 	continue
			// }
			AntiNod1 := pos1
			AntiNod2 := pos2
			AntinNodeMap2[AntiNod1] = true
			AntinNodeMap2[AntiNod2] = true
			for isValid(AntiNod1, height, width) {
				AntiNod1 = AddPos(AntiNod1, SubPos(pos1, pos2))
			}
			for isValid(AntiNod2, height, width) {
				AntiNod2 = AddPos(AntiNod2, SubPos(pos2, pos1))
			}

		}

	}

}
