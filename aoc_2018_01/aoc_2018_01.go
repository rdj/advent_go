package aoc_2018_01

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var _ = fmt.Printf

const inputFile = "input.txt"

type AdventResult int

func openInput() io.Reader {
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

func ParseInput(input io.Reader) []int {
	ints := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		ints = append(ints, n)
	}
	return ints
}

func DoPart1(input []int) AdventResult {
	state := 0
	for _, n := range input {
		state += n
	}
	return AdventResult(state)
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() AdventResult {
	return math.MaxInt
}
