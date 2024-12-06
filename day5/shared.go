package day5

import (
	"AdventofCode2024/utiles"
	"strconv"
	"strings"
)

var pageOrdering, Update = parser("day5/input.txt")

func parser(path string) (map[int]map[int]bool, [][]int) {
	scanner, err, file := utiles.ReturnScanner(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	pageOrdering := make(map[int]map[int]bool)
	Update := make([][]int, 0)
	i := 0
	flag := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			flag = true
			continue
		}
		if !flag {
			temparr := strings.Split(line, "|")
			firstNum, _ := strconv.Atoi(temparr[0])
			secondNum, _ := strconv.Atoi(temparr[1])
			if _, exist := pageOrdering[secondNum]; !exist {
				pageOrdering[secondNum] = make(map[int]bool)

			}

			pageOrdering[secondNum][firstNum] = false

		} else {
			temparr := strings.Split(line, ",")
			Update = append(Update, make([]int, 0))
			for _, num := range temparr {
				tempNum, _ := strconv.Atoi(num)
				Update[i] = append(Update[i], tempNum)

			}
			i++

		}

	}
	return pageOrdering, Update

}
func findmiddlePage(line []int) int {
	size := len(line)
	return line[size/2]
}
