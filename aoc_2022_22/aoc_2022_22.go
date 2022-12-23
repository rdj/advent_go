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

const Part2Want = 34426

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

func (p Point) Div(n int) Point {
	return Point{p.x / n, p.y / n}
}

func (p Point) Rem(n int) Point {
	return Point{p.x % n, p.y % n}
}

func (p Point) Times(n int) Point {
	return Point{p.x * n, p.y * n}
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

type EdgeKey struct {
	face, facing int
}

type EdgeValue struct {
	face, facing int
	xform        func(g Grove, p Point) Point
}

var edges map[EdgeKey]EdgeValue

func minusY(g Grove, p Point) Point {
	return Point{p.x, g.faceWidth - p.y - 1}
}

func swapXY(g Grove, p Point) Point {
	return Point{p.y, p.x}
}

func wrapY(g Grove, p Point) Point {
	y := 0
	if p.y == 0 {
		y = g.faceWidth - 1
	}
	return Point{p.x, y}
}

func init() {
	edges = map[EdgeKey]EdgeValue{
		// 1 < to 4 >    4 < to 1 >     -y
		EdgeKey{1, West}: EdgeValue{4, East, minusY},
		EdgeKey{4, West}: EdgeValue{1, East, minusY},

		// 1 ^ to 6 >    6 < to 1 v     x<>y
		EdgeKey{1, North}: EdgeValue{6, East, swapXY},
		EdgeKey{6, West}:  EdgeValue{1, South, swapXY},

		// 2 > to 5 <    5 > to 2 <     -y
		EdgeKey{2, East}: EdgeValue{5, West, minusY},
		EdgeKey{5, East}: EdgeValue{2, West, minusY},

		// 2 v to 3 <    3 > to 2 ^     x<>y
		EdgeKey{2, South}: EdgeValue{3, West, swapXY},
		EdgeKey{3, East}:  EdgeValue{2, North, swapXY},

		// 2 ^ to 6 ^    6 v to 2 v     y-wrap
		EdgeKey{2, North}: EdgeValue{6, North, wrapY},
		EdgeKey{6, South}: EdgeValue{2, South, wrapY},

		// 3 < to 4 v    4 ^ to 3 >     x<>y
		EdgeKey{3, West}:  EdgeValue{4, South, swapXY},
		EdgeKey{4, North}: EdgeValue{3, East, swapXY},

		// 5 v to 6 <    6 > to 5 ^     x<>y
		EdgeKey{5, South}: EdgeValue{6, West, swapXY},
		EdgeKey{6, East}:  EdgeValue{5, North, swapXY},
	}
}

type Grove struct {
	points    map[Point]byte
	moves     []Move
	faceWidth int
	faces     map[Point]int
	cubic     bool
}

func NewGrove() Grove {
	return Grove{map[Point]byte{}, []Move{}, 0, map[Point]int{}, false}
}

type FacePoint struct {
	face int
	p    Point
}

func (g Grove) facePoint(p Point) FacePoint {
	corner := p.Div(g.faceWidth)
	face, ok := g.faces[corner]
	if !ok {
		fmt.Println(g.faces)
		panic(fmt.Sprintf("no face for point %v, corner %v, width %d\n", p, corner, g.faceWidth))
	}

	return FacePoint{face, p.Rem(g.faceWidth)}
}

func (g Grove) faceCorner(n int) Point {
	for c, x := range g.faces {
		if n == x {
			return c
		}
	}
	panic("bad face")
}

func (g Grove) worldPoint(fp FacePoint) Point {
	c := g.faceCorner(fp.face)
	return c.Times(g.faceWidth).Plus(fp.p)
}

type Traveler struct {
	pos    Point
	facing int
}

func ParseInput(input io.Reader) Grove {
	g := NewGrove()
	g.faceWidth = int(^uint(0) >> 1)

	scanner := bufio.NewScanner(input)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			break
		}

		xMin := int(^uint(0) >> 1)
		xMax := 0
		for x, b := range line {
			switch b {
			case Open, Wall:
				g.points[Point{x, y}] = byte(b)
				if x < xMin {
					xMin = x
				}
				if x > xMax {
					xMax = x
				}
			}
		}
		width := 1 + xMax - xMin
		if width < g.faceWidth {
			g.faceWidth = width
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

	n := 1
FindFaces:
	for y := 0; y < 6; y++ {
		for x := 0; x < 6; x++ {
			p := Point{x * g.faceWidth, y * g.faceWidth}
			if _, ok := g.points[p]; ok {
				g.faces[Point{x, y}] = n
				n++
				if n > 6 {
					break FindFaces
				}
			}
		}
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

func (g Grove) Next(a Point, facing int) (Point, int, byte) {
	d := Point{}
	switch facing {
	case East:
		d.x = 1
	case South:
		d.y = 1
	case West:
		d.x = -1
	case North:
		d.y = -1
	default:
		panic(fmt.Sprintf("Facing %d", facing))
	}

	b := a.Plus(d)
	if t, ok := g.points[b]; ok {
		return b, facing, t
	}

	if !g.cubic {
		r := d.Reverse()
		b = a
		for {
			c := b.Plus(r)
			if _, ok := g.points[c]; !ok {
				break
			}
			b = c
		}
		return b, facing, g.points[b]
	}

	originFP := g.facePoint(a)
	edge, ok := edges[EdgeKey{originFP.face, facing}]
	if !ok {
		panic(fmt.Sprintf("LOST IN SPACE: %v", EdgeKey{originFP.face, facing}))
	}
	destFP := FacePoint{edge.face, edge.xform(g, originFP.p)}
	wp := g.worldPoint(destFP)
	return wp, edge.facing, g.points[wp]
}

func (g Grove) Move(pos Point, dir int, n int) (Point, int) {
	for i := 0; i < n; i++ {
		p, d, t := g.Next(pos, dir)
		if t == Wall {
			break
		}
		pos = p
		dir = d
	}
	return pos, dir
}

func (g Grove) Answer() int {
	t := Traveler{}
	for g.points[t.pos] != Open {
		t.pos.x++
	}

	for _, m := range g.moves {
		t.facing += m.turn
		t.facing %= 4

		t.pos, t.facing = g.Move(t.pos, t.facing, m.steps)
	}

	return 1000*(t.pos.y+1) + 4*(t.pos.x+1) + t.facing
}

func DoPart1(g Grove) Part1Result {
	return Part1Result(g.Answer())
}

func DoPart2(g Grove) Part2Result {
	g.cubic = true
	return Part2Result(g.Answer())
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
