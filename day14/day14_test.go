package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/pedrokiefer/adventofcode-2023/advent"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`)))

	v := InputToRocksMap(input)

	assert.Equal(t, Column{"O", "O", ".", "O", ".", "O", ".", ".", "#", "#"}, v.Columns[0])
	assert.Equal(t, Column{".", "O", ".", ".", ".", "#", "O", ".", ".", "O"}, v.Columns[2])

	v.RollNorth()
	v.Roll(North)
	assert.Equal(t, Column{"O", "O", "O", "O", ".", ".", ".", ".", "#", "#"}, v.Columns[0])

	load := v.CalculateLoad()
	assert.Equal(t, 136, load)
}

func TestRollNorth(t *testing.T) {
	c := Column{"O", "O", ".", "O", ".", "O", ".", ".", "#", "#"}
	c.RollNorth()

	assert.Equal(t, Column{"O", "O", "O", "O", ".", ".", ".", ".", "#", "#"}, c)

	c2 := Column{".", "O", ".", ".", ".", "#", "O", ".", ".", "O"}
	c2.RollNorth()

	assert.Equal(t, Column{"O", ".", ".", ".", ".", "#", "O", "O", ".", "."}, c2)
}

func TestRoll(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`)))

	v := InputToRocksMap(input)

	assert.Equal(t, "O", v.Map[advent.Point{X: 0, Y: 0}])
	assert.Equal(t, "#", v.Map[advent.Point{X: 5, Y: 0}])
	assert.Equal(t, ".", v.Map[advent.Point{X: 0, Y: 2}])

	v.Roll(North)
	assert.Equal(t, "O", v.Map[advent.Point{X: 0, Y: 2}])

	v.Print()

	assert.Equal(t, "O", v.Map[advent.Point{X: 4, Y: 2}])
	v.Roll(West)
	assert.Equal(t, ".", v.Map[advent.Point{X: 4, Y: 2}])
	v.Print()

	assert.Equal(t, ".", v.Map[advent.Point{X: 2, Y: 4}])
	v.Roll(South)
	assert.Equal(t, "O", v.Map[advent.Point{X: 2, Y: 4}])
	v.Print()

	assert.Equal(t, ".", v.Map[advent.Point{X: 3, Y: 9}])
	v.Roll(East)
	assert.Equal(t, "O", v.Map[advent.Point{X: 3, Y: 9}])
	v.Print()
}

func TestCycle(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`)))

	v := InputToRocksMap(input)
	v.Cycle()
	assert.Equal(t, ".", v.Map[advent.Point{X: 0, Y: 0}])
	assert.Equal(t, "O", v.Map[advent.Point{X: 3, Y: 9}])
}

func TestRunLongCycles(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`)))

	v := InputToRocksMap(input)
	load := v.RunLongCycles()
	assert.Equal(t, ".", v.Map[advent.Point{X: 0, Y: 0}])
	assert.Equal(t, "O", v.Map[advent.Point{X: 3, Y: 9}])
	assert.Equal(t, 64, load)
}
