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

type Card struct {
	CardID         int64
	WinningNumbers []int64
	MyNumbers      []int64
	Points         int64
	MatchingCount  int64
}

func (c Card) String() string {
	return fmt.Sprintf("Card %d: %v | %v", c.CardID, c.WinningNumbers, c.MyNumbers)
}

func InputToCards(input io.ReadCloser) []*Card {
	results := []*Card{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		c, err := ReadCard(s.Text())
		if err != nil {
			continue
		}
		results = append(results, c)
	}
	return results
}

func ReadCard(s string) (*Card, error) {
	c := &Card{}
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid card string: %s", s)
	}

	cardIDStr := strings.TrimPrefix(parts[0], "Card ")
	cardIDStr = strings.TrimSpace(cardIDStr)
	cardID, err := strconv.ParseInt(cardIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid card id: %s", cardIDStr)
	}
	c.CardID = cardID

	numbers := strings.Split(parts[1], "|")
	if len(numbers) != 2 {
		return nil, fmt.Errorf("invalid card numbers: %s", parts[1])
	}

	c.WinningNumbers = NumbersToSlice(numbers[0])
	c.MyNumbers = NumbersToSlice(numbers[1])

	c.Points = CheckPoints(c.WinningNumbers, c.MyNumbers)
	c.MatchingCount = CheckMatching(c.WinningNumbers, c.MyNumbers)

	return c, nil
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

func CheckPoints(winningNumbers []int64, myNumbers []int64) int64 {
	points := int64(0)
	for _, wn := range winningNumbers {
		for _, n := range myNumbers {
			if wn == n {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
	}
	return points
}

func CheckMatching(winningNumbers []int64, myNumbers []int64) int64 {
	points := int64(0)
	for _, wn := range winningNumbers {
		for _, n := range myNumbers {
			if wn == n {
				points++
			}
		}
	}
	return points
}

func CountMatching(cards []*Card) []int64 {
	result := Ones(len(cards))
	for x, c := range cards {
		if c.MatchingCount == 0 {
			continue
		}
		for c2 := int64(x + 1); c2 < c.MatchingCount+int64(x+1); c2++ {
			if c2 >= int64(len(cards)) {
				break
			}
			result[c2] = result[c2] + result[x]
		}
	}
	return result
}

func PrintSlice(vv []int64) string {
	s := "[ "
	for _, v := range vv {
		s = fmt.Sprintf("%s%d ", s, v)
	}
	s = fmt.Sprintf("%s]", s)
	return s
}

func Ones(lenght int) []int64 {
	results := make([]int64, lenght)
	for i := 0; i < lenght; i++ {
		results[i] = 1
	}
	return results
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	cards := InputToCards(f)

	sum := int64(0)
	for _, c := range cards {
		sum += c.Points
	}
	fmt.Printf("Total points: %d\n", sum)

	matching := CountMatching(cards)
	sum2 := int64(0)
	for _, v := range matching {
		sum2 += v
	}
	fmt.Printf("Total matching: %d\n", sum2)
}
