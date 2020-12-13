package exercise2

import (
	. "exercise1/utils"
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

func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}


func (p Password) isValid2() bool {
	password := p.Password
	firstPosition := p.Policy.min
	lastPosition := p.Policy.max

	return boolToInt(password[firstPosition-1] == p.Policy.letter) + boolToInt(password[lastPosition-1] == p.Policy.letter) == 1
}

func (e *Exercise2) Prepare() error {
	lines := ReadInput(2)

	var ret []Password

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rule := parts[0]
		password := parts[1]

		parts = strings.Split(rule, " ")
		rng := parts[0]
		letter := parts[1][0]
		parts = strings.Split(rng, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		policy := Policy{ letter, min, max }

		ret = append(ret, Password{ policy, password})
	}

	e.input = ret

	return nil
}

func (e *Exercise2) Solution1() (Solution, error) {
	validCount := 0

	for _, password := range e.input {
		if password.isValid() {
			validCount++
		}
	}

	return New(strconv.Itoa(validCount)), nil
}

func (e *Exercise2) Solution2() (Solution, error) {
	validCount := 0

	for _, password := range e.input {
		if password.isValid2() {
			validCount++
		}
	}

	return New(strconv.Itoa(validCount)), nil
}