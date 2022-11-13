package aoc_yyyy_mm

import (
	"bufio"
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

func ParseInput(input io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func DoPart1(input []string) AdventResult {
	return AdventResult(0)
}

func DoPart2(input []string) AdventResult {
	return AdventResult(0)
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() AdventResult {
	return DoPart2(ParseInput(openInput()))
}
