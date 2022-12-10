package aoc_2022_10

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 16480

type Part2Result string

const Part2Want = `###..#....####.####.#..#.#....###..###..
#..#.#....#....#....#..#.#....#..#.#..#.
#..#.#....###..###..#..#.#....#..#.###..
###..#....#....#....#..#.#....###..#..#.
#....#....#....#....#..#.#....#....#..#.
#....####.####.#.....##..####.#....###..`

type Instruction struct {
	op  string
	arg int
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) []Instruction {
	prog := make([]Instruction, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		i := Instruction{}

		_, err := fmt.Sscanf(line, "%s %d", &i.op, &i.arg)
		if err == nil {
			prog = append(prog, i)
			continue
		}

		_, err = fmt.Sscanf(line, "%s", &i.op)
		if err != nil {
			panic(err)
		}
		prog = append(prog, i)
	}
	return prog
}

func sum(ints []int) int {
	t := 0
	for _, n := range ints {
		t += n
	}
	return t
}

func cycleDeltas(prog []Instruction) []int {
	x := []int{1}

	for _, i := range prog {
		switch i.op {
		case "noop":
			x = append(x, 0)

		case "addx":
			x = append(x, 0)
			x = append(x, i.arg)
		}
	}

	return x
}

func DoPart1(prog []Instruction) Part1Result {
	deltas := cycleDeltas(prog)

	r := 0
	for c := 20; c <= 220; c += 40 {
		xval := sum(deltas[0:c])
		str := xval * c
		r += str
	}

	return Part1Result(r)
}

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func DoPart2(prog []Instruction) Part2Result {
	deltas := cycleDeltas(prog)

	x := 0

	sb := new(strings.Builder)
	for t, d := range deltas {
		if t == 240 {
			break
		}

		if t > 0 && t%40 == 0 {
			sb.WriteByte('\n')
		}

		x += d
		px := t % 40
		if absdiff(x, px) <= 1 {
			sb.WriteByte('#')
		} else {
			sb.WriteByte('.')
		}
	}

	//fmt.Println(sb)

	return Part2Result(sb.String())
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
