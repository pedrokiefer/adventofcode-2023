package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pedrokiefer/adventofcode-2023/advent"
)

var (
	RoundRock = string("O")
	CubeRock  = string("#")
	EmpySpace = string(".")
)

type Column []string

func (c Column) RollNorth() {
	for i := 0; i < len(c); i++ {
		if c[i] == CubeRock || c[i] == RoundRock {
			continue
		}
		if c[i] == EmpySpace {
			nextRock := c.nextRock(i)
			if nextRock == -1 {
				continue
			}
			c[i], c[nextRock] = c[nextRock], c[i]
		}
	}
}

func (c Column) nextRock(i int) int {
	for j := i + 1; j < len(c); j++ {
		if c[j] == CubeRock {
			return -1
		}
		if c[j] == RoundRock {
			return j
		}
	}
	return -1
}

type RocksMap struct {
	Columns []Column
	Map     map[advent.Point]string
	Width   int
	Height  int
}

func (r *RocksMap) RollNorth() {
	for i := range r.Columns {
		r.Columns[i].RollNorth()
	}
}

type Direction int

const (
	North Direction = iota
	West
	South
	East
)

var counterClock = []Direction{North, West, South, East}

func (r *RocksMap) Cycle() {
	for _, d := range counterClock {
		r.Roll(d)
	}
}

func (r *RocksMap) Roll(d Direction) {
	switch d {
	case North:
		for i := 0; i < r.Width; i++ {
			r.rollLine(r.Width, func(j int) advent.Point {
				return advent.Point{X: i, Y: j}
			})
		}
	case West:
		for i := 0; i < r.Height; i++ {
			r.rollLine(r.Height, func(j int) advent.Point {
				return advent.Point{X: j, Y: i}
			})
		}
	case South:
		for i := r.Width; i >= 0; i-- {
			r.rollLine(r.Width, func(j int) advent.Point {
				return advent.Point{X: r.Height - i - 1, Y: r.Width - j - 1}
			})
		}
	case East:
		for i := r.Height; i >= 0; i-- {
			r.rollLine(r.Height, func(j int) advent.Point {
				return advent.Point{X: r.Height - j - 1, Y: r.Width - i - 1}
			})
		}
	}
}

func (r *RocksMap) rollLine(size int, mkPoint func(int) advent.Point) {
	for i := 0; i < size; i++ {
		p := mkPoint(i)
		if r.Map[p] == CubeRock || r.Map[p] == RoundRock {
			continue
		}
		if r.Map[p] == EmpySpace {
			nextRock := r.nextRock(i, size, mkPoint)
			if nextRock == -1 {
				continue
			}
			np := mkPoint(nextRock)
			r.Map[p], r.Map[np] = r.Map[np], r.Map[p]
		}
	}
}

func (r RocksMap) nextRock(i, size int, mkPoint func(int) advent.Point) int {
	for j := i + 1; j < size; j++ {
		p := mkPoint(j)
		if r.Map[p] == CubeRock {
			return -1
		}
		if r.Map[p] == RoundRock {
			return j
		}
	}
	return -1
}

func (r *RocksMap) Hash() string {
	h := sha256.New()
	for i := 0; i < r.Height; i++ {
		for j := 0; j < r.Width; j++ {
			p := advent.Point{X: j, Y: i}
			h.Write([]byte(p.String()))
			h.Write([]byte(r.Map[p]))
		}
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (r *RocksMap) RunLongCycles() int {
	type cycle struct {
		Hash       string
		Load       int
		Interation int
	}
	seen := map[string]cycle{}
	seenList := []cycle{}
	cycleLength := 0
	i := 0
	h := ""
	for ; i < 1000; i++ {
		r.Cycle()
		h = r.Hash()
		l := r.CalculateLoad()
		fmt.Printf("Cycle: %d Load: %d Hash: %s\n", i, l, h)
		if _, ok := seen[h]; !ok {
			c := cycle{
				Hash:       h,
				Load:       l,
				Interation: i,
			}
			seenList = append(seenList, c)
			seen[h] = c
			continue
		}
		break
	}
	// We found a cycle
	cycleStart := seen[h]
	cycleLength = i - cycleStart.Interation
	fmt.Printf("Cycle: %d Lenght: %d\n", i, cycleLength)
	fmt.Printf("Skip: %d \n", (1000000000 - i))
	m := (1000000000 - i) % cycleLength
	fmt.Printf("m: %d\n", m)
	return seenList[cycleStart.Interation+m-1].Load
}

func (r *RocksMap) CalculateLoad() int {
	load := 0
	for i := 0; i < r.Height; i++ {
		for j := 0; j < r.Width; j++ {
			p := advent.Point{X: j, Y: i}
			if r.Map[p] != RoundRock {
				continue
			}
			load += r.Height - i
		}
	}
	return load
}

func (r RocksMap) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", r.Width))
	for i := 0; i < r.Height; i++ {
		for j := 0; j < r.Width; j++ {
			fmt.Print(r.Map[advent.Point{X: j, Y: i}])
		}
		fmt.Println()
	}
	fmt.Printf("%s\n", strings.Repeat("-", r.Width))
}

func InputToRocksMap(input io.ReadCloser) RocksMap {
	s := bufio.NewScanner(input)
	defer input.Close()
	cols := []Column{}
	m := map[advent.Point]string{}
	i := 0
	for s.Scan() {
		l := s.Text()
		if i == 0 {
			cols = make([]Column, len(l))
		}
		for j, c := range l {
			m[advent.Point{X: j, Y: i}] = string(c)
			if c == '#' {
				cols[j] = append(cols[j], CubeRock)
			} else if c == '.' {
				cols[j] = append(cols[j], EmpySpace)
			} else {
				cols[j] = append(cols[j], RoundRock)
			}
		}
		i++
	}
	return RocksMap{
		Columns: cols,
		Map:     m,
		Width:   len(cols[0]),
		Height:  len(cols),
	}
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	rm := InputToRocksMap(f)

	rm.Roll(North)
	fmt.Printf("Load: %d\n", rm.CalculateLoad())

	f, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	rm2 := InputToRocksMap(f)

	load := rm2.RunLongCycles()
	fmt.Printf("Load: %d\n", load)
}
