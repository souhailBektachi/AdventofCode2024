package day13

import (
	. "AdventofCode2024/utiles"
	"regexp"
	"strconv"
)

type buttonBehavior struct {
	A     Position
	B     Position
	prize Position
}

func parser(path string) []buttonBehavior {
	scanner, err, file := ReturnScanner(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	behavior := make([]buttonBehavior, 0)
	i := 0
	j := 0
	pattern := `\d+`
	re := regexp.MustCompile(pattern)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if i%3 == 0 {
			j = 0
			behavior = append(behavior, buttonBehavior{})
		}
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		Xint, _ := strconv.Atoi(matches[0][0])
		Yint, _ := strconv.Atoi(matches[1][0])
		var pointer *Position
		switch j {
		case 0:
			pointer = &behavior[i/3].A
		case 1:
			pointer = &behavior[i/3].B
		case 2:
			pointer = &behavior[i/3].prize
		}
		(*pointer) = Position{Xint, Yint}

		i++
		j++
	}

	return behavior
}
