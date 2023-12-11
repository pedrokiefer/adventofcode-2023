package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`)))

	cards := InputToCards(input)

	assert.Equal(t, []*Card{
		{
			CardID:         1,
			WinningNumbers: []int64{41, 48, 83, 86, 17},
			MyNumbers:      []int64{83, 86, 6, 31, 17, 9, 48, 53},
			Points:         8,
			MatchingCount:  4,
		},
		{
			CardID:         2,
			WinningNumbers: []int64{13, 32, 20, 16, 61},
			MyNumbers:      []int64{61, 30, 68, 82, 17, 32, 24, 19},
			Points:         2,
			MatchingCount:  2,
		},
		{
			CardID:         3,
			WinningNumbers: []int64{1, 21, 53, 59, 44},
			MyNumbers:      []int64{69, 82, 63, 72, 16, 21, 14, 1},
			Points:         2,
			MatchingCount:  2,
		},
		{
			CardID:         4,
			WinningNumbers: []int64{41, 92, 73, 84, 69},
			MyNumbers:      []int64{59, 84, 76, 51, 58, 5, 54, 83},
			Points:         1,
			MatchingCount:  1,
		},
		{
			CardID:         5,
			WinningNumbers: []int64{87, 83, 26, 28, 32},
			MyNumbers:      []int64{88, 30, 70, 12, 93, 22, 82, 36},
			Points:         0,
			MatchingCount:  0,
		},
		{
			CardID:         6,
			WinningNumbers: []int64{31, 18, 13, 56, 72},
			MyNumbers:      []int64{74, 77, 10, 23, 35, 67, 36, 11},
			Points:         0,
			MatchingCount:  0,
		},
	}, cards)

	sum := int64(0)
	for _, c := range cards {
		sum += c.Points
	}
	assert.Equal(t, int64(13), sum)

	matching := CountMatching(cards)
	assert.Equal(t, []int64{1, 2, 4, 8, 14, 1}, matching)
}
