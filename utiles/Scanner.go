package utiles

import (
	"bufio"

	"os"
)

func ReturnScanner(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, err
	} else {
		return scanner, nil

	}

}
