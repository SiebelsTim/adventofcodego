package exercise8

import (
	"adventofcode/common/solution"
	"adventofcode/common/utils"
	"runtime"
	"strconv"
)

type Exericse8 struct {
	input Program
}

func (e *Exericse8) Prepare(isTest bool) error {
	input := utils.ReadInput(8, isTest)

	var ret []Instruction
	for line := range input {
		parse, err := Parse(line)
		if err != nil {
			return err
		}
		ret = append(ret, parse)
	}

	e.input = ret

	return nil
}

func (e *Exericse8) Solution1() (solution.Solution, error) {
	vm := stops(e.input)

	return solution.New(strconv.Itoa(vm.Accumulator())), nil
}

func generateModifiedPrograms(originalProgram Program, newPrograms chan<- Program) {
	modifiedEntry := 0
	var program Program

	for {
		program, modifiedEntry = swapJmp(originalProgram, modifiedEntry, INST_NOP, INST_JMP)
		if modifiedEntry == -1 {
			break
		}
		modifiedEntry++
		newPrograms <- program
	}

	modifiedEntry = 0

	for {
		program, modifiedEntry = swapJmp(originalProgram, modifiedEntry, INST_JMP, INST_NOP)
		if modifiedEntry == -1 {
			break
		}
		modifiedEntry++
		newPrograms <- program
	}

	close(newPrograms)
}

func worker(newPrograms <-chan Program, keepRunning *bool, result chan<- int) {
	for *keepRunning {
		program := <-newPrograms
		if vm := stops(program); vm.Stopped() {
			*keepRunning = false
			result <- vm.Accumulator()
			return
		}
	}
}

func (e *Exericse8) Solution2() (solution.Solution, error) {
	newPrograms := make(chan Program, 128)
	go generateModifiedPrograms(e.input, newPrograms) // this probably leaks, but whatever

	result := make(chan int)
	keepRunning := true // Don't need locking, as it only turns from true to false once

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(newPrograms, &keepRunning, result)
	}

	// We assume that we always get a result. Otherwise we wait here forever
	return solution.New(strconv.Itoa(<-result)), nil
}

// returns the VM after the program ran
// either infinitely (stopping after first loop) or until the end
// if vm.Stopped() returns true after running this function, the VM came to a regular halt
// if vm.Stopped() returns false after running this function, the VM loops after this point
func stops(program Program) VM {
	vm := NewVM(program)

	visitedInstructions := make([]bool, len(program))

	for !vm.Stopped() {
		if visitedInstructions[vm.Ip()] {
			break
		}
		visitedInstructions[vm.Ip()] = true
		vm.Step()
	}

	return vm
}

// swaps a single instruction from "take" to "replace" and returns the new program, and the index that changed
func swapJmp(program Program, entry int, take int, replace int) (Program, int) {
	if entry >= len(program) {
		return program, -1
	}

	newProg := append([]Instruction{}, program...) // copy
	var modifiedEntry = -1

	for idx, instruction := range newProg[entry:] {
		if instruction.Kind() == take {
			modifiedEntry = idx + entry
			newProg[modifiedEntry] = basicInstruction{kind: replace, argument: instruction.Argument()}
			break
		}
	}

	return newProg, modifiedEntry
}
