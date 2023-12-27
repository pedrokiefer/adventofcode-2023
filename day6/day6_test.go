package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`Time:      7  15   30
Distance:  9  40  200`)))

	v := InputToRace(input)

	assert.Equal(t, []Race{
		{Time: 7, RecordDistance: 9},
		{Time: 15, RecordDistance: 40},
		{Time: 30, RecordDistance: 200},
	}, v)

	assert.Equal(t, 4, v[0].CountPossibleRecords())
	assert.Equal(t, 8, v[1].CountPossibleRecords())
	assert.Equal(t, 9, v[2].CountPossibleRecords())
}

func Test2(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`Time:      7  15   30
Distance:  9  40  200`)))

	v := InputToRace2(input)

	assert.Equal(t, Race{
		Time:           71530,
		RecordDistance: 940200,
	}, v)

	assert.Equal(t, 71503, v.CountPossibleRecords())
}
