package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Element string

var (
	EmptySpace         = Element(".")
	VerticalSplitter   = Element("|")
	HorizontalSplitter = Element("-")
	MirrorUpward       = Element("/")
	MirrorDownward     = Element("\\")
)

type Position struct {
	X int
	Y int
}

var (
	Up    = Position{X: 0, Y: -1}
	Down  = Position{X: 0, Y: 1}
	Left  = Position{X: -1, Y: 0}
	Right = Position{X: 1, Y: 0}
)

func (p Position) Add(d Position) Position {
	return Position{
		X: p.X + d.X,
		Y: p.Y + d.Y,
	}
}

type Directions []Position

func (d Directions) Contains(p Position) bool {
	for _, v := range d {
		if v == p {
			return true
		}
	}
	return false
}

type Cavern struct {
	Elements map[Position]Element

	Width  int
	Height int
}

type ScatterMap struct {
	Map            map[Position]Directions
	StartPosition  Position
	StartDirection Position
	Energized      int
}

func NewScatterMap() ScatterMap {
	return ScatterMap{
		Map:       map[Position]Directions{},
		Energized: 0,
	}
}

func InputToCavern(input io.ReadCloser) Cavern {
	cavern := Cavern{
		Elements: map[Position]Element{},
	}
	s := bufio.NewScanner(input)
	defer input.Close()
	y := 0
	for s.Scan() {
		l := s.Text()
		for i, c := range l {
			p := Position{
				X: i,
				Y: y,
			}
			cavern.Elements[p] = Element(string(c))
		}
		cavern.Width = len(l)
		y++
	}
	cavern.Height = y
	return cavern
}

func (sm *ScatterMap) ScatterLight(c Cavern, d Position, p Position) {
	if p.X < 0 || p.X >= c.Width || p.Y < 0 || p.Y >= c.Height {
		return
	}
	//fmt.Printf("Scattering light from direction %v on %v\n", d, p)
	if _, ok := sm.Map[p]; !ok {
		sm.Map[p] = Directions{d}
		sm.Energized++
	} else {
		if sm.Map[p].Contains(d) {
			//fmt.Printf("Already energized %v from %v\n", p, d)
			return
		}
		sm.Map[p] = append(sm.Map[p], d)
	}

	e := c.Elements[p]
	switch e {
	case EmptySpace:
		sm.ScatterLight(c, d, p.Add(d))
		return
	case VerticalSplitter:
		if d == Up || d == Down {
			sm.ScatterLight(c, d, p.Add(d))
			return
		}
		sm.ScatterLight(c, Up, p.Add(Up))
		sm.ScatterLight(c, Down, p.Add(Down))
		return
	case HorizontalSplitter:
		if d == Left || d == Right {
			sm.ScatterLight(c, d, p.Add(d))
			return
		}
		sm.ScatterLight(c, Left, p.Add(Left))
		sm.ScatterLight(c, Right, p.Add(Right))
		return
	case MirrorUpward:
		switch d {
		case Up:
			sm.ScatterLight(c, Right, p.Add(Right))
			return
		case Down:
			sm.ScatterLight(c, Left, p.Add(Left))
			return
		case Left:
			sm.ScatterLight(c, Down, p.Add(Down))
			return
		case Right:
			sm.ScatterLight(c, Up, p.Add(Up))
			return
		}
	case MirrorDownward:
		switch d {
		case Up:
			sm.ScatterLight(c, Left, p.Add(Left))
			return
		case Down:
			sm.ScatterLight(c, Right, p.Add(Right))
			return
		case Left:
			sm.ScatterLight(c, Up, p.Add(Up))
			return
		case Right:
			sm.ScatterLight(c, Down, p.Add(Down))
			return
		}
	}
}

func FindMaximumScatterMap(c Cavern) ScatterMap {
	maximum := ScatterMap{Energized: 0}

	for x := 0; x < c.Width; x++ {
		sm := NewScatterMap()
		sm.StartDirection = Down
		sm.StartPosition = Position{X: x, Y: 0}
		sm.ScatterLight(c, Down, Position{X: x, Y: 0})
		if sm.Energized > maximum.Energized {
			fmt.Printf("New maximum: %v position: %v direction: %v\n", sm.Energized, sm.StartPosition, sm.StartDirection)
			maximum = sm
		}

		sm = NewScatterMap()
		sm.StartDirection = Up
		sm.StartPosition = Position{X: x, Y: c.Height - 1}
		sm.ScatterLight(c, Up, Position{X: x, Y: c.Height - 1})
		if sm.Energized > maximum.Energized {
			fmt.Printf("New maximum: %v position: %v direction: %v\n", sm.Energized, sm.StartPosition, sm.StartDirection)
			maximum = sm
		}
	}

	for y := 0; y < c.Height; y++ {
		sm := NewScatterMap()
		sm.StartDirection = Right
		sm.StartPosition = Position{X: 0, Y: y}
		sm.ScatterLight(c, Right, Position{X: 0, Y: y})
		if sm.Energized > maximum.Energized {
			fmt.Printf("New maximum: %v position: %v direction: %v\n", sm.Energized, sm.StartPosition, sm.StartDirection)
			maximum = sm
		}

		sm = NewScatterMap()
		sm.StartDirection = Left
		sm.StartPosition = Position{X: c.Width - 1, Y: y}
		sm.ScatterLight(c, Left, Position{X: c.Width - 1, Y: y})
		if sm.Energized > maximum.Energized {
			fmt.Printf("New maximum: %v position: %v direction: %v\n", sm.Energized, sm.StartPosition, sm.StartDirection)
			maximum = sm
		}
	}

	return maximum
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := InputToCavern(f)
	sm := NewScatterMap()
	sm.ScatterLight(c, Right, Position{X: 0, Y: 0})

	fmt.Printf("Energized: %d\n", sm.Energized)

	maximum := FindMaximumScatterMap(c)
	fmt.Printf("Maximum: %d\n", maximum.Energized)
}
