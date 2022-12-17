package aoc_2022_17

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

const Part1Want = 3098

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

type Point struct{ x, y int }

func (a Point) Plus(b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

type Shape struct {
	points []Point
}

var shapes = []Shape{
	{[]Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},         // -
	{[]Point{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}}, // +
	{[]Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}}, // L (backwards)
	{[]Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},         // |
	{[]Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},         // []
}

func (s Shape) at(pos Point) []Point {
	return lo.Map(s.points, func(p Point, _ int) Point { return pos.Plus(p) })
}

func shape(n int) *Shape {
	return &shapes[n%len(shapes)]
}

type Piece struct {
	shape *Shape
	pos   Point
}

func NewPiece(n, height int) *Piece {
	p := new(Piece)
	p.shape = shape(n)
	p.pos = Point{2, height + 3}
	return p
}

func (p *Piece) Place(cave map[Point]bool) int {
	maxY := 0
	for _, p := range p.shape.at(p.pos) {
		cave[p] = true
		if p.y > maxY {
			maxY = p.y
		}
	}
	return maxY
}

func (p *Piece) TryMove(delta Point, width int, cave map[Point]bool) bool {
	pos := p.pos.Plus(delta)
	points := p.shape.at(pos)

	for _, p := range points {
		if p.x < 0 || p.x >= width {
			return false
		}
		if p.y < 0 {
			return false
		}
		if cave[p] {
			return false
		}
	}

	p.pos = pos
	return true
}

func play(moves string, width, totalPieces int) int {
	height := 0
	cave := map[Point]bool{}

	m := 0

	for n := 0; n < totalPieces; n++ {
		p := NewPiece(n, height)

		for {
			delta := Point{0, 0}
			switch moves[m%len(moves)] {
			case '<':
				delta.x = -1
			case '>':
				delta.x = 1
			}
			m++
			p.TryMove(delta, width, cave)

			if !p.TryMove(Point{0, -1}, width, cave) {
				y := p.Place(cave)
				if y+1 > height {
					height = y + 1
				}
				//dump(cave, width, height)
				break
			}
		}
	}

	return height
}

func dump(cave map[Point]bool, width, height int) {
	sb := new(strings.Builder)
	for y := height; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if cave[Point{x, y}] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines[0]
}

func DoPart1(input string) Part1Result {
	return Part1Result(play(input, 7, 2022))
}

func DoPart2(input string) Part2Result {
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
