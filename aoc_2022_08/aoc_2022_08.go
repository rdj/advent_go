package aoc_2022_08

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 1538

type Part2Result int

const Part2Want = 496125

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

func isVisible(a []string, row int, col int) bool {
	if row == 0 || col == 0 || row == len(a)-1 || col == len(a[0])-1 {
		return true
	}

	n := a[row][col]

	blocked := false
	for r := row + 1; r < len(a); r++ {
		if a[r][col] >= n {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for r := row - 1; r >= 0; r-- {
		if a[r][col] >= n {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for c := col + 1; c < len(a[0]); c++ {
		if a[row][c] >= n {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for c := col - 1; c >= 0; c-- {
		if a[row][c] >= n {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	return false
}

func scenicScore(a []string, row int, col int) int {
	if row == 0 || col == 0 || row == len(a)-1 || col == len(a[0])-1 {
		return 0
	}

	n := a[row][col]

	cum := 1
	s := 0

	for r := row + 1; r < len(a); r++ {
		s++
		if a[r][col] >= n {
			break
		}
	}
	cum *= s
	s = 0

	for r := row - 1; r >= 0; r-- {
		s++
		if a[r][col] >= n {
			break
		}
	}
	cum *= s
	s = 0

	for c := col + 1; c < len(a[0]); c++ {
		s++
		if a[row][c] >= n {
			break
		}
	}
	cum *= s
	s = 0

	for c := col - 1; c >= 0; c-- {
		s++
		if a[row][c] >= n {
			break
		}
	}
	cum *= s

	return cum
}

func DoPart1(input ParsedInput) Part1Result {
	trees := 0
	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[0]); c++ {
			if isVisible(input, r, c) {
				//fmt.Printf("(%d, %d): %c\n", r, c, input[r][c])
				trees++
			}
		}
	}

	return Part1Result(trees)
}

func DoPart2(input ParsedInput) Part2Result {
	best := 0

	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[0]); c++ {
			s := scenicScore(input, r, c)
			if s > best {
				best = s
			}
		}
	}

	return Part2Result(best)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
