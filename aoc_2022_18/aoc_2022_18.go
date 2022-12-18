package aoc_2022_18

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 4548

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

type Point struct{ x, y, z int }

func (a Point) Plus(b Point) Point {
	return Point{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
	}
}

func (p Point) Neighbors() []Point {
	return []Point{
		p.Plus(Point{-1, 0, 0}),
		p.Plus(Point{1, 0, 0}),
		p.Plus(Point{0, -1, 0}),
		p.Plus(Point{0, 1, 0}),
		p.Plus(Point{0, 0, -1}),
		p.Plus(Point{0, 0, 1}),
	}
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) map[Point]bool {
	points := map[Point]bool{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		p := Point{}
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &p.x, &p.y, &p.z)
		if err != nil {
			panic(err)
		}
		points[p] = true
	}
	return points
}

func DoPart1(points map[Point]bool) Part1Result {
	faces := 0
	for p := range points {
		pf := 6
		for _, n := range p.Neighbors() {
			if points[n] {
				pf--
			}
		}
		faces += pf
	}
	return Part1Result(faces)
}

func DoPart2(points map[Point]bool) Part2Result {
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
