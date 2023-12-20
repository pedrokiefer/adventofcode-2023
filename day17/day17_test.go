package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/pedrokiefer/adventofcode-2023/advent"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`)))

	v := InputToBlockMap(input)

	assert.Equal(t, 2, v.Map[advent.Point{X: 0, Y: 0}])
	assert.Equal(t, 13, v.Width)
	assert.Equal(t, 13, v.Height)
	assert.Equal(t, 3, v.Map[advent.Point{X: 12, Y: 12}])
}
