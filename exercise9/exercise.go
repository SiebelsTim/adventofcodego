package exercise9

import (
	"adventofcode/common/solution"
	"adventofcode/common/utils"
	"errors"
	"math"
	"strconv"
)

type Exericse9 struct {
	preambleLength int
	isTest         bool
}

func (e *Exericse9) Input() <-chan int {
	ret := make(chan int)
	go func() {
		for line := range utils.ReadInput(9, e.isTest) {
			atoi, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ret <- atoi
		}
		close(ret)
	}()

	return ret
}

func (e *Exericse9) Prepare(isTest bool) error {
	if isTest {
		e.isTest = true
		e.preambleLength = 5
	} else {
		e.isTest = false
		e.preambleLength = 25
	}

	return nil
}

func createBuffer(input <-chan int, preambleLength int) RingBuffer {
	values := make([]int, preambleLength)

	for i := 0; i < preambleLength; i++ {
		values[i] = <-input
	}

	return NewBuffer(values)
}

func allows(buffer RingBuffer, number int) bool {
	values := buffer.Values()

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == number {
				return true
			}
		}
	}

	return false
}

func (e *Exericse9) findBadNumber() (int, error) {
	input := e.Input()
	defer func() {
		for range input {
		} // consume rest
	}()
	buffer := createBuffer(input, e.preambleLength)

	for number := range input {
		if allows(buffer, number) {
			buffer.Add(number)
		} else {
			return number, nil
		}
	}
	return 0, errors.New("no solution found")
}

func (e *Exericse9) Solution1() (solution.Solution, error) {
	number, err := e.findBadNumber()
	if err != nil {
		return nil, err
	}

	return solution.New(strconv.Itoa(number)), nil
}

func (e *Exericse9) Solution2() (solution.Solution, error) {
	number, err := e.findBadNumber()
	if err != nil {
		return nil, err
	}

	var inputs []int
	input := e.Input()
	defer func() {
		for range input {
		} // consume rest
	}()
	for n := range input {
		inputs = append(inputs, n)

		sum := 0
		var smallest int64 = math.MaxInt64
		var largest int64 = 0
		for _, i := range inputs {
			if smallest > int64(i) {
				smallest = int64(i)
			}
			if largest < int64(i) {
				largest = int64(i)
			}
			sum += i
			if sum == number {
				return solution.New(strconv.FormatInt(smallest+largest, 10)), nil
			} else if sum > number {
				inputs = inputs[1:] // pop first
				break
			}
		}
	}

	return solution.New("No solution found"), nil
}
