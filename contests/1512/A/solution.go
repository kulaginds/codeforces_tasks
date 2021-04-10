package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		line        string
		setCount    int
		arrayLength int
		array       []int
	)
	for scanner.Scan() {
		line = scanner.Text()

		if setCount == 0 {
			setCount = parseInt(line)
			continue
		}

		if arrayLength == 0 {
			arrayLength = parseInt(line)
			continue
		}

		array = parseIntArray(line)

		fmt.Println(whatIndexOfSpy(array))
		arrayLength = 0
		array = nil
	}
}

func parseInt(line string) int {
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return i
}

func parseIntArray(line string) []int {
	parts := strings.Split(line, " ")

	arr := make([]int, 0, len(parts))
	for _, part := range parts {
		arr = append(arr, parseInt(part))
	}

	return arr
}

func whatIndexOfSpy(array []int) int {
	if len(array) <= 2 {
		return 2
	}

	first := array[0]
	second := array[1]
	third := array[2]

	if first == second && second == third {
		for i, num := range array[3:] {
			if first != num {
				return i + 3 + 1
			}
		}

		return 0
	} else {
		// странный один из первых трех
		if first == second {
			if first == third {
				return 2
			} else {
				return 3
			}
		} else {
			if first == third {
				return 2
			} else {
				return 1
			}
		}
	}
}
