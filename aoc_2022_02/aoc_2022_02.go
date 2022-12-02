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

type ParsedInput [][]int

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) ParsedInput {
	lines := make([][]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := make([]int, 2)
		_, err := fmt.Sscanf(scanner.Text(), "%c %c", &line[0], &line[1])
		if err != nil {
			panic(err)
		}
		line[0] -= 'A'
		line[1] -= 'X'
		lines = append(lines, line)
	}
	return lines
}

// rock paper scissors
// outcomes[elf][me] = my score for the round
var outcomes [][]int = [][]int{
	{3, 6, 0},
	{0, 3, 6},
	{6, 0, 3},
}

func score(elf, me int) int {
	return 1 + me + outcomes[elf][me]
}

func DoPart1(input ParsedInput) Part1Result {
	myscore := 0
	for _, r := range input {
		myscore += score(r[0], r[1])
	}

	return Part1Result(myscore)
}

func DoPart2(input ParsedInput) Part2Result {
	// lose draw win
	// choices[elf][desiredOutcome] = my play
	choices := [][]int{
		{2, 0, 1},
		{0, 1, 2},
		{1, 2, 0},
	}

	myscore := 0
	for _, r := range input {
		elf := r[0]
		outcome := r[1]
		myscore += score(elf, choices[elf][outcome])
	}

	return Part2Result(myscore)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
