package exercise1

import (
	"adventofcode/common/solution"
	. "adventofcode/common/utils"
	"errors"
	"strconv"
)

type Exercise1 struct {
	input []int
}

func (e *Exercise1) Prepare(isTest bool) error {
	input := ReadInputArray(1, isTest)
	numbers := make([]int, len(input))
	for i, line := range input {
		number, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		numbers[i] = number
	}

	e.input = numbers

	return nil
}

func (e *Exercise1) Solution1() (solution.Solution, error) {
	numbers := e.input
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j && numbers[i]+numbers[j] == 2020 {
				return solution.New(strconv.Itoa(numbers[i] * numbers[j])), nil
			}
		}
	}

	return nil, errors.New("no solution found")
}

func (e *Exercise1) Solution2() (solution.Solution, error) {
	numbers := e.input
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if i != j && j != k && k != i && numbers[i]+numbers[j]+numbers[k] == 2020 {
					return solution.New(strconv.Itoa(numbers[i] * numbers[j] * numbers[k])), nil
				}
			}
		}
	}

	return nil, errors.New("no solution found")
}
