package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Interpreter struct {
	input      io.Reader
	output     chan interface{}
	executable bool
	commands   map[string]Instruction
}

func (i *Interpreter) RegisterCommand(name string, instr Instruction) {
	i.commands[name] = instr
}

func (i *Interpreter) Run() {
	sc := bufio.NewScanner(i.input)
	commands := make([]string, 0)
	for command := range i.commands {
		commands = append(commands, command)
	}
	sc.Split(i.instructionSplitFunc(commands))
	for sc.Scan() {
		parts := bytes.Split(sc.Bytes(), []byte("("))
		if len(parts) < 2 {
			continue
		}
		cmdName := string(parts[0])
		params := parts[1][:max(0, len(parts[1])-1)]
		instr := i.commands[cmdName]
		if err := instr.ParseParams(params); err != nil {
			i.output <- fmt.Errorf("parse parameters of %q error: %s", string(parts[0]), err)
			continue
		}
		if _, proceed := instr.(*DoInstr); !i.executable && !proceed {
			// fmt.Printf("Skip %q instruction\n", sc.Text())
			continue
		}
		res, err := instr.Execute(i)
		if err != nil {
			i.output <- fmt.Errorf("interpreter error: %s", err)
			continue
		}
		i.output <- res
		// fmt.Printf("Execute %q instruction\n", sc.Text())
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error scanning:", err)
	}

	close(i.output)
}

func (i *Interpreter) ChangeExecutable(turn bool) {
	// fmt.Printf("Executable = %t\n", turn)
	i.executable = turn
}

func (i *Interpreter) Output() <-chan interface{} {
	return i.output
}

func NewInterpreter(input io.Reader) *Interpreter {
	output := make(chan interface{})
	i := &Interpreter{input: input, output: output, executable: true, commands: make(map[string]Instruction)}
	return i
}

func (i *Interpreter) instructionSplitFunc(commands []string) bufio.SplitFunc {
	pattern := regexp.MustCompile(fmt.Sprintf(`(%s)\(([^()]*?)\)`, strings.Join(commands, "|")))
	return func(data []byte, atEOF bool) (avdance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if indeces := pattern.FindIndex(data); len(indeces) > 0 {
			return indeces[1], data[indeces[0]:indeces[1]], nil
		}

		if atEOF {
			return len(data), nil, nil
		}

		return
	}
}
