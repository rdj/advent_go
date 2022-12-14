package aoc_2022_14

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 1513

type Part2Result int

const Part2Want = 22646

const (
	Empty  = '.'
	Stone  = '#'
	Sand   = 'o'
	Source = '+'
)

type Cave struct {
	src, min, max Point
	state         map[Point]byte
}

func NewCave() Cave {
	src := Point{500, 0}
	return Cave{
		src:   src,
		min:   src,
		max:   src,
		state: map[Point]byte{},
	}
}

func (c *Cave) AddStone(a, b Point) {
	inc := Point{}
	switch {
	case a.x < b.x:
		inc.x = 1
	case a.x > b.x:
		inc.x = -1
	case a.y < b.y:
		inc.y = 1
	case a.y > b.y:
		inc.y = -1
	}

	p := a
	for {
		c.state[p] = Stone
		if p.x < c.min.x {
			c.min.x = p.x
		}
		if p.x > c.max.x {
			c.max.x = p.x
		}
		if p.y > c.max.y {
			c.max.y = p.y
		}
		if p == b {
			break
		}
		p = p.Plus(inc)
	}
}

func (c *Cave) AddSand(inf bool) bool {
	if _, ok := c.state[c.src]; ok {
		return false
	}

	steps := []Point{
		Point{0, 1},  // down
		Point{-1, 1}, // down left
		Point{1, 1},  // down right
	}

	p := c.src
Fall:
	for {
		if p.y == c.max.y+1 {
			if inf {
				return false
			} else {
				c.state[p] = Sand
				return true
			}
		}

		for _, s := range steps {
			n := p.Plus(s)
			if _, ok := c.state[n]; !ok {
				p = n
				continue Fall
			}
		}
		c.state[p] = Sand
		return true
	}
}

func (c *Cave) String() string {
	sb := new(strings.Builder)

	for y := 0; y <= c.max.y; y++ {
		if y > 0 {
			sb.WriteByte('\n')
		}
		for x := c.min.x; x <= c.max.x; x++ {
			p := Point{x, y}
			if b, ok := c.state[p]; ok {
				sb.WriteByte(b)
			} else if p.Eq(c.src) {
				sb.WriteByte(Source)
			} else {
				sb.WriteByte(Empty)
			}
		}
	}

	return sb.String()
}

type Point struct{ x, y int }

func (a Point) Plus(b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func (a Point) Eq(b Point) bool {
	return a.x == b.x && a.y == b.y
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func parsePoint(s string) Point {
	p := Point{}
	lo.Must(fmt.Sscanf(s, "%d,%d", &p.x, &p.y))
	return p
}

func ParseInput(input io.Reader) Cave {
	c := NewCave()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		pts := lo.Map(strings.Split(scanner.Text(), " -> "), func(s string, _ int) Point { return parsePoint(s) })
		for i := 0; i+1 < len(pts); i++ {
			c.AddStone(pts[i], pts[i+1])
		}
	}
	return c
}

func DoPart1(c Cave) Part1Result {
	i := 0
	for c.AddSand(true) {
		i++
	}
	return Part1Result(i)
}

func DoPart2(c Cave) Part2Result {
	i := 0
	for c.AddSand(false) {
		i++
	}
	return Part2Result(i)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
