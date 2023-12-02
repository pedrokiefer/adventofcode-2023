package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanReadAGame(t *testing.T) {
	g, err := ReadGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), g.ID)
	assert.Equal(t, []Set{
		{Red: 4, Green: 0, Blue: 3},
		{Red: 1, Green: 2, Blue: 6},
		{Red: 0, Green: 2, Blue: 0},
	}, g.Sets)
}

func TestFileToGame(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`)))

	list := InputToGame(input)

	assert.Equal(t, []Game{
		{ID: 1, Sets: []Set{
			{Red: 4, Green: 0, Blue: 3},
			{Red: 1, Green: 2, Blue: 6},
			{Red: 0, Green: 2, Blue: 0},
		}},
		{ID: 2, Sets: []Set{
			{Red: 0, Green: 2, Blue: 1},
			{Red: 1, Green: 3, Blue: 4},
			{Red: 0, Green: 1, Blue: 1},
		}},
		{ID: 3, Sets: []Set{
			{Red: 20, Green: 8, Blue: 6},
			{Red: 4, Green: 13, Blue: 5},
			{Red: 1, Green: 5, Blue: 0},
		}},
		{ID: 4, Sets: []Set{
			{Red: 3, Green: 1, Blue: 6},
			{Red: 6, Green: 3, Blue: 0},
			{Red: 14, Green: 3, Blue: 15},
		}},
		{ID: 5, Sets: []Set{
			{Red: 6, Green: 3, Blue: 1},
			{Red: 1, Green: 2, Blue: 2},
		}},
	}, list)
}

func TestCanGetMinimalSet(t *testing.T) {
	g, err := ReadGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), g.ID)
	assert.Equal(t, []Set{
		{Red: 4, Green: 0, Blue: 3},
		{Red: 1, Green: 2, Blue: 6},
		{Red: 0, Green: 2, Blue: 0},
	}, g.Sets)
	assert.Equal(t, Set{Red: 4, Green: 2, Blue: 6}, g.MinimalSet())
	assert.Equal(t, int64(48), g.MinimalSet().Power())
}

func TestCanGetMinimalSet2(t *testing.T) {
	g, err := ReadGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	assert.Nil(t, err)
	assert.Equal(t, int64(3), g.ID)
	assert.Equal(t, []Set{
		{Red: 20, Green: 8, Blue: 6},
		{Red: 4, Green: 13, Blue: 5},
		{Red: 1, Green: 5, Blue: 0},
	}, g.Sets)
	assert.Equal(t, Set{Red: 20, Green: 13, Blue: 6}, g.MinimalSet())
}
