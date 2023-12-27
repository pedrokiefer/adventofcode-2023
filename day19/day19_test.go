package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`)))

	pf, parts := InputToPuzzleInput(input)

	assert.Equal(t, 11, len(pf.Workflows))
	assert.Equal(t, Workflow{
		Label: "px",
		Steps: []Step{
			&StepLessThan{
				Field:  "a",
				Value:  2006,
				Action: "qkq",
			},
			&StepGreaterThan{
				Field:  "m",
				Value:  2090,
				Action: "A",
			},
			&StepGoTo{
				Target: "rfg",
			},
		},
	}, pf.Workflows["px"])

	assert.Equal(t, Part{
		X: 787, M: 2655, A: 1222, S: 2876,
	}, parts[0])

	accepted := pf.FilterParts(parts)
	assert.Equal(t, []Part{
		{X: 787, M: 2655, A: 1222, S: 2876},
		{X: 2036, M: 264, A: 79, S: 2244},
		{X: 2127, M: 1623, A: 2188, S: 1013},
	}, accepted)

	assert.Equal(t, 19114, Sum(accepted))
}

func TestFilterAnalyzer(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(`px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`)))

	pf, _ := InputToPuzzleInput(input)

	assert.Equal(t, 11, len(pf.Workflows))

	pr := pf.Analyze()
	assert.Equal(t, []PartRange{
		{X: Range{Min: 1, Max: 1415}, M: Range{Min: 1, Max: 4000}, A: Range{Min: 1, Max: 2005}, S: Range{Min: 1, Max: 1350}},
		{X: Range{Min: 2663, Max: 4000}, M: Range{Min: 1, Max: 4000}, A: Range{Min: 1, Max: 2005}, S: Range{Min: 1, Max: 1350}},
		{X: Range{Min: 1, Max: 4000}, M: Range{Min: 2091, Max: 4000}, A: Range{Min: 2006, Max: 4000}, S: Range{Min: 1, Max: 1350}},
		{X: Range{Min: 1, Max: 2440}, M: Range{Min: 1, Max: 2090}, A: Range{Min: 2006, Max: 4000}, S: Range{Min: 537, Max: 1350}},
		{X: Range{Min: 1, Max: 4000}, M: Range{Min: 1, Max: 4000}, A: Range{Min: 1, Max: 4000}, S: Range{Min: 3449, Max: 4000}},
		{X: Range{Min: 1, Max: 4000}, M: Range{Min: 1549, Max: 4000}, A: Range{Min: 1, Max: 4000}, S: Range{Min: 2771, Max: 3448}},
		{X: Range{Min: 1, Max: 4000}, M: Range{Min: 1, Max: 1548}, A: Range{Min: 1, Max: 4000}, S: Range{Min: 2771, Max: 3448}},
		{X: Range{Min: 1, Max: 4000}, M: Range{Min: 839, Max: 1800}, A: Range{Min: 1, Max: 4000}, S: Range{Min: 1351, Max: 2770}},
		{X: Range{Min: 1, Max: 4000}, M: Range{Min: 1, Max: 838}, A: Range{Min: 1, Max: 1716}, S: Range{Min: 1351, Max: 2770}},
	}, pr)
	assert.Equal(t, 15320205000000, pr[0].Value())

	s := 0
	for _, pr := range pr {
		s += pr.Value()
	}
	assert.Equal(t, 167409079868000, s)
}
