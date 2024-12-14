package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Instruction interface {
	ParseParams([]byte) error
	Execute(*Interpreter) (interface{}, error)
}

type DoInstr struct{}

func (di *DoInstr) ParseParams(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("wrong number of params")
	}
	return nil
}

func (di *DoInstr) Execute(interpreter *Interpreter) (interface{}, error) {
	interpreter.ChangeExecutable(true)
	return nil, nil
}

type DontInstr struct{}

func (di *DontInstr) ParseParams(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("wrong number of params")
	}
	return nil
}

func (di *DontInstr) Execute(interpreter *Interpreter) (interface{}, error) {
	interpreter.ChangeExecutable(false)
	return nil, nil
}

type MulInstr struct{ X, Y int }

func (mi *MulInstr) Execute(i *Interpreter) (interface{}, error) {
	return mi.X * mi.Y, nil
}

func (mi *MulInstr) ParseParams(params []byte) (err error) {
	pattern := regexp.MustCompile(`^([1-9][0-9]{0,2}),([1-9][0-9]{0,2})$`)
	parts := pattern.FindAllStringSubmatch(string(params), -1)
	if len(parts) == 0 || len(parts[0]) != 3 {
		return fmt.Errorf("MulInstr must have two int params: %s", string(params))
	}
	if mi.X, err = strconv.Atoi(string(parts[0][1])); err != nil {
		return err
	}
	if mi.Y, err = strconv.Atoi(string(parts[0][2])); err != nil {
		return err
	}
	return nil
}
