package day2

import (
	"AdventofCode2024/utiles"
	"strings"
)

func parser(path string) [][]string {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()
	if err != nil {
		panic(err)

	}
	Levels := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		Levels = append(Levels, strings.Split(line, " "))

	}
	return Levels
}
