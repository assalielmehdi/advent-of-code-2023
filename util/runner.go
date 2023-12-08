package util

import (
	"log"
	"os"
	"strings"
	"time"
)

func Run(title string, runner func(*Scanner) any, paths ...string) {
	for _, path := range paths {
		scanner := NewFileScanner(path)
		startAt := time.Now().UnixMicro()
		result := runner(scanner)

		log.Printf("Running %s on file: \033[1m%s\033[0m\t| Result: \033[1m%v\033[0m\t(took %vμs)", title, path, result, time.Now().UnixMicro()-startAt)
	}
}

func RunAll(title string, runner func(*Scanner) any) {
	Run(title, runner, getCwdFiles()...)
}

func getCwdFiles() []string {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	paths := make([]string, 0, len(files))
	for _, file := range files {
		path := file.Name()

		if strings.HasSuffix(path, ".txt") {
			paths = append(paths, path)
		}
	}

	return paths
}
