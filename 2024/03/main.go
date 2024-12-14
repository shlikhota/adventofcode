package main

import (
	"fmt"
	"os"
)

func main() {
	interp := NewInterpreter(os.Stdin)
	interp.RegisterCommand("mul", &MulInstr{})
	interp.RegisterCommand("do", &DoInstr{})
	interp.RegisterCommand("don't", &DontInstr{})
	go interp.Run()
	result := 0
	for output := range interp.Output() {
		switch output.(type) {
		case int:
			result += output.(int)
		case error:
			// fmt.Println("Warning:", output.(error).Error())
		}
	}

	fmt.Println("Puzzle answer: ", result)
}
