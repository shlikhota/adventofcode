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

		if CheckReportValidity(report) {
			reports = append(reports, report)
		} else {
			fmt.Printf("Level %+v is unsafe\n", levels)
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return reports, nil
}

func CheckReportValidity(report []int) bool {
	cLevels := len(report)
	if cLevels < 2 {
		return true
	}

	for i := cLevels - 1; i >= 1; i-- {
		diffLastTwo := report[i] - report[i-1]

		// it should either gradually increasing or gradually decreasing
		if diffLastTwo == 0 {
			fmt.Println("diffLastTwo == 0")
			return false
		}
		if diffLastTwo < 0 && diffLastTwo < -3 || diffLastTwo > 0 && diffLastTwo > 3 {
			fmt.Println("diffLastTwo > 3", diffLastTwo, report)
			return false
		}

		// no need to check order for last levels
		if i == 1 {
			return true
		}

		// check if either decreasing or increasing order
		diffPrevious := report[i-1] - report[i-2]
		if diffLastTwo > 0 && diffPrevious < 0 ||
			diffLastTwo < 0 && diffPrevious > 0 {

			fmt.Println("diffLastTwo != diffPrevious", diffLastTwo, diffPrevious)
			return false
		}
	}

	return true
}

func noError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
