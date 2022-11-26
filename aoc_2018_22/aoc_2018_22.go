package aoc_2018_22

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	xMultiplier = 16807 // 7⁵
	yMultiplier = 48271 // prime
	modulus     = 20183 // prime

	rocky  = 0
	wet    = 1
	narrow = 2
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 10204

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

type Point struct{ x, y int }

func (p Point) Neighbors() []Point {
	n := make([]Point, 4)
	if p.y > 0 {
		n = append(n, p.Up())
	}
	if p.x > 0 {
		n = append(n, p.Left())
	}
	n = append(n, p.Right(), p.Down())
	return n
}

func (a Point) Eq(b Point) bool {
	return a.x == b.x && a.y == b.y
}

func (p Point) Up() Point {
	return Point{p.x, p.y - 1}
}

func (p Point) Left() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) Right() Point {
	return Point{p.x + 1, p.y}
}

func (p Point) Down() Point {
	return Point{p.x, p.y + 1}
}

type Cave struct {
	depth    int
	target   Point
	geo, ero map[Point]int
}

func NewCave() *Cave {
	c := new(Cave)
	c.geo = make(map[Point]int)
	c.ero = make(map[Point]int)
	return c
}

func (c *Cave) Geology(p Point) int {
	if g, ok := c.geo[p]; ok {
		return g
	}

	var g int
	switch {
	case p.Eq(Point{0, 0}) || p.Eq(c.target):
		g = 0
	case p.Eq(Point{1, 0}):
		g = xMultiplier
	case p.y == 0:
		g = c.Geology(p.Left()) + xMultiplier
	case p.Eq(Point{0, 1}):
		g = yMultiplier
	case p.x == 0:
		g = c.Geology(p.Up()) + yMultiplier
	default:
		g = c.Erosion(p.Left()) * c.Erosion(p.Up())
	}

	g %= modulus
	c.geo[p] = g
	return g
}

func (c *Cave) Erosion(p Point) int {
	if e, ok := c.ero[p]; ok {
		return e
	}

	e := (c.Geology(p) + c.depth) % modulus
	c.ero[p] = e
	return e
}

func (c *Cave) Terrain(p Point) int {
	return c.Erosion(p) % 3
}

func (c *Cave) TotalRisk() int {
	risk := 0
	p := Point{}
	for p.y = 0; p.y <= c.target.y; p.y++ {
		for p.x = 0; p.x <= c.target.x; p.x++ {
			risk += c.Terrain(p)
		}
	}

	return risk
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) *Cave {
	c := NewCave()

	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		panic("bad input")
	}
	fmt.Sscanf(scanner.Text(), "depth: %d", &c.depth)

	if !scanner.Scan() {
		panic("bad input")
	}
	fmt.Sscanf(scanner.Text(), "target: %d,%d", &c.target.x, &c.target.y)

	return c
}

type Grid [][]int

func (g Grid) String() string {
	sb := new(strings.Builder)
	for y, row := range g {
		if y > 0 {
			sb.WriteByte('\n')
		}
		for _, n := range row {
			fmt.Fprintf(sb, "%7d", n)
		}
	}
	return sb.String()
}

type Terrain [][]int

func (g Terrain) String() string {
	sb := new(strings.Builder)
	for y, row := range g {
		if y > 0 {
			sb.WriteByte('\n')
		}
		for _, n := range row {
			var b byte
			switch n % 3 {
			case 0:
				b = '.'
			case 1:
				b = '='
			case 2:
				b = '|'
			}
			sb.WriteByte(b)
		}
	}
	return sb.String()
}

func DoPart1(c *Cave) Part1Result {
	return Part1Result(c.TotalRisk())
}

func DoPart2(c *Cave) Part2Result {
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
