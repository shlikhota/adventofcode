package main_test

import (
	"os"
	"testing"

	main "github.com/shlikhota/adventofcode/2024/02"
)

func TestLoadData(t *testing.T) {
	file, err := os.Open("test.dat")
	noError(t, err)

	reports, err := main.LoadData(file)
	noError(t, err)

	expected := 2
	if len(reports) != expected {
		t.Errorf("Safe reports count is %d, must be %d", len(reports), expected)
	}
}

func TestCheckReportValidity(t *testing.T) {
	type Report struct {
		levels []int
		valid  bool
	}
	testCases := map[string]Report{
		"Safe because the levels are all decreasing by 1 or 2":    {valid: true, levels: []int{7, 6, 4, 2, 1}},
		"Safe because the levels are all increasing by 1 or 2":    {valid: true, levels: []int{4, 6, 7, 9, 11}},
		"Unsafe because 2 7 is an increase of 5":                  {valid: false, levels: []int{1, 2, 7, 8, 9}},
		"Unsafe because 6 2 is a decrease of 4":                   {valid: false, levels: []int{9, 7, 6, 2, 1}},
		"Unsafe because 1 3 is increasing but 3 2 is decreasing":  {valid: false, levels: []int{1, 3, 2, 4, 5}},
		"Unsafe because 4 4 is neither an increase or a decrease": {valid: false, levels: []int{8, 6, 4, 4, 1}},
	}

	for testCaseName, report := range testCases {
		t.Run(testCaseName, func(t *testing.T) {
			if main.CheckReportValidity(report.levels) != report.valid {
				t.Errorf("Report %+v must be valid=%t", report.levels, report.valid)
			}
		})
	}
}

func noError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}
