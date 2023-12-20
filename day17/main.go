package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pedrokiefer/adventofcode-2023/advent"
)

type BlockMap struct {
	Map map[advent.Point]int

	Width  int
	Height int
}

func InputToBlockMap(input io.ReadCloser) *BlockMap {
	bm := &BlockMap{
		Map: map[advent.Point]int{},
	}
	s := bufio.NewScanner(input)
	defer input.Close()
	y := 0
	for s.Scan() {
		l := s.Text()
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		for x, c := range l {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			bm.Map[advent.Point{X: x, Y: y}] = v
		}
		bm.Width = len(l)
		y++
	}
	bm.Height = y
	return bm
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	_ = InputToBlockMap(f)
}
