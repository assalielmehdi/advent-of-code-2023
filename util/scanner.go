package util

import (
	"log"
	"os"
	"strconv"
	"unicode"
)

type Scanner struct {
	bytes []byte
	idx   int
}

func NewScanner(bytes []byte) *Scanner {
	return &Scanner{
		bytes: bytes,
		idx:   0,
	}
}

func NewFileScanner(path string) *Scanner {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return NewScanner(bytes)
}

func (sc *Scanner) HasNext() bool {
	if sc.idx >= len(sc.bytes) {
		return false
	}

	nextIdx := sc.idx
	for nextIdx < len(sc.bytes) && unicode.IsSpace(rune(sc.bytes[nextIdx])) {
		nextIdx++
	}

	return nextIdx < len(sc.bytes)
}

func (sc *Scanner) skipSpace() {
	for sc.idx < len(sc.bytes) && unicode.IsSpace(rune(sc.bytes[sc.idx])) {
		sc.idx++
	}
}

func (sc *Scanner) NextBytes() []byte {
	if !sc.HasNext() {
		log.Fatalf("no more tokens to scan!")
	}

	sc.skipSpace()

	from := sc.idx

	for sc.idx < len(sc.bytes) && !unicode.IsSpace(rune(sc.bytes[sc.idx])) {
		sc.idx++
	}

	return sc.bytes[from:sc.idx]
}

func (sc *Scanner) Next() string {
	return string(sc.NextBytes())
}

func (sc *Scanner) NextInt() int {
	val, err := strconv.Atoi(sc.Next())
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func (sc *Scanner) HasNextLine() bool {
	return sc.idx < len(sc.bytes)
}

func (sc *Scanner) NextLineBytes() []byte {
	if !sc.HasNext() {
		log.Fatalf("no more tokens to scan!")
	}

	from := sc.idx

	for sc.idx < len(sc.bytes) && sc.bytes[sc.idx] != '\n' {
		sc.idx++
	}

	line := sc.bytes[from:sc.idx]

	sc.idx++

	return line
}

func (sc *Scanner) NextLine() string {
	return string(sc.NextLineBytes())
}
