package aoc_2022_03

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 7967

type Part2Result int

const Part2Want = 2716

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

func priority(r rune) int {
	if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	}
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	}
	panic("bad priority")
}

func counts(s string) map[rune]int {
	c := map[rune]int{}
	for _, r := range s {
		c[r]++
	}
	return c
}

func sharedType(s string) rune {
	if len(s)%2 != 0 {
		panic("odd length")
	}

	split := len(s) / 2
	c1 := counts(s[:split])
	c2 := counts(s[split:])

	for r := range c1 {
		if c2[r] != 0 {
			return r
		}
	}

	panic("no shared type")
}

func DoPart1(input ParsedInput) Part1Result {
	sum := 0
	for _, s := range input {
		sum += priority(sharedType(s))
	}

	return Part1Result(sum)
}

func sharedType2(elves []string) rune {
	has := map[rune]int{}
	allFlags := 0
	for i, s := range elves {
		flag := 1 << i
		allFlags |= flag
		for _, r := range s {
			has[r] |= flag
		}
	}
	for r, f := range has {
		if f == allFlags {
			return r
		}
	}
	panic("no shared type")
}

func DoPart2(input ParsedInput) Part2Result {
	sum := 0

	for len(input) > 0 {
		group := input[:3]
		input = input[3:]
		shared := sharedType2(group)
		sum += priority(shared)
	}

	return Part2Result(sum)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
