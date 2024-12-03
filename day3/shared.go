package day3

import (
	"AdventofCode2024/utiles"
)

func parser(path string) string {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()
	if err != nil {
		panic(err)

	}
	var result string
	for scanner.Scan() {
		result += scanner.Text()
	}
	return result
}
