package aoc_2022_06

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 1953

type Part2Result int

const Part2Want = 2301

type ParsedInput []string

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) ParsedInput {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func indexAfterUniqCount(s string, n int) int {
	for i := n; i < len(s); i++ {
		chars := map[byte]bool{}
		for j := i - n; j < i; j++ {
			chars[s[j]] = true
		}
		if len(chars) == n {
			return i
		}
	}
	panic("did not find")
}

func DoPart1(input ParsedInput) Part1Result {
	r := indexAfterUniqCount(input[0], 4)
	return Part1Result(r)
}

func DoPart2(input ParsedInput) Part2Result {
	r := indexAfterUniqCount(input[0], 14)
	return Part2Result(r)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
