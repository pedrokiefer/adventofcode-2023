package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`)))
	is := InputToInitializationSequence(input)

	assert.Equal(t, []string{"rn=1", "cm-", "qp=3", "cm=2", "qp-", "pc=4", "ot=9", "ab=5", "pc-", "pc=6", "ot=7"}, is)

	h := InitializationSequenceHASH(is)
	assert.Equal(t, 1320, h)
}

func TestHolidayASCIIStringHelperAlgorithm(t *testing.T) {
	x := HolidayASCIIStringHelperAlgorithm("rn=1")
	assert.Equal(t, 30, x)
}

func TestManualArrangementProcedure(t *testing.T) {
	op := InputToOperation("rn=1")
	assert.Equal(t, Operation{
		Label:       "rn",
		Box:         0,
		Type:        AddLens,
		FocalLength: 1,
	}, op)

	op2 := InputToOperation("cm-")
	assert.Equal(t, Operation{
		Label:       "cm",
		Box:         0,
		Type:        RemoveLens,
		FocalLength: 0,
	}, op2)

	b := ManualArrangementProcedure([]string{"rn=1", "cm-", "qp=3", "cm=2", "qp-", "pc=4", "ot=9", "ab=5", "pc-", "pc=6", "ot=7"})
	assert.Equal(t, Box{
		Lenses: []*Lens{
			{Label: "rn", Box: 0, FocalLength: 1},
			{Label: "cm", Box: 0, FocalLength: 2},
		},
	}, b[0])
	assert.Equal(t, Box{
		Lenses: []*Lens{
			{Label: "ot", Box: 3, FocalLength: 7},
			{Label: "ab", Box: 3, FocalLength: 5},
			{Label: "pc", Box: 3, FocalLength: 6},
		},
	}, b[3])

	assert.Equal(t, 5, b[0].FocusingPower())
	assert.Equal(t, 0, b[1].FocusingPower())
	assert.Equal(t, 0, b[2].FocusingPower())
	assert.Equal(t, 140, b[3].FocusingPower())
}
