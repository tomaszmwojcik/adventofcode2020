package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Instruction represents instruction in the game console
type Instruction struct {
	operation string
	argument  int
}

func evaluate(instructions []Instruction) (int, bool) {
	visited := make([]bool, len(instructions))
	instruction := instructions[0]
	i := 0
	acc := 0
	loop := false
	for i < len(instructions) {
		instruction = instructions[i]
		if visited[i] {
			loop = true
			break
		} else {
			visited[i] = true
		}
		if instruction.operation == "nop" {
			i++
		} else if instruction.operation == "acc" {
			acc += instruction.argument
			i++
		} else if instruction.operation == "jmp" {
			i += instruction.argument
		}
	}
	return acc, loop
}

func swapInstructions(instructions []Instruction, i int) bool {
	operation := instructions[i].operation
	if operation == "jmp" {
		instructions[i].operation = "nop"
		return true
	} else if operation == "nop" {
		instructions[i].operation = "jmp"
		return true
	}
	return false
}

func main() {
	instructions := make([]Instruction, 0, 10)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, " \n")
		lineSplit := strings.Split(line, " ")
		operation := lineSplit[0]
		argument, _ := strconv.Atoi(lineSplit[1])
		instructions = append(instructions, Instruction{operation, argument})
		if err != nil {
			break
		}
	}
	accumulatorVal, _ := evaluate(instructions)
	fmt.Printf("Accumulator value is %d just before the loop\n", accumulatorVal)
	fixedAccumulatorVal := 0
	for i := 0; i < len(instructions); i++ {
		swapped := swapInstructions(instructions, i)
		if swapped {
			var loop bool
			fixedAccumulatorVal, loop = evaluate(instructions)
			if !loop {
				break
			}
		}
		swapInstructions(instructions, i)
	}
	fmt.Printf("Accumulator value after the fix is %d\n", fixedAccumulatorVal)
}
