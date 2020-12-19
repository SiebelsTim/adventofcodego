package exercise2

import (
	"adventofcode/common/solution"
	. "adventofcode/common/utils"
	"strconv"
	"strings"
)

type Policy struct {
	letter byte
	min int
	max int
}

type Password struct {
	Policy Policy
	Password string
}

type Exercise2 struct {
	input []Password
}

func (p Password) isValid() bool {
	password := p.Password

	count := 0

	for idx := range password {
		if password[idx] == p.Policy.letter {
			count++
		}
	}

	return count >= p.Policy.min && count <= p.Policy.max
}


func (p Password) isValid2() bool {
	password := p.Password
	firstPosition := p.Policy.min
	lastPosition := p.Policy.max

	return BoolToInt(password[firstPosition-1] == p.Policy.letter) + BoolToInt(password[lastPosition-1] == p.Policy.letter) == 1
}

func (e *Exercise2) Prepare() error {
	lines := ReadInput(2)

	rows := make(chan Password)
	linecount := 0

	for line := range lines {
		line := line
		go func() {
			parts := strings.Split(line, ": ")
			rule := parts[0]
			password := parts[1]

			parts = strings.Split(rule, " ")
			rng := parts[0]
			letter := parts[1][0]
			parts = strings.Split(rng, "-")
			min, _ := strconv.Atoi(parts[0])
			max, _ := strconv.Atoi(parts[1])

			policy := Policy{letter, min, max}

			rows <- Password{policy, password}
		}()
		linecount++
	}

	var ret []Password
	for i:=0; i < linecount; i++ {
		ret = append(ret, <-rows)
	}

	e.input = ret

	return nil
}

func (e *Exercise2) Solution1() (solution.Solution, error) {
	validCount := 0

	for _, password := range e.input {
		if password.isValid() {
			validCount++
		}
	}

	return solution.New(strconv.Itoa(validCount)), nil
}

func (e *Exercise2) Solution2() (solution.Solution, error) {
	validCount := 0

	for _, password := range e.input {
		if password.isValid2() {
			validCount++
		}
	}

	return solution.New(strconv.Itoa(validCount)), nil
}