package aoc_2018_08

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var _ = fmt.Println

const inputFile = "input.txt"

type ParsedInput []int
type Part1Result int
type Part2Result int

func openInput() io.Reader {
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

func ParseInput(input io.Reader) ParsedInput {
	ints := make([]int, 0)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if nil != err {
			panic("bad input")
		}
		ints = append(ints, n)
	}
	return ints
}

func totalMetadata(ints []int) ([]int, int) {
	// header
	nkids := ints[0]
	nmeta := ints[1]
	ints = ints[2:]

	// children
	sum := 0
	for i := 0; i < nkids; i++ {
		var kidsum int
		ints, kidsum = totalMetadata(ints)
		sum += kidsum
	}

	// metadata
	for i := 0; i < nmeta; i++ {
		sum += ints[i]
	}
	return ints[nmeta:], sum
}

func nodeValue(ints []int) ([]int, int) {
	// header
	nkids := ints[0]
	nmeta := ints[1]
	ints = ints[2:]

	// children
	kids := make([]int, nkids)
	for i := 0; i < nkids; i++ {
		ints, kids[i] = nodeValue(ints)
	}

	// metadata
	value := 0
	for i := 0; i < nmeta; i++ {
		meta := ints[i]
		if nkids == 0 {
			value += meta
			continue
		}
		if meta > 0 && meta <= len(kids) {
			value += kids[meta-1]
		}
	}

	return ints[nmeta:], value
}

func DoPart1(input ParsedInput) Part1Result {
	_, sum := totalMetadata(input)
	return Part1Result(sum)
}

func DoPart2(input ParsedInput) Part2Result {
	_, value := nodeValue(input)
	return Part2Result(value)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
