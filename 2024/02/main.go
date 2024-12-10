package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := LoadData(os.Stdin)
	noError(err)

	fmt.Println(len(reports), "are safe")
}

func LoadData(r io.Reader) (reports [][]int, err error) {
	reports = make([][]int, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		levels := strings.Fields(scanner.Text())
		if len(levels) < 2 {
			return nil, fmt.Errorf("number of levels less than 2: %s\n", scanner.Text())
		}

		report := make([]int, 0)
		for _, levelStr := range levels {
			levelInt, err := strconv.Atoi(levelStr)
			if err != nil {
				return nil, err
			}
			report = append(report, levelInt)
		}

		if CheckReportValidity(report, 1) {
			reports = append(reports, report)
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return reports, nil
}

// CheckReportValidity can remove single level (system tolerate a single bad level)
func CheckReportValidity(report []int, possibleFaults int) bool {
	cLevels := len(report)
	if cLevels < 2 {
		return true
	}

	trend := 0
	diffs := make([]int, 0)
	for i := cLevels - 1; i >= 1; i-- {
		diff := report[i] - report[i-1]
		diffs = append([]int{diff}, diffs...)
		if diff < 0 {
			trend -= 1
		} else if diff > 0 {
			trend += 1
		}
	}

	invalidLevels := make([]int, 0)
	for i, diff := range diffs {
		if (trend > 0 && diff >= 1 && diff <= 3) ||
			(trend < 0 && diff <= -1 && diff >= -3) {
			continue
		}
		invalidLevels = append(invalidLevels, i+1)
	}
	if len(invalidLevels) == 0 {
		return true
	}

	// cannot be more than possible faults (+1 is because adjucents may affect each other)
	if len(invalidLevels) > possibleFaults+1 {
		return false
	}

	for _, l := range invalidLevels {
		reportCopy := append([]int{}, report[:l]...)
		if CheckReportValidity(append(reportCopy, report[l+1:]...), -1) {
			return true
		}
		reportCopy = append([]int{}, report[:l-1]...)
		if CheckReportValidity(append(reportCopy, report[l:]...), -1) {
			return true
		}
	}

	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func noError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
