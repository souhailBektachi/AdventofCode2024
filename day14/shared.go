package day14

import (
	. "AdventofCode2024/utiles"
	"strconv"

	"regexp"
)

type PosVel struct {
	Pos      Position
	Velocity Position
}

func parser(path string) []PosVel {
	scanner, err, file := ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	posVel := make([]PosVel, 0)

	pattern := `\d+|-\d+`
	re := regexp.MustCompile(pattern)
	for i := 0; scanner.Scan(); i++ {
		posVel = append(posVel, PosVel{})
		matches := re.FindAllString(scanner.Text(), -1)
		nums := make([]int, 4)
		for i, m := range matches {
			nums[i], _ = strconv.Atoi(m)
		}

		posVel[i].Pos = Position{nums[0], nums[1]}
		posVel[i].Velocity = Position{nums[2], nums[3]}

	}

	return posVel
}
