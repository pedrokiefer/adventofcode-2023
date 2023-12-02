package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileToIntList(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`)))

	list := InputToIntList(input, GetCalibrationValue)

	assert.Equal(t, []int64{12, 38, 15, 77}, list)
}

func TestFileToIntList2(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`)))

	list := InputToIntList(input, GetCalibrationValue2)

	assert.Equal(t, []int64{29, 83, 13, 24, 42, 14, 76}, list)
}

func TestFindFirstDigit(t *testing.T) {
	d := FindFirstDigit("1abc2")
	assert.Equal(t, "1", d)
}

func TestFindLastDigit(t *testing.T) {
	d := FindLastDigit("1abc2")
	assert.Equal(t, "2", d)
}

func TestGetCalibrationValue(t *testing.T) {
	v, err := GetCalibrationValue("1abc2")
	assert.Nil(t, err)
	assert.Equal(t, int64(12), v)
}

func TestFindFirstDigit2(t *testing.T) {
	d := FindFirstDigit2("two1nine")
	assert.Equal(t, "2", d)
}

func TestFindLastDigit2(t *testing.T) {
	d := FindLastDigit2("two1nine")
	assert.Equal(t, "9", d)
}

func TestGetCalibrationValue2(t *testing.T) {
	v, err := GetCalibrationValue2("two1nine")
	assert.Nil(t, err)
	assert.Equal(t, int64(29), v)
}
