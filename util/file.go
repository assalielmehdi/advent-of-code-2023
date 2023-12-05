package util

import (
	"bufio"
	"log"
	"os"
)

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type FileIterator struct {
	file    os.File
	scanner bufio.Scanner
}

func NewFileIterator(filepath string) *FileIterator {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		log.Fatal(err)
	}

	return &FileIterator{
		file:    *file,
		scanner: *scanner,
	}
}

func (it *FileIterator) HasNext() bool {
	return it.scanner.Scan()
}

func (it *FileIterator) Next() string {
	return it.scanner.Text()
}

func (it *FileIterator) Close() {
	it.file.Close()
}
