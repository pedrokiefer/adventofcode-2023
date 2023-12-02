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

type Game struct {
	ID   int64
	Sets []Set
}

type Set struct {
	Red   int64
	Green int64
	Blue  int64
}

func InputToGame(input io.ReadCloser) []Game {
	results := []Game{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		g, err := ReadGame(s.Text())
		if err != nil {
			continue
		}
		results = append(results, g)
	}
	return results
}

func ReadGame(s string) (Game, error) {
	g := Game{
		Sets: []Set{},
	}
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return Game{}, fmt.Errorf("invalid game string: %s", s)
	}

	gameIDStr := strings.TrimPrefix(parts[0], "Game ")
	gameID, err := strconv.ParseInt(gameIDStr, 10, 64)
	if err != nil {
		return Game{}, fmt.Errorf("invalid game id: %s", gameIDStr)
	}

	g.ID = gameID

	sets := strings.Split(parts[1], ";")
	for _, set := range sets {
		s := Set{}
		colors := strings.Split(set, ",")
		for _, color := range colors {
			color = strings.TrimSpace(color)
			if strings.HasSuffix(color, "red") {
				c, err := ReadColor(color, "red")
				if err != nil {
					return g, err
				}
				s.Red = c
			} else if strings.HasSuffix(color, "green") {
				c, err := ReadColor(color, "green")
				if err != nil {
					return g, err
				}
				s.Green = c
			} else if strings.HasSuffix(color, "blue") {
				c, err := ReadColor(color, "blue")
				if err != nil {
					return g, err
				}
				s.Blue = c
			}
		}
		g.Sets = append(g.Sets, s)
	}

	return g, nil
}

func ReadColor(s string, color string) (int64, error) {
	v := strings.TrimSuffix(s, color)
	v = strings.TrimSpace(v)
	val, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("invalid %s value: %s", color, v)
	}
	return val, nil
}

func (g Game) Valid(maxRed, maxGreen, maxBlue int64) bool {
	for _, set := range g.Sets {
		if set.Red > maxRed || set.Green > maxGreen || set.Blue > maxBlue {
			return false
		}
	}
	return true
}

func (g Game) MinimalSet() Set {
	min := Set{}
	for _, set := range g.Sets {
		if set.Red > min.Red {
			min.Red = set.Red
		}
		if set.Green > min.Green {
			min.Green = set.Green
		}
		if set.Blue > min.Blue {
			min.Blue = set.Blue
		}
	}
	return min
}

func (s Set) Power() int64 {
	return s.Red * s.Green * s.Blue
}

var MaxRed = int64(12)
var MaxGreen = int64(13)
var MaxBlue = int64(14)

func Sum(l []int64) int64 {
	var sum int64
	for _, v := range l {
		sum += v
	}
	return sum
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	games := InputToGame(f)

	gIds := []int64{}
	for _, g := range games {
		if g.Valid(MaxRed, MaxGreen, MaxBlue) {
			gIds = append(gIds, g.ID)
		}
	}
	fmt.Println(Sum(gIds))

	power := []int64{}
	for _, g := range games {
		power = append(power, g.MinimalSet().Power())
	}
	fmt.Println(Sum(power))
}
