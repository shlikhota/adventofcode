package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	count := LookUp(os.Stdin, []byte("XMAS"))
	fmt.Println("Puzzle answer:", count)
}

const Left = 0
const TopLeft = 1
const Top = 2
const TopRight = 3

func LookUp(input io.Reader, needle []byte) (result int) {
	if len(needle) == 0 {
		return 0
	}

	var previousRowStates [][4]int

	sc := bufio.NewScanner(input)
	var irow int
	var previousIndex int

	for sc.Scan() {
		row := sc.Bytes()
		rowLength := len(row)
		currentRowStates := make([][4]int, rowLength)
		if len(previousRowStates) == 0 {
			previousRowStates = make([][4]int, rowLength)
		}

		for col := 0; col < rowLength; col++ {
			currentByte := row[col]

			// check horizontal orientation
			previousIndex = 0
			if col > 0 {
				previousIndex = currentRowStates[col-1][Left]
			}
			leftIndex, found := getNextIndex(needle, currentByte, previousIndex)
			currentRowStates[col][Left] = leftIndex
			if found {
				result++
				fmt.Printf("horizontal\t(%d)\t[%2d,%2d]\n", result, irow, col)
			}

			// check diagonal orientation (left to right)
			previousIndex = 0
			if col > 0 {
				previousIndex = previousRowStates[col-1][TopLeft]
			}
			topLeftIndex, found := getNextIndex(needle, currentByte, previousIndex)
			currentRowStates[col][TopLeft] = topLeftIndex
			if found {
				result++
				fmt.Printf("diagonal(l)\t(%d)\t[%2d,%2d]\n", result, irow, col)
			}

			// check vertical orientation
			previousIndex = previousRowStates[col][Top]
			topIndex, found := getNextIndex(needle, currentByte, previousIndex)
			currentRowStates[col][Top] = topIndex
			if found {
				result++
				fmt.Printf("vertical\t(%d)\t[%2d,%2d]\n", result, irow, col)
			}

			// check diagonal orientation (right to left)
			previousIndex = 0
			if col < rowLength-1 {
				previousIndex = previousRowStates[col+1][TopRight]
			}
			topRightIndex, found := getNextIndex(needle, currentByte, previousIndex)
			currentRowStates[col][TopRight] = topRightIndex
			if found {
				result++
				fmt.Printf("diagonal(r)\t(%d)\t[%2d,%2d]\n", result, irow, col)
			}
			irow++
		}

		previousRowStates = currentRowStates
	}
	return result
}

func getNextIndex(needle []byte, currentByte byte, previousIndex int) (nextIndex int, found bool) {
	// end of found before can be beginning of the new one
	if len(needle) == abs(previousIndex) {
		previousIndex = -sign(previousIndex)
	}
	index := abs(previousIndex)
	if previousIndex < 0 && needle[len(needle)-1-index] == currentByte {
		nextIndex = previousIndex - 1
	} else if previousIndex > 0 && needle[index] == currentByte {
		nextIndex = previousIndex + 1
	} else {
		if needle[0] == currentByte {
			nextIndex = 1
		} else if needle[len(needle)-1] == currentByte {
			nextIndex = -1
		}
	}
	return nextIndex, len(needle) == abs(nextIndex)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}