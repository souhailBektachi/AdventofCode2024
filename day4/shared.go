package day4

import "AdventofCode2024/utiles"

func parser(path string) [][]rune {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	Words := make([][]rune, 0)
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		Words = append(Words, make([]rune, 0))

		for _, word := range line {
			Words[j] = append(Words[j], word)
		}
		j++

	}
	return Words
}
