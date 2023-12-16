package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCavern(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`)))
	c := InputToCavern(input)

	assert.Equal(t, 10, c.Width)
	assert.Equal(t, 10, c.Height)
	assert.Equal(t, VerticalSplitter, c.Elements[Position{X: 1, Y: 0}])
}

func TestScatterLight(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`)))
	c := InputToCavern(input)

	assert.Equal(t, 10, c.Width)
	assert.Equal(t, 10, c.Height)
	assert.Equal(t, VerticalSplitter, c.Elements[Position{X: 1, Y: 0}])

	sm := NewScatterMap()
	sm.ScatterLight(c, Right, Position{X: 0, Y: 0})

	assert.Equal(t, Directions{Right}, sm.Map[Position{X: 0, Y: 0}])
	assert.Equal(t, Directions{Right, Left}, sm.Map[Position{X: 1, Y: 0}])
	assert.Equal(t, Directions{Down}, sm.Map[Position{X: 1, Y: 1}])
	assert.Equal(t, Directions{Left}, sm.Map[Position{X: 0, Y: 7}])
	assert.Equal(t, Directions{Down, Up}, sm.Map[Position{X: 1, Y: 7}])
	assert.Equal(t, Directions{Right}, sm.Map[Position{X: 4, Y: 7}])
	assert.Equal(t, Directions{Up}, sm.Map[Position{X: 4, Y: 6}])
	assert.Equal(t, Directions{Right, Down}, sm.Map[Position{X: 5, Y: 6}])
	assert.Equal(t, Directions{Right, Left}, sm.Map[Position{X: 6, Y: 6}])

	assert.Equal(t, 46, sm.Energized)
}

func TestScatterLightMaximum(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`)))
	c := InputToCavern(input)

	assert.Equal(t, 10, c.Width)
	assert.Equal(t, 10, c.Height)
	assert.Equal(t, VerticalSplitter, c.Elements[Position{X: 1, Y: 0}])

	sm := FindMaximumScatterMap(c)

	assert.Equal(t, Position{X: 3, Y: 0}, sm.StartPosition)
	assert.Equal(t, Down, sm.StartDirection)
	assert.Equal(t, 51, sm.Energized)

}
