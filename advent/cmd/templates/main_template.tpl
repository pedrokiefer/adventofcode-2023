package main

import (
    "os"
    "log"
	"bufio"
	"io"
)

func InputTo{{ .Name }}(input io.ReadCloser) string {
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		l := s.Text()
	}
	return ""
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	_ = InputTo{{ .Name }}(f)
}
