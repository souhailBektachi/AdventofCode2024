package utiles

import (
	"bufio"

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
