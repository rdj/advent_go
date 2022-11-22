package aoc_2018_18

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

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 481_290

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 180_752

const (
	Outside = 0
	Open    = '.'
	Trees   = '|'
	Lumber  = '#'
)

type Point struct{ x, y int }

func (p Point) neighbors() []Point {
	return []Point{
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y + 1},
	}
}

type Acreage struct {
	dim   int
	acres []byte
}

func NewAcreage(dim int) *Acreage {
	return &Acreage{dim, make([]byte, dim*dim)}
}

func (a *Acreage) advance(steps int) (base int, cycle int) {
	seen := map[string]int{}
	swap := NewAcreage(a.dim)

	for n := 0; n < steps; n++ {
		state := a.String()
		if prev, ok := seen[state]; ok {
			return prev, n - prev
		} else {
			seen[state] = n
		}
		p := Point{}
		for p.y = 0; p.y < a.dim; p.y++ {
			for p.x = 0; p.x < a.dim; p.x++ {
				swap.set(p, a.advanceOne(p))
			}
		}
		a.acres, swap.acres = swap.acres, a.acres
	}

	return
}

func (a *Acreage) advanceOne(p Point) byte {
	neighs := p.neighbors()
	switch a.get(p) {
	case Open:
		trees := 0
		for _, n := range neighs {
			if a.get(n) == Trees {
				trees++
				if trees >= 3 {
					return Trees
				}
			}
		}
		return Open

	case Trees:
		lumber := 0
		for _, n := range neighs {
			if a.get(n) == Lumber {
				lumber++
				if lumber >= 3 {
					return Lumber
				}
			}
		}
		return Trees

	case Lumber:
		lumber, trees := 0, 0
		for _, n := range neighs {
			c := a.get(n)
			if c == Lumber {
				lumber++
			}
			if c == Trees {
				trees++
			}
			if lumber >= 1 && trees >= 1 {
				return Lumber
			}
		}
		return Open
	}

	panic("unreachable")
}

func (a *Acreage) get(p Point) byte {
	if p.x < 0 || p.y < 0 || p.x >= a.dim || p.y >= a.dim {
		return Outside
	}
	return a.acres[p.y*a.dim+p.x]
}

func (a *Acreage) resourceValue() int {
	var trees, lumber int
	for _, c := range a.acres {
		switch c {
		case Trees:
			trees++
		case Lumber:
			lumber++
		}
	}
	return trees * lumber
}

func (a *Acreage) set(p Point, b byte) {
	a.acres[p.y*a.dim+p.x] = b
}

func (a *Acreage) String() string {
	sb := new(strings.Builder)
	for y := 0; y < a.dim; y++ {
		sb.Write(a.acres[a.dim*y : a.dim*(y+1)])
		sb.WriteByte('\n')
	}
	return sb.String()
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) *Acreage {
	var a *Acreage = nil
	scanner := bufio.NewScanner(input)
	for y := 0; scanner.Scan(); y++ {
		if a == nil {
			a = NewAcreage(len(scanner.Text()))
		}
		for x, r := range scanner.Text() {
			a.set(Point{x, y}, byte(r))
		}
	}
	return a
}

func DoPart1(a *Acreage) Part1Result {
	a.advance(10)
	return Part1Result(a.resourceValue())
}

func DoPart2(a *Acreage) Part2Result {
	init := NewAcreage(a.dim)
	copy(init.acres, a.acres)

	n := 1_000_000_000
	base, cycle := a.advance(n)

	init.advance(base + (n-base)%cycle)
	return Part2Result(init.resourceValue())
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
