package aoc_2022_02

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 11_386

type Part2Result int

const Part2Want = 13_600

type ParsedInput [][]rune

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) ParsedInput {
	lines := make([][]rune, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := make([]rune, 2)
		_, err := fmt.Sscanf(scanner.Text(), "%c %c", &line[0], &line[1])
		if err != nil {
			panic(err)
		}
		lines = append(lines, line)
	}
	return lines
}

func DoPart1(input ParsedInput) Part1Result {

	outcomes := [][]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	//elfscore := 0
	myscore := 0
	for _, r := range input {
		elf := r[0] - 'A'
		me := r[1] - 'X'

		myscore += 1 + int(me)
		myscore += outcomes[elf][me]
	}

	return Part1Result(myscore)
}

func DoPart2(input ParsedInput) Part2Result {
	// lose draw win
	choices := [][]int{
		{2, 0, 1},
		{0, 1, 2},
		{1, 2, 0},
	}

	outcomes := [][]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	//elfscore := 0
	myscore := 0
	for _, r := range input {
		elf := r[0] - 'A'
		choice := r[1] - 'X'

		me := choices[elf][choice]

		myscore += 1 + int(me)
		myscore += outcomes[elf][me]
	}

	return Part2Result(myscore)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
