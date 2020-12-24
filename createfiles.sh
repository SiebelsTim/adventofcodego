#!/bin/bash

[ $# -eq 1 ] || exit 1

[ -d "exercise$1" ] || mkdir "exercise$1"

touch "exercise$1/input.txt"
touch "exercise$1/testinput.txt"


createMain() {
  cat <<EOD | sed "s/%N%/$1/g" > "exercise$1/exercise.go"
package exercise%N%

import (
	"adventofcode/common/solution"
	"adventofcode/common/utils"
)

type Exericse%N% struct {
}

func (e *Exericse%N%) Prepare(isTest bool) error {
	// TODO
	return nil
}

func (e *Exericse%N%) Solution1() (solution.Solution, error) {
	return solution.New("Not yet implemented"), nil
}

func (e *Exericse%N%) Solution2() (solution.Solution, error) {
	return solution.New("Not yet implemented"), nil
}

EOD
}

[ -f "exercise$1/exercise.go" ] || createMain $1
