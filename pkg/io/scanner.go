package util_io

import (
	"log"
	"strconv"
)

type Scanner struct {
	tokenizer *Tokenizer
}

func NewScanner(content []byte) *Scanner {
	return &Scanner{
		tokenizer: NewTokenizer(content),
	}
}

func NewEmptyScanner() *Scanner {
	return &Scanner{
		tokenizer: NewEmptyTokenizer(),
	}
}

func (scanner *Scanner) SetContent(content []byte) {
	scanner.tokenizer.SetContent(content)
}

func (scanner *Scanner) HasNext() bool {
	return scanner.tokenizer.HasNext()
}

func (scanner *Scanner) Next() []byte {
	return scanner.tokenizer.Next()
}

func (scanner *Scanner) NextInt() int {
	valBytes := scanner.tokenizer.Next()
	if valBytes == nil {
		log.Fatalf("nothing to scan")
	}

	val, err := strconv.Atoi(string(valBytes))
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func (scanner *Scanner) NextInts() []int {
	ints := make([]int, 0)

	for scanner.tokenizer.HasNext() {
		valBytes := scanner.tokenizer.Next()

		val, err := strconv.Atoi(string(valBytes))
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, val)
	}

	return ints
}
