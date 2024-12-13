package main_test

import (
	"bytes"
	"testing"

	main "github.com/shlikhota/adventofcode/2024/03"
)

func TestParseInstructions(t *testing.T) {
	testData := bytes.NewReader([]byte("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"))
	instructions := main.ParseInstructions(testData)
	expectedResults := []int{8, 25, 88, 40}
	for i, instruction := range instructions {
		expected := expectedResults[i]
		got := instruction.Execute().(int)
		if got != expected {
			t.Errorf("Instruction %d expected %d, got %d", i, expected, got)
		}
	}
}
