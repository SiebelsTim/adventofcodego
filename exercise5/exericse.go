package exercise5

import (
	"adventofcode/common/solution"
	"adventofcode/common/utils"
	"errors"
	"strconv"
	"strings"
)

type Exericse5 struct {
	input []Pass
}

type Pass struct {
	partitioningRow    [7]bool
	partitioningColumn [3]bool
}

func parsePass(line string) Pass {
	row := [7]bool{
		line[0] == 'B', line[1] == 'B', line[2] == 'B', line[3] == 'B', line[4] == 'B', line[5] == 'B', line[6] == 'B',
	}
	column := [3]bool{
		line[7] == 'R', line[8] == 'R', line[9] == 'R',
	}
	return Pass{partitioningRow: row, partitioningColumn: column}
}

func partition(min int, max int, directions []bool) int {
	direction := directions[0]
	if len(directions) == 1 {
		if direction {
			return max
		} else {
			return min
		}
	}

	if direction {
		min = (max-min)/2 + min + 1
	} else {
		max = (max-min)/2 + min
	}

	return partition(min, max, directions[1:])
}

func (p Pass) coord() (int, int) {
	return partition(0, 127, p.partitioningRow[:]), partition(0, 7, p.partitioningColumn[:])
}

func (p Pass) seatId() int {
	row, column := p.coord()

	return row * 8 + column
}

func (p Pass) String() string {
	ret := strings.Builder{}
	ret.WriteByte('[')
	for _, row := range p.partitioningRow {
		if row {
			ret.WriteByte('B')
		} else {
			ret.WriteByte('F')
		}
	}
	ret.WriteByte(']')
	ret.WriteByte('[')
	for _, row := range p.partitioningColumn {
		if row {
			ret.WriteByte('R')
		} else {
			ret.WriteByte('L')
		}
	}
	ret.WriteByte(']')

	return ret.String()
}

func (e *Exericse5) Prepare() error {
	input := utils.ReadInput(5)

	for line := range input {
		e.input = append(e.input, parsePass(line))
	}

	return nil
}

func (e *Exericse5) Solution1() (solution.Solution, error) {
	ret := 0

	for _, pass := range e.input {
		id := pass.seatId()
		if id > ret {
			ret = id
		}
	}

	return solution.New(strconv.Itoa(ret)), nil
}

func (e *Exericse5) Solution2() (solution.Solution, error) {
	allSeatIds := make(map [int]bool)

	for _, pass := range e.input {
		allSeatIds[pass.seatId()] = true
	}

	for seatId := 0; seatId < 901; seatId++ {
		if _, found := allSeatIds[seatId]; !found {
			_, prev := allSeatIds[seatId - 1]
			_, next := allSeatIds[seatId + 1]
			if prev && next {
				return solution.New(strconv.Itoa(seatId)), nil
			}
		}
	}

	return nil, errors.New("no solution found")
}
