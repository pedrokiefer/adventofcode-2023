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

type Lens struct {
	Label       string
	Box         int
	FocalLength int
}

type OperationType string

var (
	AddLens    = OperationType("=")
	RemoveLens = OperationType("-")
)

type Operation struct {
	Label       string
	Box         int
	Type        OperationType
	FocalLength int
}

type Box struct {
	Lenses []*Lens
}

func InputToInitializationSequence(input io.ReadCloser) []string {
	result := []string{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		l := s.Text()
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		p := strings.Split(l, ",")
		result = append(result, p...)
	}
	return result
}

func HolidayASCIIStringHelperAlgorithm(i string) int {
	hash := 0
	for _, c := range i {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func InitializationSequenceHASH(is []string) int {
	sum := 0
	for _, i := range is {
		sum += HolidayASCIIStringHelperAlgorithm(i)
	}
	return sum
}

func InputToOperation(i string) Operation {
	if strings.Contains(i, "=") {
		parts := strings.Split(i, "=")
		f, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		return Operation{
			Label:       parts[0],
			Box:         HolidayASCIIStringHelperAlgorithm(parts[0]),
			Type:        AddLens,
			FocalLength: f,
		}
	} else if strings.Contains(i, "-") {
		l := strings.TrimSuffix(i, "-")
		return Operation{
			Label: l,
			Box:   HolidayASCIIStringHelperAlgorithm(l),
			Type:  RemoveLens,
		}
	}
	return Operation{}
}

func (b *Box) GetLens(label string) (*Lens, bool) {
	for _, l := range b.Lenses {
		if l.Label == label {
			return l, true
		}
	}
	return nil, false
}

func ManualArrangementProcedure(is []string) []Box {
	boxes := make([]Box, 256)
	for _, i := range is {
		op := InputToOperation(i)
		if op.Type == AddLens {
			if l, ok := boxes[op.Box].GetLens(op.Label); ok {
				l.FocalLength = op.FocalLength
			} else {
				boxes[op.Box].Lenses = append(boxes[op.Box].Lenses, &Lens{
					Label:       op.Label,
					Box:         op.Box,
					FocalLength: op.FocalLength,
				})
			}
		} else if op.Type == RemoveLens {
			for j, l := range boxes[op.Box].Lenses {
				if l.Label == op.Label {
					boxes[op.Box].Lenses = append(boxes[op.Box].Lenses[:j], boxes[op.Box].Lenses[j+1:]...)
					break
				}
			}
		}
	}
	return boxes
}

func (b *Box) FocusingPower() int {
	sum := 0
	for c, l := range b.Lenses {
		p := (l.Box + 1) * (c + 1) * l.FocalLength
		sum += p
	}
	return sum
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	is := InputToInitializationSequence(f)
	h := InitializationSequenceHASH(is)
	fmt.Println(h)

	boxes := ManualArrangementProcedure(is)
	total := 0
	for _, b := range boxes {
		total += b.FocusingPower()
	}
	fmt.Println(total)
}
