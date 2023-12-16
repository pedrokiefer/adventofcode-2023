package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	Seeds                 = "seeds"
	SeedToSoil            = "seed-to-soil"
	SoilToFertilizer      = "soil-to-fertilizer"
	FertilizerToWater     = "fertilizer-to-water"
	WaterToLight          = "water-to-light"
	LightToTemperature    = "light-to-temperature"
	TemperatureToHumidity = "temperature-to-humidity"
	HumidityToLocation    = "humidity-to-location"
)

type Interval struct {
	Start  int64
	Length int64
}

type Range struct {
	DestinationStart int64
	SourceStart      int64
	Length           int64
}

type Almanac struct {
	Seeds                 []int64
	SeedToSoil            []Range
	SoilToFertilizer      []Range
	FertilizerToWater     []Range
	WaterToLight          []Range
	LightToTemperature    []Range
	TemperatureToHumidity []Range
	HumidityToLocation    []Range
}

func (a *Almanac) SetRange(name string, r []Range) {
	switch name {
	case SeedToSoil:
		a.SeedToSoil = r
	case SoilToFertilizer:
		a.SoilToFertilizer = r
	case FertilizerToWater:
		a.FertilizerToWater = r
	case WaterToLight:
		a.WaterToLight = r
	case LightToTemperature:
		a.LightToTemperature = r
	case TemperatureToHumidity:
		a.TemperatureToHumidity = r
	case HumidityToLocation:
		a.HumidityToLocation = r
	}
}

func InputToAlmanac(input io.ReadCloser) *Almanac {
	a := &Almanac{}
	s := bufio.NewScanner(input)
	defer input.Close()
	rangeName := ""
	curRange := []Range{}
	for s.Scan() {
		l := s.Text()
		l = strings.TrimSpace(l)
		if l == "" {
			a.SetRange(rangeName, curRange)
			curRange = []Range{}
			rangeName = ""
			continue
		}

		if strings.Contains(l, ":") {
			parts := strings.Split(l, ":")

			if parts[0] == Seeds {
				a.Seeds = NumbersToSlice(parts[1])
				continue
			}

			rangeName = strings.TrimSuffix(parts[0], " map")
		} else {
			parts := strings.Split(l, " ")
			if len(parts) != 3 {
				continue
			}
			d, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				continue
			}
			s, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				continue
			}
			l, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				continue
			}
			curRange = append(curRange, Range{
				DestinationStart: d,
				SourceStart:      s,
				Length:           l,
			})
		}

	}
	if rangeName != "" {
		a.SetRange(rangeName, curRange)
	}
	return a
}

func NumbersToSlice(numbers string) []int64 {
	results := []int64{}
	parts := strings.Split(strings.TrimSpace(numbers), " ")
	for _, n := range parts {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		v, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			continue
		}
		results = append(results, v)
	}
	return results
}

func (a Almanac) Map(v int64, rs []Range) int64 {
	for _, r := range rs {
		if v >= r.SourceStart && v < r.SourceStart+r.Length {
			return r.DestinationStart + (v - r.SourceStart)
		}
	}
	return v
}

func (a Almanac) Location(v int64) int64 {
	// fmt.Printf("Seed: %d\n", v)
	v = a.Map(v, a.SeedToSoil)
	// fmt.Printf("Seed to soil: %d\n", v)
	v = a.Map(v, a.SoilToFertilizer)
	// fmt.Printf("Soil to fertilizer: %d\n", v)
	v = a.Map(v, a.FertilizerToWater)
	// fmt.Printf("Fertilizer to water: %d\n", v)
	v = a.Map(v, a.WaterToLight)
	// fmt.Printf("Water to light: %d\n", v)
	v = a.Map(v, a.LightToTemperature)
	// fmt.Printf("Light to temperature: %d\n", v)
	v = a.Map(v, a.TemperatureToHumidity)
	// fmt.Printf("Temperature to humidity: %d\n", v)
	v = a.Map(v, a.HumidityToLocation)
	// fmt.Printf("Humidity to location: %d\n", v)
	return v
}

func FindLowestLocation(a *Almanac) int64 {
	lowest := int64(0)
	for _, s := range a.Seeds {
		l := a.Location(s)
		if lowest == 0 || l < lowest {
			lowest = l
		}
	}
	return lowest
}

func SeedsToIntervals(seeds []int64) []Interval {
	results := []Interval{}
	scp := make([]int64, len(seeds))
	if copy(scp, seeds) != len(seeds) {
		log.Fatal("copy failed")
		return []Interval{}
	}
	for len(scp) > 0 {
		s := scp[0]
		r := scp[1]
		results = append(results, Interval{
			Start:  s,
			Length: r,
		})
		fmt.Printf("Start Seed: %d range %d\n", s, r)
		scp = scp[2:]
	}
	return results
}

func PartitionInterval(i Interval, rs []Range) []Interval {
	results := []Interval{}
	iMin := i.Start
	iMax := i.Start + i.Length
	parts := []Interval{i}
	for len(parts) > 0 {
		for _, r := range rs {
			rMin := r.SourceStart
			rMax := r.SourceStart + r.Length
			if iMin > rMax || iMax < rMin {
				fmt.Printf("No overlap: %+v, %+v\n", i, r)
				continue
			}

			maxMin := max(iMin, rMin)
			minMax := min(iMax, rMax)
			ra := minMax - maxMin

			results = append(results, Interval{
				Start:  r.DestinationStart + (maxMin - r.SourceStart),
				Length: ra,
			})

		}
	}
	return results
}

func MapIntervals(intervals []Interval, rs []Range) []Interval {
	results := []Interval{}
	for _, i := range intervals {
		ni := PartitionInterval(i, rs)
		results = append(results, ni...)
	}
	return results
}

func FindLowestLocation2(a *Almanac) int64 {
	lowest := []int64{}
	intervals := SeedsToIntervals(a.Seeds)
	intervals = MapIntervals(intervals, a.SeedToSoil)
	fmt.Printf("New intervals: %+v\n", intervals)
	intervals = MapIntervals(intervals, a.SoilToFertilizer)
	fmt.Printf("New intervals: %+v\n", intervals)
	intervals = MapIntervals(intervals, a.FertilizerToWater)
	fmt.Printf("New intervals: %+v\n", intervals)
	intervals = MapIntervals(intervals, a.WaterToLight)
	fmt.Printf("New intervals: %+v\n", intervals)
	intervals = MapIntervals(intervals, a.LightToTemperature)
	fmt.Printf("New intervals: %+v\n", intervals)
	intervals = MapIntervals(intervals, a.TemperatureToHumidity)
	fmt.Printf("New intervals: %+v\n", intervals)
	intervals = MapIntervals(intervals, a.HumidityToLocation)
	fmt.Printf("New intervals: %+v\n", intervals)

	for _, i := range intervals {
		lowest = append(lowest, i.Start)
	}
	return lowest[0]
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	a := InputToAlmanac(f)

	lowest := FindLowestLocation(a)
	fmt.Printf("Lowest location: %d\n", lowest)

	lowest2 := FindLowestLocation2(a)
	fmt.Printf("Lowest location: %d\n", lowest2)
}
