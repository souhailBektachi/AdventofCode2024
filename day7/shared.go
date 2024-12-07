package day7

import (
	"AdventofCode2024/utiles"
	"strconv"
	"strings"
)

func parser(path string) map[int][]int {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	equation := make(map[int][]int)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		result, _ := strconv.Atoi(line[0])
		equation[result] = StringToNUms(line[1])

	}
	return equation
}

func StringToNUms(String string) []int {
	arr := strings.Split(String, " ")
	result := make([]int, len(arr))
	for i, num := range arr {
		result[i], _ = strconv.Atoi(num)
	}

	return result
}
