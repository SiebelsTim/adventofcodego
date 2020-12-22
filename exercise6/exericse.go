package exercise6

import (
	"adventofcode/common/solution"
	"adventofcode/common/utils"
	"strconv"
)

type group struct {
	answers          []byte
	participantCount int
}

type Exericse6 struct {
	input []group
}

func (e *Exericse6) Prepare(isTest bool) error {
	e.input = []group{}
	input := utils.ReadInput(6, isTest)

	currentGroup := group{}
	for line := range input {
		if line == "" {
			e.input = append(e.input, currentGroup)
			currentGroup = group{}
		} else {
			currentGroup.answers = append(currentGroup.answers, []byte(line)...)
			currentGroup.participantCount++
		}
	}
	e.input = append(e.input, currentGroup)

	return nil
}

func (g group) toMap() [26]int {
	ret := [26]int{}

	for _, c := range g.answers {
		ret[c - 'a'] += 1
	}

	return ret
}

func (e *Exericse6) Solution1() (solution.Solution, error) {
	ret := 0

	for _, g := range e.input {
		for _, count := range g.toMap() {
			if count > 0 {
				ret++
			}
		}
	}

	return solution.New(strconv.Itoa(ret)), nil
}

func (e *Exericse6) Solution2() (solution.Solution, error) {
	ret := 0

	for _, g := range e.input {
		expectedCount := g.participantCount
		answers := g.toMap()
		for _, count := range answers {
			if count == expectedCount {
				ret++
			}
		}
	}

	return solution.New(strconv.Itoa(ret)), nil
}
