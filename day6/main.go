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

type Race struct {
	Time           int
	RecordDistance int
}

func (r Race) CountPossibleRecords() int {
	count := 0
	max := 0
	add := 0
	if r.Time%2 == 0 {
		max = r.Time/2 + 1
		add = -1
	} else {
		max = (r.Time + 1) / 2
	}
	for i := 0; i < max; i++ {
		d := i * (r.Time - i)
		if d > r.RecordDistance {
			count++
		}
	}
	return count*2 + add
}

func InputToRace(input io.ReadCloser) []Race {
	s := bufio.NewScanner(input)
	defer input.Close()
	time := []int{}
	distance := []int{}
	for s.Scan() {
		l := s.Text()
		if strings.Contains(l, "Time") {
			ts := strings.TrimLeft(l, "Time:")
			ts = strings.TrimSpace(ts)
			for _, v := range strings.Split(ts, " ") {
				v = strings.TrimSpace(v)
				if v == "" {
					continue
				}
				i, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				time = append(time, i)
			}
		} else if strings.Contains(l, "Distance") {
			ts := strings.TrimLeft(l, "Distance:")
			ts = strings.TrimSpace(ts)
			for _, v := range strings.Split(ts, " ") {
				v = strings.TrimSpace(v)
				if v == "" {
					continue
				}
				i, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				distance = append(distance, i)
			}
		}
	}
	races := []Race{}
	for i, v := range time {
		races = append(races, Race{
			Time:           v,
			RecordDistance: distance[i],
		})
	}
	return races
}

func InputToRace2(input io.ReadCloser) Race {
	s := bufio.NewScanner(input)
	defer input.Close()
	time := 0
	distance := 0
	for s.Scan() {
		l := s.Text()
		if strings.Contains(l, "Time") {
			ts := strings.TrimLeft(l, "Time:")
			ts = strings.ReplaceAll(ts, " ", "")
			i, err := strconv.Atoi(ts)
			if err != nil {
				log.Fatal(err)
			}
			time = i
		} else if strings.Contains(l, "Distance") {
			ts := strings.TrimLeft(l, "Distance:")
			ts = strings.ReplaceAll(ts, " ", "")
			i, err := strconv.Atoi(ts)
			if err != nil {
				log.Fatal(err)
			}
			distance = i
		}
	}
	return Race{
		Time:           time,
		RecordDistance: distance,
	}
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	rs := InputToRace(f)

	total := 1
	for _, r := range rs {
		total *= r.CountPossibleRecords()
	}
	fmt.Printf("Total: %d\n", total)

	f, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	race := InputToRace2(f)
	fmt.Printf("Total: %d\n", race.CountPossibleRecords())
}
