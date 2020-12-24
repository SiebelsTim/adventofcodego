package exercise8

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	INST_ACC = iota
	INST_JMP
	INST_NOP
)

type Instruction interface {
	Kind() int
	Argument() int
	String() string
}

type Program []Instruction

type basicInstruction struct {
	kind     int
	argument int
}

func (b basicInstruction) Kind() int {
	return b.kind
}

func (b basicInstruction) Argument() int {
	return b.argument
}

func (b basicInstruction) String() string {
	name, err := instructionToName(b.kind)
	if err != nil {
		name = err.Error()
	}

	sign := '+'
	factor := 1
	if b.argument < 0 {
		sign = '-'
		factor = -1
	}

	return fmt.Sprintf("%s  %c%*d", name, sign, 3, factor * b.argument)
}

func (p Program) String() string {
	var ret strings.Builder

	for _, inst := range p {
		ret.WriteString(inst.String())
		ret.WriteString("\n")
	}

	return ret.String()
}

func NewNop(arg int) Instruction {
	return basicInstruction{kind: INST_NOP, argument: arg}
}
func NewAcc(arg int) Instruction {
	return basicInstruction{kind: INST_ACC, argument: arg}
}
func NewJmp(arg int) Instruction {
	return basicInstruction{kind: INST_JMP, argument: arg}
}

func Parse(line string) (Instruction, error) {
	parts := strings.Split(line, " ")

	arg, err := strconv.Atoi(parts[1])

	if err != nil {
		return nil, err
	}

	instruction, err := nameToInstruction(parts[0])
	if err != nil {
		return nil, err
	}

	return basicInstruction{
		kind:     instruction,
		argument: arg,
	}, nil
}

func nameToInstruction(name string) (int, error) {
	switch name {
	case "acc":
		return INST_ACC, nil
	case "jmp":
		return INST_JMP, nil
	case "nop":
		return INST_NOP, nil
	default:
		return 0, errors.New("unknown instruction " + name)
	}
}

func instructionToName(inst int) (string, error) {
	names := []string{
		"acc",
		"jmp",
		"nop",
	}

	if len(names) <= inst {
		return "", errors.New("could not find name for instruction " + strconv.Itoa(inst))
	}

	return names[inst], nil
}
