package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// parse all the numbers into slices
	pointsA := make([]int, 0)
	pointsB := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) < 2 {
			Fail("wrong input data: %s", scanner.Text())
		}
		num1, err := strconv.Atoi(numbers[0])
		NoError(err)
		num2, err := strconv.Atoi(numbers[1])
		NoError(err)

		pointsA = append(pointsA, num1)
		pointsB = append(pointsB, num2)
	}
	NoError(scanner.Err())

	// sort them
	sort.Ints(pointsA)
	sort.Ints(pointsB)

	// sum diff between them
	var result int
	for idx := range pointsA {
		result += diff(pointsA[idx], pointsB[idx])
	}

	// print the result
	fmt.Println(result)
	return
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func NoError(err error) {
	if err != nil {
		Fail(err.Error())
	}
}

func Fail(format string, a ...any) {
	fmt.Errorf(format, a)
	os.Exit(1)
}
