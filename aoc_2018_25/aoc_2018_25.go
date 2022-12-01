package aoc_2018_25

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 310

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

func absdiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type Point [4]int

func (a Point) Manhattan(b Point) int {
	if len(a) != len(b) {
		panic("mismatched coords")
	}
	dist := 0
	for i := range a {
		dist += absdiff(a[i], b[i])
	}
	return dist
}

type ParsedInput []Point

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) ParsedInput {
	points := make([]Point, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		p := Point{}
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d,%d", &p[0], &p[1], &p[2], &p[3])
		if err != nil {
			panic(fmt.Sprintf("%s: %q", err, scanner.Text()))
		}
		points = append(points, p)
	}
	return points
}

func DoPart1(input ParsedInput) Part1Result {
	cons := make([][]Point, 0)
	con := []Point{input[0]}
	input = input[1:]

	unplaced := map[Point]bool{}
	for _, p := range input {
		unplaced[p] = true
	}

Outer:
	for len(unplaced) > 0 {
		for p := range unplaced {
			if !unplaced[p] {
				continue
			}
			if len(con) == 0 {
				con = append(con, p)
				delete(unplaced, p)
				continue Outer
			}
			for _, m := range con {
				if m.Manhattan(p) <= 3 {
					con = append(con, p)
					delete(unplaced, p)
					continue Outer
				}
			}
		}

		cons = append(cons, con)
		con = make([]Point, 0)
	}
	if len(con) > 0 {
		cons = append(cons, con)
	}

	return Part1Result(len(cons))
}

func DoPart2(input ParsedInput) Part2Result {
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
