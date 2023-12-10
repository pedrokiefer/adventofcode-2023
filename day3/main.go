package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

type Element struct {
	Value    int64
	Symbol   bool
	SymbolC  rune
	StartPos Position
	EndPos   Position
	Valid    bool
}

func (e Element) String() string {
	if e.Symbol {
		return string(e.SymbolC) + " " + e.StartPos.String() + " " + e.EndPos.String()
	}
	return strconv.FormatInt(e.Value, 10) + " " + e.StartPos.String() + " " + e.EndPos.String()
}

type PuzzleMap map[Position]*Element

func (pMap PuzzleMap) AddToPuzzleMap(e *Element) {
	for x := e.StartPos.X; x <= e.EndPos.X; x++ {
		p := Position{
			X: x,
			Y: e.StartPos.Y,
		}
		pMap[p] = e
	}
}

type Gear struct {
	Part1 *Element
	Part2 *Element
	Ratio int64
}

func (g Gear) String() string {
	return strconv.FormatInt(g.Ratio, 10)
}

func InputToPuzzleMap(input io.ReadCloser) (PuzzleMap, []*Element, []*Element) {
	symbols := []*Element{}
	parts := []*Element{}
	pMap := PuzzleMap{}
	s := bufio.NewScanner(input)
	defer input.Close()
	y := 0
	for s.Scan() {
		l := s.Text()
		var lastDigit *rune
		var curValue string
		var E *Element
		E = nil
		for x, c := range l {
			if lastDigit == nil {
				if c >= '0' && c <= '9' {
					lastDigit = &c
					curValue = string(c)
					E = &Element{
						Symbol: false,
						StartPos: Position{
							X: x,
							Y: y,
						},
					}
				} else if c != '.' {
					p := Position{
						X: x,
						Y: y,
					}
					E = &Element{
						Symbol:   true,
						SymbolC:  c,
						StartPos: p,
						EndPos:   p,
					}
					lastDigit = nil
					log.Printf("Found Symbol %c at %d,%d", c, x, y)
					symbols = append(symbols, E)
				}
				continue
			} else {
				if c >= '0' && c <= '9' {
					curValue += string(c)
					lastDigit = &c
				} else {
					lastDigit = nil
					v, err := strconv.ParseInt(curValue, 10, 64)
					if err != nil {
						continue
					}
					log.Printf("Found value %d at %d,%d", v, x-1, y)
					E.Value = v
					E.EndPos = Position{
						X: x - 1,
						Y: y,
					}
					parts = append(parts, E)

					pMap.AddToPuzzleMap(E)

					if c != '.' {
						p := Position{
							X: x,
							Y: y,
						}
						E = &Element{
							Symbol:   true,
							SymbolC:  c,
							StartPos: p,
							EndPos:   p,
						}
						log.Printf("Found Symbol %c at %d,%d", c, x, y)
						symbols = append(symbols, E)
					}
				}
			}
		}
		if lastDigit != nil {
			lastDigit = nil
			v, err := strconv.ParseInt(curValue, 10, 64)
			if err != nil {
				continue
			}
			log.Printf("Found value %d at %d,%d", v, len(l)-1, y)
			E.Value = v
			E.EndPos = Position{
				X: len(l) - 1,
				Y: y,
			}
			parts = append(parts, E)
			pMap.AddToPuzzleMap(E)
		}
		y++
	}
	return pMap, symbols, parts
}

func CheckValid(pMap PuzzleMap, symbols []*Element) bool {
	for _, s := range symbols {
		if e, ok := pMap[Position{X: s.StartPos.X, Y: s.StartPos.Y - 1}]; ok {
			e.Valid = true
		}
		if e, ok := pMap[Position{X: s.StartPos.X + 1, Y: s.StartPos.Y - 1}]; ok {
			e.Valid = true
		}
		if e, ok := pMap[Position{X: s.StartPos.X + 1, Y: s.StartPos.Y}]; ok {
			e.Valid = true
		}
		if e, ok := pMap[Position{X: s.StartPos.X + 1, Y: s.StartPos.Y + 1}]; ok {
			e.Valid = true
		}
		if e, ok := pMap[Position{X: s.StartPos.X, Y: s.StartPos.Y + 1}]; ok {
			e.Valid = true
		}
		if e, ok := pMap[Position{X: s.StartPos.X - 1, Y: s.StartPos.Y + 1}]; ok {
			e.Valid = true
		}
		if e, ok := pMap[Position{X: s.StartPos.X - 1, Y: s.StartPos.Y}]; ok {
			e.Valid = true
		}
		if e, ok := pMap[Position{X: s.StartPos.X - 1, Y: s.StartPos.Y - 1}]; ok {
			e.Valid = true
		}
	}
	return true
}

func ElementInSlice(e *Element, list []*Element) bool {
	for _, l := range list {
		if l == e {
			return true
		}
	}
	return false
}

func FindGears(pMap PuzzleMap, symbols []*Element) []*Gear {
	result := []*Gear{}
	for _, s := range symbols {
		if s.SymbolC != '*' {
			continue
		}
		parts := []*Element{}
		if e, ok := pMap[Position{X: s.StartPos.X, Y: s.StartPos.Y - 1}]; ok {
			if !ElementInSlice(e, parts) {
				parts = append(parts, e)
			}
		}
		if e, ok := pMap[Position{X: s.StartPos.X + 1, Y: s.StartPos.Y - 1}]; ok {
			if !ElementInSlice(e, parts) {
				parts = append(parts, e)
			}
		}
		if e, ok := pMap[Position{X: s.StartPos.X + 1, Y: s.StartPos.Y}]; ok {
			if !ElementInSlice(e, parts) {
				parts = append(parts, e)
			}
		}
		if e, ok := pMap[Position{X: s.StartPos.X + 1, Y: s.StartPos.Y + 1}]; ok {
			if !ElementInSlice(e, parts) {
				parts = append(parts, e)
			}
		}
		if e, ok := pMap[Position{X: s.StartPos.X, Y: s.StartPos.Y + 1}]; ok {
			if !ElementInSlice(e, parts) {
				parts = append(parts, e)
			}
		}
		if e, ok := pMap[Position{X: s.StartPos.X - 1, Y: s.StartPos.Y + 1}]; ok {
			if !ElementInSlice(e, parts) {
				parts = append(parts, e)
			}
		}
		if e, ok := pMap[Position{X: s.StartPos.X - 1, Y: s.StartPos.Y}]; ok {
			if !ElementInSlice(e, parts) {
				parts = append(parts, e)
			}
		}
		if e, ok := pMap[Position{X: s.StartPos.X - 1, Y: s.StartPos.Y - 1}]; ok {
			if !ElementInSlice(e, parts) {
				parts = append(parts, e)
			}
		}

		if len(parts) == 2 {
			g := &Gear{
				Part1: parts[0],
				Part2: parts[1],
				Ratio: parts[0].Value * parts[1].Value,
			}
			result = append(result, g)
		}
	}
	return result
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	m, symbols, parts := InputToPuzzleMap(f)

	CheckValid(m, symbols)

	sum := int64(0)
	for _, e := range parts {
		if e.Valid {
			sum += e.Value
		}
	}

	log.Printf("Sum of valid parts: %d", sum)

	gears := FindGears(m, symbols)

	gearSum := int64(0)
	for _, g := range gears {
		gearSum += g.Ratio
	}

	log.Printf("Sum of gears ratio: %d", gearSum)
}
