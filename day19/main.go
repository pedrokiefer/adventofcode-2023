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
	InitialWorkflow = "in"
)

type Part struct {
	X int
	M int
	A int
	S int
}

func (p *Part) Value() int {
	return p.X + p.M + p.A + p.S
}

type Range struct {
	Min int
	Max int
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (r Range) Split(value int, op string) (Range, Range) {
	if value > r.Max {
		return r, Range{}
	}
	var min, max int
	if op == "<" {
		min, max = value-1, value
	} else {
		min, max = value, value+1
	}
	r1 := Range{Min: Max(1, r.Min), Max: Min(min, r.Max)}
	r2 := Range{Min: Max(max, r.Min), Max: Min(4000, r.Max)}
	return r1, r2
}

func (r Range) Value() int {
	return r.Max - r.Min + 1
}

func (r Range) String() string {
	return fmt.Sprintf("[%d,%d]", r.Min, r.Max)
}

type PartRange struct {
	X Range
	M Range
	A Range
	S Range
}

func (pr *PartRange) Clone(field string, r Range) PartRange {
	p := PartRange{
		X: pr.X,
		M: pr.M,
		A: pr.A,
		S: pr.S,
	}
	switch field {
	case "x":
		p.X = r
	case "m":
		p.M = r
	case "a":
		p.A = r
	case "s":
		p.S = r
	}
	return p
}

func (pr PartRange) Value() int {
	//fmt.Printf("x=%d, m=%d, a=%d, s=%d ==> ", pr.X.Value(), pr.M.Value(), pr.A.Value(), pr.S.Value())
	v := pr.X.Value() * pr.M.Value() * pr.A.Value() * pr.S.Value()
	//fmt.Printf("%d\n", v)
	return v
}

func (pr PartRange) String() string {
	return fmt.Sprintf("{x=%s, m=%s, a=%s, s=%s}", pr.X.String(), pr.M.String(), pr.A.String(), pr.S.String())
}

type Step interface {
	Check(p *Part) (bool, string)
	Analyze(r PartRange) ([]PartRange, string)
	Label() string
}

type StepLessThan struct {
	Field  string
	Value  int
	Action string
}

func (s *StepLessThan) Check(p *Part) (bool, string) {
	pv := 0
	switch s.Field {
	case "x":
		pv = p.X
	case "m":
		pv = p.M
	case "a":
		pv = p.A
	case "s":
		pv = p.S
	}
	if pv < s.Value {
		return true, s.Action
	}
	return false, ""
}

func (s *StepLessThan) Analyze(r PartRange) ([]PartRange, string) {
	switch s.Field {
	case "x":
		r1, r2 := r.X.Split(s.Value, "<")
		return []PartRange{
			r.Clone("x", r1),
			r.Clone("x", r2),
		}, s.Action
	case "m":
		r1, r2 := r.M.Split(s.Value, "<")
		return []PartRange{
			r.Clone("m", r1),
			r.Clone("m", r2),
		}, s.Action
	case "a":
		r1, r2 := r.A.Split(s.Value, "<")
		return []PartRange{
			r.Clone("a", r1),
			r.Clone("a", r2),
		}, s.Action
	case "s":
		r1, r2 := r.S.Split(s.Value, "<")
		return []PartRange{
			r.Clone("s", r1),
			r.Clone("s", r2),
		}, s.Action
	}
	return nil, ""
}

func (s *StepLessThan) Label() string {
	return fmt.Sprintf("%s<%d:%s", s.Field, s.Value, s.Action)
}

type StepGreaterThan struct {
	Field  string
	Value  int
	Action string
}

func (s *StepGreaterThan) Check(p *Part) (bool, string) {
	pv := 0
	switch s.Field {
	case "x":
		pv = p.X
	case "m":
		pv = p.M
	case "a":
		pv = p.A
	case "s":
		pv = p.S
	}
	if pv > s.Value {
		return true, s.Action
	}
	return false, ""
}

func (s *StepGreaterThan) Analyze(r PartRange) ([]PartRange, string) {
	var p1, p2 PartRange
	switch s.Field {
	case "x":
		r1, r2 := r.X.Split(s.Value, ">")
		p1 = r.Clone("x", r1)
		p2 = r.Clone("x", r2)
	case "m":
		r1, r2 := r.M.Split(s.Value, ">")
		p1 = r.Clone("m", r1)
		p2 = r.Clone("m", r2)
	case "a":
		r1, r2 := r.A.Split(s.Value, ">")
		p1 = r.Clone("a", r1)
		p2 = r.Clone("a", r2)
	case "s":
		r1, r2 := r.S.Split(s.Value, ">")
		p1 = r.Clone("s", r1)
		p2 = r.Clone("s", r2)
	}
	return []PartRange{
		p2,
		p1,
	}, s.Action
}

func (s *StepGreaterThan) Label() string {
	return fmt.Sprintf("%s>%d:%s", s.Field, s.Value, s.Action)
}

type StepGoTo struct {
	Target string
}

func (s *StepGoTo) Check(p *Part) (bool, string) {
	return true, s.Target
}

func (s *StepGoTo) Analyze(r PartRange) ([]PartRange, string) {
	return []PartRange{r, r}, s.Target
}

func (s *StepGoTo) Label() string {
	return s.Target
}

type Workflow struct {
	Label string
	Steps []Step
}

type PartFilter struct {
	Workflows map[string]Workflow
}

func (pf *PartFilter) analyzeRecursive(w Workflow, pr PartRange, level int) []PartRange {
	validPrs := []PartRange{}
	for _, s := range w.Steps {
		// fmt.Printf("%s- analyzing: %s step %s range %s\n", strings.Repeat(" ", level), w.Label, s.Label(), pr.String())
		rr, target := s.Analyze(pr)
		if target == "R" {
			pr = rr[1]
			continue
		}
		if target == "A" {
			// fmt.Printf("%s* accepting: %s\n", strings.Repeat(" ", level), rr[0].String())
			validPrs = append(validPrs, rr[0])
			pr = rr[1]
			continue
		}
		if _, has := pf.Workflows[target]; !has {
			log.Printf("missing workflow: %s", target)
			continue
		}
		validPrs = append(validPrs, pf.analyzeRecursive(pf.Workflows[target], rr[0], level+1)...)
		pr = rr[1]
	}
	return validPrs
}

func (pf *PartFilter) Analyze() []PartRange {
	w := pf.Workflows[InitialWorkflow]
	pr := PartRange{
		X: Range{Min: 1, Max: 4000},
		M: Range{Min: 1, Max: 4000},
		A: Range{Min: 1, Max: 4000},
		S: Range{Min: 1, Max: 4000},
	}

	return pf.analyzeRecursive(w, pr, 0)
}

func (pf *PartFilter) Filter(part *Part) bool {
	w := pf.Workflows[InitialWorkflow]

begin:
	for _, s := range w.Steps {
		ok, target := s.Check(part)
		if !ok {
			continue
		}
		if target == "" {
			log.Printf("missing target!")
			return false
		}
		if target == "R" {
			return false
		}
		if target == "A" {
			return true
		}
		if _, has := pf.Workflows[target]; !has {
			log.Printf("missing workflow: %s", target)
			return false
		}
		w = pf.Workflows[target]
		goto begin
	}
	return false
}

func (pf *PartFilter) FilterParts(parts []Part) []Part {
	accepted := []Part{}
	for _, p := range parts {
		v := pf.Filter(&p)
		if v {
			accepted = append(accepted, p)
		}
	}
	return accepted
}

func ToSteps(s string) []Step {
	steps := []Step{}
	for _, v := range strings.Split(s, ",") {
		if strings.Contains(v, "<") {
			vStr := strings.Split(v, "<")
			if len(vStr) != 2 {
				log.Printf("invalid step: %s", v)
				continue
			}
			targetStr := strings.Split(vStr[1], ":")
			if len(targetStr) != 2 {
				log.Printf("invalid step: %s", v)
				continue
			}
			value, err := strconv.Atoi(targetStr[0])
			if err != nil {
				log.Printf("invalid step: %s", v)
				continue
			}
			steps = append(steps, &StepLessThan{
				Value:  value,
				Field:  vStr[0],
				Action: targetStr[1],
			})
		} else if strings.Contains(v, ">") {
			vStr := strings.Split(v, ">")
			if len(vStr) != 2 {
				log.Printf("invalid step: %s", v)
				continue
			}
			targetStr := strings.Split(vStr[1], ":")
			if len(targetStr) != 2 {
				log.Printf("invalid step: %s", v)
				continue
			}
			value, err := strconv.Atoi(targetStr[0])
			if err != nil {
				log.Printf("invalid step: %s", v)
				continue
			}
			steps = append(steps, &StepGreaterThan{
				Value:  value,
				Field:  vStr[0],
				Action: targetStr[1],
			})
		} else {
			steps = append(steps, &StepGoTo{
				Target: v,
			})
		}
	}
	return steps
}

func InputToPuzzleInput(input io.ReadCloser) (*PartFilter, []Part) {
	s := bufio.NewScanner(input)
	defer input.Close()
	p := "workflows"
	pf := &PartFilter{
		Workflows: map[string]Workflow{},
	}
	parts := []Part{}
	for s.Scan() {
		l := s.Text()
		if l == "" {
			p = "parts"
			continue
		}

		if p == "workflows" {
			label := strings.Split(l, "{")
			if len(label) != 2 {
				log.Printf("invalid workflow: %s", l)
				continue
			}
			steps := ToSteps(strings.TrimSuffix(label[1], "}"))
			wf := Workflow{
				Label: label[0],
				Steps: steps,
			}
			pf.Workflows[wf.Label] = wf
			continue
		} else if p == "parts" {
			pp := strings.TrimPrefix(l, "{")
			pp = strings.TrimSuffix(pp, "}")
			part := Part{}
			for _, vv := range strings.Split(pp, ",") {
				kv := strings.Split(vv, "=")
				value, err := strconv.Atoi(kv[1])
				if err != nil {
					log.Printf("Error parsing: %v", vv)
					continue
				}
				switch kv[0] {
				case "x":
					part.X = value
				case "m":
					part.M = value
				case "a":
					part.A = value
				case "s":
					part.S = value
				}
			}
			parts = append(parts, part)
		}

	}
	return pf, parts
}

func Sum(parts []Part) int {
	s := 0
	for _, p := range parts {
		s += p.Value()
	}
	return s
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	pf, parts := InputToPuzzleInput(f)
	accepted := pf.FilterParts(parts)
	v := Sum(accepted)
	fmt.Printf("Total parts: %d\n", v)

	pr := pf.Analyze()
	s := 0
	for _, pr := range pr {
		s += pr.Value()
	}
	fmt.Printf("Total combinations: %d\n", s)
}
