package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		line      string
		char      uint8
		setCount  int
		size      int
		rowsCount int
	)

	firstX := -1
	firstY := -1
	secondX := -1
	secondY := -1
	thirdX := -1
	thirdY := -1
	fourthX := -1
	fourthY := -1

	for scanner.Scan() {
		line = scanner.Text()

		if setCount == 0 {
			setCount = parseInt(line)
			continue
		}

		if size == 0 {
			size = parseInt(line)
			continue
		}

		if rowsCount != size {
			for x := 0; x < size; x++ {
				char = line[x]

				if char == '*' {
					if firstX == -1 {
						firstX = x
						firstY = rowsCount
					} else {
						secondX = x
						secondY = rowsCount
					}
				}
			}

			rowsCount++

			if rowsCount != size {
				continue
			}
		}

		thirdX, thirdY, fourthX, fourthY = findRectangle(firstX, firstY, secondX, secondY, size)

		printMatrix(firstX, firstY, secondX, secondY, thirdX, thirdY, fourthX, fourthY, size)

		// обнуление
		size = 0
		rowsCount = 0
		firstX = -1
		firstY = -1
		secondX = -1
		secondY = -1
		thirdX = -1
		thirdY = -1
		fourthX = -1
		fourthY = -1
	}
}

func parseInt(line string) int {
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return i
}

func findRectangle(firstX, firstY, secondX, secondY, size int) (int, int, int, int) {
	var (
		thirdX, thirdY, fourthX, fourthY int
	)

	switch {
	// в плоскости X
	case isPointsInPlaneX(firstX, firstY, secondX, secondY):
		thirdX = firstX
		thirdY = firstY + 1

		// на границе матрицы
		if firstY == (size - 1) {
			thirdY = firstY - 1
		}

		fourthX = secondX
		fourthY = thirdY

	// в плоскости Y
	case isPointsInPlaneY(firstX, firstY, secondX, secondY):
		thirdX = firstX + 1

		// на границе матрицы
		if firstX == (size - 1) {
			thirdX = firstX - 1
		}

		thirdY = firstY
		fourthX = thirdX
		fourthY = secondY

	// по диагонали
	default:
		thirdX = secondX
		thirdY = firstY
		fourthX = firstX
		fourthY = secondY
	}

	return thirdX, thirdY, fourthX, fourthY
}

func isPointsInPlaneX(_, firstY, _, secondY int) bool {
	return firstY == secondY
}

func isPointsInPlaneY(firstX, _, secondX, _ int) bool {
	return firstX == secondX
}

var buf bytes.Buffer

func printMatrix(firstX, firstY, secondX, secondY, thirdX, thirdY, fourthX, fourthY, size int) {
	buf.Reset()

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			switch {
			case firstX == x && firstY == y,
				secondX == x && secondY == y,
				thirdX == x && thirdY == y,
				fourthX == x && fourthY == y:
				buf.WriteRune('*')
			default:
				buf.WriteRune('.')
			}
		}

		buf.WriteRune('\n')
	}

	fmt.Print(buf.String())
}
