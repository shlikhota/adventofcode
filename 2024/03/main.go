package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.dat")
	noError(err)

	result := 0
	instructions := ParseInstructions(file)
	for _, instr := range instructions {
		result += instr.Execute().(int)
	}

	fmt.Println("Puzzle answer: ", result)
}

type Instruction interface {
	ParseParams(data []byte) error
	Execute() interface{}
}

type MulInstr struct {
	X, Y int
}

func (mi *MulInstr) Execute() interface{} {
	return mi.X * mi.Y
}

func (mi *MulInstr) ParseParams(params []byte) (err error) {
	parts := bytes.Split(params, []byte(","))
	if len(parts) != 2 {
		return fmt.Errorf("MulInstr must have only 2 params")
	}
	if mi.X, err = strconv.Atoi(string(parts[0])); err != nil {
		return err
	}
	if mi.Y, err = strconv.Atoi(string(parts[1])); err != nil {
		return err
	}
	return nil
}

func InstructionSplitFunc(command string) bufio.SplitFunc {
	pattern := regexp.MustCompile(fmt.Sprintf(`%s\([1-9][0-9]{0,2},[1-9][0-9]{0,2}\)`, command))
	return func(data []byte, atEOF bool) (avdance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if indeces := pattern.FindIndex(data); len(indeces) > 0 {
			idxStart, idxEnd := indeces[0], indeces[1]
			return idxEnd, data[idxStart+len(command)+1 : idxEnd-1], nil
		}

		if atEOF {
			return len(data), nil, nil
		}

		return
	}
}

func ParseInstructions(input io.Reader) []Instruction {
	sc := bufio.NewScanner(input)
	sc.Split(InstructionSplitFunc("mul"))
	result := make([]Instruction, 0)
	for sc.Scan() {
		instr := &MulInstr{}
		instr.ParseParams(sc.Bytes())
		result = append(result, instr)
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error scanning:", err)
	}

	return result
}

func noError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
