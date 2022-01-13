package parser

import (
	"fmt"
	"os"

	"github.com/H1ghBre4k3r/go-bf/internal/tokens"
)

func Parse(lexed []int) []Instruction {
	parsed, index := parse(lexed, 0)
	// if we didn't parse everything before returning, there was a bracket error
	if index != len(lexed) {
		fmt.Println("Bracket error!")
		os.Exit(-1)
	}
	return parsed
}

func parse(lexed []int, index int) ([]Instruction, int) {
	instructions := make([]Instruction, 0)
	i := 0

loop:
	for index < len(lexed) {
		l := lexed[index]
		instructionCount := len(instructions)
		// switch through current
		switch l {
		case tokens.PLUS:
			if instructionCount == 0 || instructions[instructionCount-1].Operation != ADD {
				instructions = append(instructions, Instruction{
					Operation: ADD,
					Operand:   1,
				})
			} else {
				instructions[instructionCount-1].Operand += 1
			}

		case tokens.MINUS:
			if instructionCount == 0 || instructions[instructionCount-1].Operation != ADD {
				instructions = append(instructions, Instruction{
					Operation: ADD,
					Operand:   -1,
				})
			} else {
				instructions[instructionCount-1].Operand -= 1
			}

		case tokens.RIGHT:
			if instructionCount == 0 || instructions[instructionCount-1].Operation != MOVE {
				instructions = append(instructions, Instruction{
					Operation: MOVE,
					Operand:   1,
				})
			} else {
				instructions[instructionCount-1].Operand += 1
			}

		case tokens.LEFT:
			if instructionCount == 0 || instructions[instructionCount-1].Operation != MOVE {
				instructions = append(instructions, Instruction{
					Operation: MOVE,
					Operand:   -1,
				})
			} else {
				instructions[instructionCount-1].Operand -= 1
			}

		case tokens.OUT:
			instructions = append(instructions, Instruction{
				Operation: PRINT,
			})

		case tokens.IN:
			instructions = append(instructions, Instruction{
				Operation: READ,
			})

		case tokens.START_LOOP:
			// TODO lome: actually move code inside of parsed look token to properly represent control flow
			parsed, newIndex := parse(lexed, index+1)
			instructions = append(instructions, Instruction{
				Operation: START_LOOP,
				Operand:   newIndex - index,
			})
			instructions = append(instructions, parsed...)
			index = newIndex + 1
			continue

		case tokens.END_LOOP:
			// TODO lome: remove that in favor of bundled loops
			instructions = append(instructions, Instruction{
				Operation: END_LOOP,
			})
			break loop
		}

		index += 1
	}

	return instructions, index + i
}
