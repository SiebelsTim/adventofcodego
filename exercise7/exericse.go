package exercise7

import (
	"adventofcode/common/solution"
	"adventofcode/common/utils"
	"strconv"
)

type Exericse7 struct {
	input map[string]rule
}

func (e *Exericse7) Prepare(isTest bool) error {
	input := utils.ReadInput(7, isTest)

	var rules = make(map[string]rule)

	for line := range input {
		rule, err := parseRule(line)
		if err != nil {
			return err
		}
		rules[rule.outerBag.String()] = rule
	}

	e.input = rules

	return nil
}

func createContainedInMap(input map[string]rule) map[string][]bag {
	ret := make(map[string][]bag, len(input))
	for _, rule := range input {
		for _, contained := range rule.containedBags {
			ret[contained.String()] = append(ret[contained.String()], rule.outerBag)
		}
	}

	return ret
}

func (e *Exericse7) Solution1() (solution.Solution, error) {
	inMap := createContainedInMap(e.input)

	result := make(map[string]bool)
	queue := inMap["shiny gold"]

	for len(queue) > 0 {
		top := queue[0]
		if _, ok := result[top.String()]; ok { // already visited
			queue = queue[1:] // pop
			continue
		}
		result[top.String()] = true
		queue = append(queue, inMap[top.String()]...)

		queue = queue[1:] // pop
	}

	return solution.New(strconv.Itoa(len(result))), nil
}

func (e *Exericse7) Solution2() (solution.Solution, error) {
	queue := e.input["shiny gold"].containedBags
	ret := 0
	for len(queue) > 0 {
		top := queue[0]
		ret += top.count
		bags := e.input[top.String()].containedBags
		for i:=0; i < top.count; i++ {
			queue = append(queue, bags...)
		}

		queue = queue[1:] // pop
	}

	return solution.New(strconv.Itoa(ret)), nil
}
