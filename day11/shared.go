package day11

import (
	"AdventofCode2024/utiles"
	"strings"
)

func parser(path string) []string {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	scanner.Scan()
	StonesLine := strings.Split(scanner.Text(), " ")

	return StonesLine
}
