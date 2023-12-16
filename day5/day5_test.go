package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`)))

	a := InputToAlmanac(input)

	assert.Equal(t, []int64{79, 14, 55, 13}, a.Seeds)
	assert.Equal(t, []Range{
		{DestinationStart: 50, SourceStart: 98, Length: 2},
		{DestinationStart: 52, SourceStart: 50, Length: 48},
	}, a.SeedToSoil)

	v := a.Map(int64(79), a.SeedToSoil)
	assert.Equal(t, int64(81), v)

	v = a.Location(int64(79))
	assert.Equal(t, int64(82), v)

	l := FindLowestLocation2(a)
	assert.Equal(t, int64(46), l)
}

func TestPartitionInterval(t *testing.T) {
	ranges := []Range{
		{DestinationStart: 49, SourceStart: 53, Length: 8},
		{DestinationStart: 0, SourceStart: 11, Length: 42},
		{DestinationStart: 42, SourceStart: 0, Length: 7},
		{DestinationStart: 57, SourceStart: 7, Length: 4},
	}
	interval := Interval{Start: 81, Length: 14}
	partitioned := PartitionInterval(interval, ranges)
	assert.Equal(t, []Interval{
		{Start: 81, Length: 14},
	}, partitioned)
}

func TestPartitionInterval2(t *testing.T) {
	ranges := []Range{
		{DestinationStart: 60, SourceStart: 20, Length: 10},
		{DestinationStart: 20, SourceStart: 40, Length: 10},
		{DestinationStart: 40, SourceStart: 60, Length: 10},
	}
	interval := Interval{Start: 10, Length: 80}
	partitioned := PartitionInterval(interval, ranges)
	assert.Equal(t, []Interval{
		{Start: 10, Length: 10},
		{Start: 20, Length: 10},
		{Start: 30, Length: 10},
		{Start: 40, Length: 10},
		{Start: 50, Length: 10},
		{Start: 60, Length: 10},
		{Start: 70, Length: 10},
	}, partitioned)
}
