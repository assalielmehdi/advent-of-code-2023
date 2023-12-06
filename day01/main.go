package main

import (
	"fmt"
	"unicode"

	util_io "assalielmehdi/adventofcode2023/pkg/io"
)

func compareSubstr(str string, i int, substr string) bool {
	if i >= len(str) || i+len(substr) > len(str) {
		return false
	}

	ok := true

	for j := 0; ok && j < len(substr); j++ {
		ok = ok && str[i+j] == substr[j]
	}

	return ok
}

func scanNextDigit(str string, i int) int {
	if i >= len(str) {
		return -1
	}

	if unicode.IsDigit(rune(str[i])) {
		return int(rune(str[i]) - '0')
	}

	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for j, dig := range digits {
		if compareSubstr(str, i, dig) {
			return j
		}
	}

	return -1
}

func calculateCalibration(line string) int {
	firstDig, lastDig := -1, -1

	for i := range line {
		dig := scanNextDigit(line, i)

		if dig != -1 {
			if firstDig == -1 {
				firstDig = dig
			}

			lastDig = dig
		}
	}

	return firstDig*10 + lastDig
}

func main() {
	lines := util_io.ReadLines("input2.txt")

	sum := 0

	for _, line := range lines {
		sum += calculateCalibration(string(line))
	}

	fmt.Println(sum)
}
