package exercise3

import (
	"adventofcode/common/solution"
	"adventofcode/common/utils"
	"strconv"
	"strings"
)

type Exercise3 struct {
	input  [][]byte
	width  int
	height int
}

type Position struct {
	X int
	Y int
}

func (p Position) add(p2 Position) Position {
	return Position{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (e *Exercise3) getCell(p Position) byte {
	return e.input[p.Y][p.X%e.width]
}

func (e *Exercise3) String() string {
	ret := strings.Builder{}
	for _, line := range e.input {
		ret.Write(line)
		ret.WriteByte('\n')
	}

	return ret.String()
}

func (e *Exercise3) Prepare(isTest bool) error {
	var input [][]byte
	i := 0
	for line := range utils.ReadInput(3, isTest) {
		input = append(input, make([]byte, len(line)))
		input[i] = []byte(line)
		i++
	}

	e.input = input
	e.height = i
	e.width = len(input[0])

	return nil
}

func (e *Exercise3) walk(slope Position) int {
	position := Position{X: 0, Y: 0}
	treeCount := 0

	for position.Y < e.height {
		if e.getCell(position) == '#' {
			treeCount++
		}
		position = position.add(slope)
	}

	return treeCount
}

func (e *Exercise3) Solution1() (solution.Solution, error) {
	slope := Position{X: 3, Y: 1}

	return solution.New(strconv.Itoa(e.walk(slope))), nil
}

func (e *Exercise3) Solution2() (solution.Solution, error) {
	slopes := []Position{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	ch := make(chan int, len(slopes))
	ret := 1
	for _, slope := range slopes {
		slope := slope
		go func() {
			ch <- e.walk(slope)
		}()
	}

	for i := 0; i < len(slopes); i++ {
		ret *= <-ch
	}

	return solution.New(strconv.Itoa(ret)), nil
}
