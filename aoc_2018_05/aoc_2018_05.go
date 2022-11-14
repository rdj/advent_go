package aoc_2018_05

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

var _ = fmt.Println
var _ = strconv.Atoi

const asciiCaseDiff = 'a' - 'A'
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

func reduce(input []byte) []byte {
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
			if diff == asciiCaseDiff {
				input = append(input[:i], input[i+2:]...)
			}
		}

		if before == len(input) {
			break
		}
	}

	return input
}

func DoPart1(input []byte) AdventResult {
	input = reduce(input)
	return AdventResult(len(input))
}

func makeFilter(cap rune) func(rune) rune {
	lower := cap + asciiCaseDiff
	return func(r rune) rune {
		if r == cap || r == lower {
			return -1
		}
		return r
	}
}

func DoPart2(input []byte) AdventResult {
	shortest := int(^uint(0) >> 1)

	for c := 'A'; c <= 'Z'; c++ {
		filtered := bytes.Map(makeFilter(c), input)
		filtered = reduce(filtered)
		if len(filtered) < shortest {
			shortest = len(filtered)
		}
	}

	return AdventResult(shortest)
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() AdventResult {
	return DoPart2(ParseInput(openInput()))
}
