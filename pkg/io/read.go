package util_io

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func ReadLines(filepath string) [][]byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		log.Fatal(err)
	}

	lines := make([][]byte, 0)
	for scanner.Scan() {
		lines = append(lines, slices.Clone(scanner.Bytes()))
	}

	return lines
}
