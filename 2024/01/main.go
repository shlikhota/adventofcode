package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	pointsLeft, pointsRight, err := LoadData(os.Stdin)
	noError(err)

	totalDistance, err := TotalDistance(pointsLeft, pointsRight)
	noError(err)
	fmt.Println("Total distance: ", totalDistance)

	similarityScore := GetSimilarityScore(pointsLeft, pointsRight)
	fmt.Println("Similarity score: ", similarityScore)
}

func LoadData(r io.Reader) (a, b []int, err error) {
	// parse all the numbers into slices
	pointsA := make([]int, 0)
	pointsB := make([]int, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) < 2 {
			return nil, nil, fmt.Errorf("wrong input data: %s\n", scanner.Text())
		}
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, nil, err
		}

		pointsA = append(pointsA, num1)
		pointsB = append(pointsB, num2)
	}
	if scanner.Err() != nil {
		return nil, nil, scanner.Err()
	}

	return pointsA, pointsB, nil
}

func TotalDistance(a, b []int) (int, error) {
	// sort them
	sort.Ints(a)
	sort.Ints(b)

	// sum diff between them
	var result int
	for idx := range a {
		result += diff(a[idx], b[idx])
	}

	return result, nil
}

func GetSimilarityScore(pointsLeft, pointsRight []int) int {
	// frequency map for similarity score
	rightColumnFreqMap := make(map[int]int, 0)
	for _, n := range pointsRight {
		rightColumnFreqMap[n]++
	}

	// count similarity score
	var result int
	for _, n := range pointsLeft {
		if score, exists := rightColumnFreqMap[n]; exists {
			result += n * score
		}
	}

	return result
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func noError(err error) {
	if err != nil {
		fail(err.Error())
	}
}

func fail(format string, a ...any) {
	fmt.Errorf(format, a...)
	os.Exit(1)
}
