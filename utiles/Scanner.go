package utiles

import (
	"bufio"
	"fmt"
	"time"

	"os"
)

func ReturnScanner(path string) (*bufio.Scanner, error, *os.File) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err, nil
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, err, nil
	} else {
		return scanner, nil, file

	}

}
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

type Position struct {
	X int
	Y int
}
