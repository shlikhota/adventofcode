package main_test

import (
	"os"
	"testing"

	main "github.com/shlikhota/adventofcode/2024/01"
)

func TestTotalDistance(t *testing.T) {
	file, err := os.Open("test.dat")
	noError(t, err)
	pointsLeft, pointsRight, err := main.LoadData(file)
	noError(t, err)

	totalDistance, err := main.TotalDistance(pointsLeft, pointsRight)
	noError(t, err)

	expected := 11
	if totalDistance != expected {
		t.Errorf("totalDistance is %d, must be %d", totalDistance, expected)
	}
}

func TestGetSimilarityScore(t *testing.T) {
	file, err := os.Open("test.dat")
	noError(t, err)
	pointsLeft, pointsRight, err := main.LoadData(file)
	noError(t, err)

	similarityScore := main.GetSimilarityScore(pointsLeft, pointsRight)

	expected := 31
	if similarityScore != expected {
		t.Errorf("similarityScore is %d, must be %d", similarityScore, expected)
	}
}

func noError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}
