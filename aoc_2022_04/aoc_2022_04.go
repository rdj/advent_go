package aoc_2022_04

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 582

type Part2Result int

const Part2Want = 893

type ParsedInput []Entry

type Entry struct{ amin, amax, bmin, bmax int }

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) ParsedInput {
	lines := make([]Entry, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		e := Entry{}
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &e.amin, &e.amax, &e.bmin, &e.bmax)
		lines = append(lines, e)
	}
	return lines
}

func DoPart1(input ParsedInput) Part1Result {
	n := 0
	for _, e := range input {
		switch {
		case e.amin >= e.bmin && e.amax <= e.bmax,
			e.bmin >= e.amin && e.bmax <= e.amax:
			n++
		}
	}
	return Part1Result(n)
}

func DoPart2(input ParsedInput) Part2Result {
	n := 0
	for _, e := range input {
		switch {
		case e.amin <= e.bmin && e.amax >= e.bmin ||
			e.bmin <= e.amin && e.bmax >= e.amin:
			n++
		}
	}
	return Part2Result(n)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
