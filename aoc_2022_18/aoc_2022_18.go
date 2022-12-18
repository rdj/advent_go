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

const Part2Want = 2588

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Point struct{ x, y, z int }

func (a Point) Plus(b Point) Point {
	return Point{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
	}
}

func (a Point) Min(b Point) Point {
	return Point{
		min(a.x, b.x),
		min(a.y, b.y),
		min(a.z, b.z),
	}
}

func (a Point) Max(b Point) Point {
	return Point{
		max(a.x, b.x),
		max(a.y, b.y),
		max(a.z, b.z),
	}
}

// -i -j -k, i, j, k
var unitVectors []Point = []Point{
	{-1, 0, 0},
	{0, -1, 0},
	{0, 0, -1},
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
}

func (p Point) Neighbors() []Point {
	return lo.Map(unitVectors,
		func(u Point, _ int) Point {
			return p.Plus(u)
		})
}

const (
	Unknown = iota
	Solid
	Inside
	Outside
)

var stateString []string = []string{
	"Unknown",
	"Solid",
	"Inside",
	"Outside",
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) map[Point]byte {
	points := map[Point]byte{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		p := Point{}
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &p.x, &p.y, &p.z)
		if err != nil {
			panic(err)
		}
		points[p] = Solid
	}
	return points
}

func DoPart1(points map[Point]byte) Part1Result {
	faces := 0
	for p := range points {
		pf := 6
		for _, n := range p.Neighbors() {
			if points[n] == Solid {
				pf--
			}
		}
		faces += pf
	}
	return Part1Result(faces)
}

const maxInt = int(^uint(0) >> 1)

func lavalavalava(points map[Point]byte) {
	minP := Point{maxInt, maxInt, maxInt}
	maxP := Point{}
	for p := range points {
		minP = minP.Min(p)
		maxP = maxP.Max(p)
	}
	minP = minP.Plus(Point{-1, -1, -1})
	maxP = maxP.Plus(Point{1, 1, 1})

	q := []Point{minP}
	visited := map[Point]bool{}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if visited[p] {
			continue
		}

		visited[p] = true
		points[p] = Outside

		for _, n := range p.Neighbors() {
			if visited[n] {
				continue
			}
			if n.x < minP.x || n.y < minP.y || n.z < minP.z ||
				n.x > maxP.x || n.y > maxP.y || n.z > maxP.z {
				continue
			}
			if points[n] == Solid {
				continue
			}

			q = append(q, n)
		}
	}
}

func DoPart2(points map[Point]byte) Part2Result {
	lavalavalava(points)

	exposed := 0
	for p, v := range points {
		if v != Solid {
			continue
		}

		for _, n := range p.Neighbors() {
			if points[n] == Outside {
				exposed++
			}
		}
	}

	return Part2Result(exposed)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
