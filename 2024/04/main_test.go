package main_test

import (
	"bytes"
	"testing"

	main "github.com/shlikhota/adventofcode/2024/04"
)

func TestLookUp(t *testing.T) {
	rawInput := []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)
	got := main.LookUp(bytes.NewReader(rawInput), []byte("MAS"))
	expected := 9
	if got != expected {
		t.Errorf("Expected %d XMAS, but got %d", expected, got)
	}
}
