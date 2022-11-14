package aoc_2018_05

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var _ = fmt.Println
var _ = strconv.Atoi

const inputFile = "input.txt"

type AdventResult int

func openInput() io.Reader {
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

func ParseInput(input io.Reader) []byte {
	b, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	// Input file has a newline at the end
	for b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	return b
}

func DoPart1(input []byte) AdventResult {
	var before int
	for {
		before = len(input)

		for i := 0; i+1 < len(input); i++ {
			a, b := input[i], input[i+1]
			var diff byte
			if a > b {
				diff = a - b
			} else {
				diff = b - a
			}
			if diff == 'a'-'A' {
				input = append(input[:i], input[i+2:]...)
			}
		}

		if before == len(input) {
			break
		}
	}

	return AdventResult(len(input))
}

func DoPart2(input []byte) AdventResult {
	return AdventResult(0)
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() AdventResult {
	return DoPart2(ParseInput(openInput()))
}
