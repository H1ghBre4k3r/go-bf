package interpreter

import (
	"fmt"

	"github.com/H1ghBre4k3r/go-bf/internal/input"
	"github.com/H1ghBre4k3r/go-bf/internal/lexer"
	"github.com/H1ghBre4k3r/go-bf/internal/parser"
)

type Interpreter struct {
	program string
}

func New(inputPath string) *Interpreter {
	return &Interpreter{
		program: input.ReadFile(inputPath),
	}
}

func (i *Interpreter) Start() {
	lexed := lexer.Lex(i.program)
	parsed := parser.Parse(lexed)
	interpret(parsed)
}

func interpret(parsed []parser.Instruction) {
	memory := make([]byte, 300000)
	pointer := 0
	index := 0
	eval(parsed, &index, &memory, &pointer)
}

func eval(parsed []parser.Instruction, index *int, memory *[]byte, pointer *int) {
	for *index < len(parsed) {
		// get current symbol
		i := parsed[*index]
		// increment...cause, why not?
		*index++

		switch i.Operation {
		case parser.MOVE:
			*pointer += i.Operand

		case parser.ADD:
			(*memory)[*pointer] += byte(i.Operand)

		case parser.START_LOOP:
			// by default, skip the loop
			newIndex := *index + 2
			// looping lui
			for (*memory)[*pointer] != 0 {
				newIndex = *index
				eval(parsed, &newIndex, memory, pointer)
			}
			// start from the new index
			*index = newIndex

		case parser.END_LOOP:
			// here, we just return. The important variables are already changed via pointers.
			return

		case parser.PRINT:
			fmt.Print(string((*memory)[*pointer]))

		case parser.READ:
			// maybe implement that later
		}
	}
}
