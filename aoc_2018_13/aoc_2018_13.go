package aoc_2018_13

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result string

const Part1Fake = "0xDEAD_BEEF"
const Part1Want = "43,111"

type Part2Result string

const Part2Fake = "0xDEAD_BEEF"
const Part2Want = "44,56"

type Direction byte

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) TurnLeft() Direction {
	return (d - 1) % 4
}

func (d Direction) TurnRight() Direction {
	return (d + 1) % 4
}

type Point struct{ x, y int }

type Cart struct {
	dir   Direction
	turns int
}

func (c *Cart) Byte() byte {
	switch c.dir {
	case Up:
		return '^'
	case Right:
		return '>'
	case Down:
		return 'v'
	case Left:
		return '>'
	default:
		panic("unreachable")
	}
}

func (c *Cart) TurnForced(corner byte) {
	d := &c.dir
	switch corner {
	case '/':
		switch *d {
		case Up, Down:
			*d = d.TurnRight()
		case Right, Left:
			*d = d.TurnLeft()
		}
	case '\\':
		switch *d {
		case Up, Down:
			*d = d.TurnLeft()
		case Right, Left:
			*d = d.TurnRight()
		}
	}
}

func (c *Cart) TurnIntersection() {
	d := &c.dir
	switch c.turns % 3 {
	case 0:
		*d = d.TurnLeft()
	case 2:
		*d = d.TurnRight()
	}
	c.turns++
}

type Mine struct {
	tiles [][]byte
	carts map[Point]*Cart
}

func (m *Mine) AdvanceUntil(firstCrash bool) Point {
	for {
		points := m.CartPoints()
		if len(points) == 1 {
			return points[0]
		}
		for _, p := range points {
			c, ok := m.carts[p]
			if !ok {
				continue
			}
			delete(m.carts, p)

			switch c.dir {
			case Up:
				p.y -= 1
			case Right:
				p.x += 1
			case Down:
				p.y += 1
			case Left:
				p.x -= 1
			}
			if _, ok := m.carts[p]; ok {
				if firstCrash {
					return p
				}
				delete(m.carts, p)
				continue
			}
			switch t := m.tiles[p.y][p.x]; t {
			case '/', '\\':
				c.TurnForced(t)
			case '+':
				c.TurnIntersection()
			}

			m.carts[p] = c
		}
	}
}

func (m *Mine) CartPoints() []Point {
	points := make([]Point, 0, len(m.carts))
	for p := range m.carts {
		points = append(points, p)
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].y == points[j].y {
			return points[i].x < points[j].x
		} else {
			return points[i].y < points[j].y
		}
	})
	return points
}

func (m *Mine) String() string {
	var b strings.Builder
	for y, row := range m.tiles {
		if y > 0 {
			b.WriteByte('\n')
		}
		for x, c := range row {
			if cart, ok := m.carts[Point{x, y}]; ok {
				b.WriteByte(cart.Byte())
			} else {
				b.WriteByte(c)
			}
		}
	}
	return b.String()
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) *Mine {
	mine := Mine{make([][]byte, 0), make(map[Point]*Cart)}
	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		read := scanner.Bytes()
		bytes := make([]byte, len(read))
		copy(bytes, read)
		for x, b := range bytes {
			var dir Direction
			switch b {
			case '^':
				dir = Up
				bytes[x] = '|'
			case '>':
				dir = Right
				bytes[x] = '-'
			case 'v':
				dir = Down
				bytes[x] = '|'
			case '<':
				dir = Left
				bytes[x] = '-'
			default:
				continue
			}
			mine.carts[Point{x, y}] = &Cart{dir, 0}
		}
		mine.tiles = append(mine.tiles, bytes)
	}
	return &mine
}

func DoPart1(mine *Mine) Part1Result {
	p := mine.AdvanceUntil(true)
	return Part1Result(fmt.Sprintf("%d,%d", p.x, p.y))
}

func DoPart2(mine *Mine) Part2Result {
	p := mine.AdvanceUntil(false)
	return Part2Result(fmt.Sprintf("%d,%d", p.x, p.y))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
