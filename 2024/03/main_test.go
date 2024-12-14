package main_test

import (
	"bytes"
	"testing"

	main "github.com/shlikhota/adventofcode/2024/03"
)

func TestParseInstructions(t *testing.T) {
	testData := bytes.NewReader([]byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`))
	interp := main.NewInterpreter(testData)
	interp.RegisterCommand("mul", &main.MulInstr{})
	interp.RegisterCommand("do", &main.DoInstr{})
	interp.RegisterCommand("don't", &main.DontInstr{})
	go interp.Run()
	expectedResults := make(chan int, 2)
	for _, i := range []int{8, 40} {
		expectedResults <- i
	}
	for output := range interp.Output() {
		switch output.(type) {
		case int:
			expected := <-expectedResults
			got := output.(int)
			if got != expected {
				t.Errorf("Expected %d, got %d", expected, got)
			}
		}
	}
}
