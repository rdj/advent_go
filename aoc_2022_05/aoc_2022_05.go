package aoc_2022_05

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result string

const Part1Want = "JDTMRWCQJ"

type Part2Result string

const Part2Want = "VHJDDCWRD"

type ParsedInput [][]byte

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

type Move struct {
	n, from, to int
}

func ParseInput(input io.Reader) ([][]byte, []Move) {
	stacks := make([][]byte, 0)
	moves := make([]Move, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if len(stacks) == 0 {
			n := (len(line) + 1) / 4
			for i := 0; i < n; i++ {
				stacks = append(stacks, make([]byte, 0))
			}
		}

		for i := 0; i < len(stacks); i++ {
			r := line[i*4+1]
			if r >= '0' && r <= '9' {
				continue
			}
			if r == ' ' {
				continue
			}
			stacks[i] = append(stacks[i], r)
		}
	}

	for scanner.Scan() {
		m := Move{}
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &m.n, &m.from, &m.to)
		m.from--
		m.to--
		moves = append(moves, m)
	}

	return stacks, moves
}

func dump(stacks [][]byte) {
	for i := range stacks {
		fmt.Println(string(stacks[i]))
	}
}

func DoPart1(stacks [][]byte, moves []Move) Part1Result {
	for _, m := range moves {
		for i := 0; i < m.n; i++ {
			r := stacks[m.from][0]
			stacks[m.from] = stacks[m.from][1:]
			stacks[m.to] = append([]byte{r}, stacks[m.to]...)
		}
	}

	sb := new(strings.Builder)
	for _, s := range stacks {
		sb.WriteByte(s[0])
	}
	return Part1Result(sb.String())
}

func DoPart2(stacks [][]byte, moves []Move) Part2Result {
	for _, m := range moves {
		// must explicitly copy to avoid overlapping
		src := append([]byte{}, stacks[m.from][0:m.n]...)
		stacks[m.from] = stacks[m.from][m.n:]
		stacks[m.to] = append(src, stacks[m.to]...)
	}

	sb := new(strings.Builder)
	for _, s := range stacks {
		sb.WriteByte(s[0])
	}
	return Part2Result(sb.String())
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
