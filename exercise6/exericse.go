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
	input := utils.ReadInput(6, isTest)

	currentGroup := group{}
	for line := range input {
		if line == "" {
			e.input = append(e.input, currentGroup)
			currentGroup = group{}
		} else {
			for _, c := range line {
				currentGroup.answers = append(currentGroup.answers, byte(c))
			}
			currentGroup.participantCount++
		}
	}
	e.input = append(e.input, currentGroup)

	return nil
}

func (g group) toMap() map[byte]int {
	ret := make(map[byte]int)

	for _, c := range g.answers {
		ret[c] += 1
	}

	return ret
}

func (e *Exericse6) Solution1() (solution.Solution, error) {
	ret := 0

	ch := make(chan int, 32)
	for _, g := range e.input {
		go func(g group) {
			ch <- len(g.toMap())
		}(g)
	}

	for i := 0; i < len(e.input); i++ {
		ret += <-ch
	}

	return solution.New(strconv.Itoa(ret)), nil
}

func (e *Exericse6) Solution2() (solution.Solution, error) {
	ret := 0

	ch := make(chan int, 32)
	for _, g := range e.input {
		go func(g group) {
			result := 0
			expectedCount := g.participantCount
			answers := g.toMap()
			for _, count := range answers {
				if count == expectedCount {
					result++
				}
			}
			ch <- result
		}(g)
	}

	for i := 0; i < len(e.input); i++ {
		ret += <-ch
	}

	return solution.New(strconv.Itoa(ret)), nil
}
