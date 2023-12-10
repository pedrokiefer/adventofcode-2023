package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePuzzle(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)))

	m, symbols, parts := InputToPuzzleMap(input)

	fmt.Printf("%v\n", symbols)

	assert.Equal(t, m[Position{X: 0, Y: 0}], &Element{Value: 467, Symbol: false, StartPos: Position{X: 0, Y: 0}, EndPos: Position{X: 2, Y: 0}})

	assert.Equal(t, []*Element{
		{Symbol: true, SymbolC: '*', StartPos: Position{X: 3, Y: 1}, EndPos: Position{X: 3, Y: 1}},
		{Symbol: true, SymbolC: '#', StartPos: Position{X: 6, Y: 3}, EndPos: Position{X: 6, Y: 3}},
		{Symbol: true, SymbolC: '*', StartPos: Position{X: 3, Y: 4}, EndPos: Position{X: 3, Y: 4}},
		{Symbol: true, SymbolC: '+', StartPos: Position{X: 5, Y: 5}, EndPos: Position{X: 5, Y: 5}},
		{Symbol: true, SymbolC: '$', StartPos: Position{X: 3, Y: 8}, EndPos: Position{X: 3, Y: 8}},
		{Symbol: true, SymbolC: '*', StartPos: Position{X: 5, Y: 8}, EndPos: Position{X: 5, Y: 8}},
	}, symbols)

	assert.Equal(t, []*Element{
		{Value: 467, Symbol: false, StartPos: Position{X: 0, Y: 0}, EndPos: Position{X: 2, Y: 0}},
		{Value: 114, Symbol: false, StartPos: Position{X: 5, Y: 0}, EndPos: Position{X: 7, Y: 0}},
		{Value: 35, Symbol: false, StartPos: Position{X: 2, Y: 2}, EndPos: Position{X: 3, Y: 2}},
		{Value: 633, Symbol: false, StartPos: Position{X: 6, Y: 2}, EndPos: Position{X: 8, Y: 2}},
		{Value: 617, Symbol: false, StartPos: Position{X: 0, Y: 4}, EndPos: Position{X: 2, Y: 4}},
		{Value: 58, Symbol: false, StartPos: Position{X: 7, Y: 5}, EndPos: Position{X: 8, Y: 5}},
		{Value: 592, Symbol: false, StartPos: Position{X: 2, Y: 6}, EndPos: Position{X: 4, Y: 6}},
		{Value: 755, Symbol: false, StartPos: Position{X: 6, Y: 7}, EndPos: Position{X: 8, Y: 7}},
		{Value: 664, Symbol: false, StartPos: Position{X: 1, Y: 9}, EndPos: Position{X: 3, Y: 9}},
		{Value: 598, Symbol: false, StartPos: Position{X: 5, Y: 9}, EndPos: Position{X: 7, Y: 9}},
	}, parts)

	CheckValid(m, symbols)

	sum := int64(0)
	for _, e := range parts {
		if e.Valid {
			sum += e.Value
		}
	}

	assert.Equal(t, int64(4361), sum)
}

func TestFindGears(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)))

	m, symbols, parts := InputToPuzzleMap(input)

	assert.Equal(t, m[Position{X: 0, Y: 0}], &Element{Value: 467, Symbol: false, StartPos: Position{X: 0, Y: 0}, EndPos: Position{X: 2, Y: 0}})

	assert.Equal(t, []*Element{
		{Symbol: true, SymbolC: '*', StartPos: Position{X: 3, Y: 1}, EndPos: Position{X: 3, Y: 1}},
		{Symbol: true, SymbolC: '#', StartPos: Position{X: 6, Y: 3}, EndPos: Position{X: 6, Y: 3}},
		{Symbol: true, SymbolC: '*', StartPos: Position{X: 3, Y: 4}, EndPos: Position{X: 3, Y: 4}},
		{Symbol: true, SymbolC: '+', StartPos: Position{X: 5, Y: 5}, EndPos: Position{X: 5, Y: 5}},
		{Symbol: true, SymbolC: '$', StartPos: Position{X: 3, Y: 8}, EndPos: Position{X: 3, Y: 8}},
		{Symbol: true, SymbolC: '*', StartPos: Position{X: 5, Y: 8}, EndPos: Position{X: 5, Y: 8}},
	}, symbols)

	assert.Equal(t, []*Element{
		{Value: 467, Symbol: false, StartPos: Position{X: 0, Y: 0}, EndPos: Position{X: 2, Y: 0}},
		{Value: 114, Symbol: false, StartPos: Position{X: 5, Y: 0}, EndPos: Position{X: 7, Y: 0}},
		{Value: 35, Symbol: false, StartPos: Position{X: 2, Y: 2}, EndPos: Position{X: 3, Y: 2}},
		{Value: 633, Symbol: false, StartPos: Position{X: 6, Y: 2}, EndPos: Position{X: 8, Y: 2}},
		{Value: 617, Symbol: false, StartPos: Position{X: 0, Y: 4}, EndPos: Position{X: 2, Y: 4}},
		{Value: 58, Symbol: false, StartPos: Position{X: 7, Y: 5}, EndPos: Position{X: 8, Y: 5}},
		{Value: 592, Symbol: false, StartPos: Position{X: 2, Y: 6}, EndPos: Position{X: 4, Y: 6}},
		{Value: 755, Symbol: false, StartPos: Position{X: 6, Y: 7}, EndPos: Position{X: 8, Y: 7}},
		{Value: 664, Symbol: false, StartPos: Position{X: 1, Y: 9}, EndPos: Position{X: 3, Y: 9}},
		{Value: 598, Symbol: false, StartPos: Position{X: 5, Y: 9}, EndPos: Position{X: 7, Y: 9}},
	}, parts)

	gears := FindGears(m, symbols)

	assert.Equal(t, []*Gear{
		{
			Part1: &Element{Value: 35, Symbol: false, StartPos: Position{X: 2, Y: 2}, EndPos: Position{X: 3, Y: 2}},
			Part2: &Element{Value: 467, Symbol: false, StartPos: Position{X: 0, Y: 0}, EndPos: Position{X: 2, Y: 0}},
			Ratio: int64(35 * 467),
		},
		{
			Part1: &Element{Value: 755, Symbol: false, StartPos: Position{X: 6, Y: 7}, EndPos: Position{X: 8, Y: 7}},
			Part2: &Element{Value: 598, Symbol: false, StartPos: Position{X: 5, Y: 9}, EndPos: Position{X: 7, Y: 9}},
			Ratio: int64(755 * 598),
		},
	}, gears)
}
