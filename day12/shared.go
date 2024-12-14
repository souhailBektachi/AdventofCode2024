package day12

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
	garden := make([][]string, 0)
	for scanner.Scan() {
		garden = append(garden, strings.Split(scanner.Text(), ""))

	}

	return garden
}
