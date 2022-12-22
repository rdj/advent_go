package aoc_2022_22

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 186_128

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

const (
	Open = '.'
	Wall = '#'
)

type Point struct{ x, y int }

func (a Point) Plus(b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func (p Point) Reverse() Point {
	return Point{-p.x, -p.y}
}

const (
	NoTurn = 0
	Left   = 3
	Right  = 1
)

type Move struct {
	turn  int
	steps int
}

// Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^).
const (
	East  = 0
	South = 1
	West  = 2
	North = 3
)

type Grove struct {
	points map[Point]byte
	moves  []Move
}

func NewGrove() Grove {
	return Grove{map[Point]byte{}, []Move{}}
}

type Traveler struct {
	pos    Point
	facing int
}

func ParseInput(input io.Reader) Grove {
	g := NewGrove()
	scanner := bufio.NewScanner(input)

	for y := 1; scanner.Scan(); y++ {
		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			break
		}

		for x, b := range line {
			switch b {
			case Open, Wall:
				g.points[Point{x + 1, y}] = byte(b)
			}
		}
	}

	movestr := scanner.Text()
	for len(movestr) > 0 {
		m := Move{}
		if len(g.moves) > 0 {
			switch movestr[0] {
			case 'L':
				m.turn = Left
			case 'R':
				m.turn = Right
			default:
				panic("RDJ")
			}
			movestr = movestr[1:]
		}

		nlen := 0
		for nlen < len(movestr) && movestr[nlen] >= '0' && movestr[nlen] <= '9' {
			nlen++
		}

		m.steps = lo.Must(strconv.Atoi(movestr[:nlen]))
		g.moves = append(g.moves, m)
		movestr = movestr[nlen:]
	}

	return g
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func (g Grove) Next(a Point, d Point) (Point, byte) {
	b := a.Plus(d)
	if t, ok := g.points[b]; ok {
		return b, t
	}

	r := d.Reverse()
	b = a
	for {
		c := b.Plus(r)
		if _, ok := g.points[c]; !ok {
			break
		}
		b = c
	}
	return b, g.points[b]
}

func (g Grove) Move(from Point, dir Point, n int) Point {
	pos := from
	for i := 0; i < n; i++ {
		p, t := g.Next(pos, dir)
		if t == Wall {
			break
		}
		pos = p
	}
	return pos
}

func DoPart1(g Grove) Part1Result {
	t := Traveler{pos: Point{1, 1}}
	for g.points[t.pos] != Open {
		t.pos.x++
	}

	for _, m := range g.moves {
		t.facing += m.turn
		t.facing %= 4

		dir := Point{}
		switch t.facing {
		case East:
			dir.x = 1
		case South:
			dir.y = 1
		case West:
			dir.x = -1
		case North:
			dir.y = -1
		default:
			panic(fmt.Sprintf("Facing %d", t.facing))
		}

		t.pos = g.Move(t.pos, dir, m.steps)
	}

	return Part1Result(1000*t.pos.y + 4*t.pos.x + t.facing)
}

func DoPart2(grove Grove) Part2Result {
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
