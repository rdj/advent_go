package aoc_2018_17

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 30495

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 24899

const (
	Clay   = '#'
	Empty  = '.'
	Spring = '+'
	Pool   = '~'
	Flow   = '|'
)

type Point struct{ x, y int }

type GroundMap map[Point]byte

func (g GroundMap) String() string {
	xmin := int(^uint(0) >> 1)
	xmax := -1
	ymax := -1

	for p := range g {
		if p.x < xmin {
			xmin = p.x
		}
		if p.x > xmax {
			xmax = p.x
		}
		if p.y > ymax {
			ymax = p.y
		}
	}

	sb := new(strings.Builder)

	for y := 0; y <= ymax; y++ {
		if y > 0 {
			sb.WriteByte('\n')
		}
		for x := xmin; x <= xmax; x++ {
			c, ok := g[Point{x, y}]
			if ok {
				sb.WriteByte(c)
			} else {
				sb.WriteByte(Empty)
			}
		}
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

func ParseInput(input io.Reader) GroundMap {
	g := GroundMap{}
	g[Point{500, 0}] = Spring
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var (
			a, b               rune
			aVal, bStart, bEnd int
		)
		_, err := fmt.Sscanf(scanner.Text(),
			"%c=%d, %c=%d..%d",
			&a, &aVal,
			&b, &bStart, &bEnd)
		if err != nil {
			fmt.Println(scanner.Text())
			panic(err)
		}

		for bVal := bStart; bVal <= bEnd; bVal++ {
			var p Point
			switch a {
			case 'x':
				p = Point{aVal, bVal}
			case 'y':
				p = Point{bVal, aVal}
			}
			g[p] = Clay
		}
	}

	return g
}

func (g GroundMap) run() (ymin, ymax int) {
	ymin = int(^uint(0) >> 1)
	ymax = -1
	for p := range g {
		if g[p] != Clay {
			continue
		}
		if p.y < ymin {
			ymin = p.y
		}
		if p.y > ymax {
			ymax = p.y
		}
	}

	g.flow(Point{500, 1}, ymax)
	return
}

func (g GroundMap) flow(p Point, ymax int) {
down:
	for {
		if p.y > ymax {
			return
		}
		if _, set := g[p]; !set {
			g[p] = Flow
		}
		down := Point{p.x, p.y + 1}
		switch g[down] {
		case Clay, Pool:
			break down
		}
		p = down
	}

	left, lBlocked := g.check(p, -1)
	right, rBlocked := g.check(p, 1)
	if lBlocked && rBlocked {
		for x := left.x; x <= right.x; x++ {
			g[Point{x, p.y}] = Pool
		}
		g.flow(Point{p.x, p.y - 1}, ymax)
	} else {
		if !lBlocked {
			g.flow(left, ymax)
		}
		if !rBlocked {
			g.flow(right, ymax)
		}
	}
}

func (g GroundMap) check(p Point, dir int) (bound Point, blocked bool) {
	for {
		down := Point{p.x, p.y + 1}
		if d, set := g[down]; !set || d == Flow {
			return p, false
		}
		if _, set := g[p]; !set {
			g[p] = Flow
		}
		next := Point{p.x + dir, p.y}
		if g[next] == Clay {
			return p, true
		}
		p = next
	}
}

func (g GroundMap) result(countIf func(byte) bool) int {
	ymin, ymax := g.run()
	n := 0
	for p, c := range g {
		if p.y < ymin || p.y > ymax {
			continue
		}
		if countIf(c) {
			n++
		}
	}
	return n
}

func DoPart1(g GroundMap) Part1Result {
	flowOrPool := func(c byte) bool {
		return c == Flow || c == Pool
	}

	return Part1Result(g.result(flowOrPool))
}

func DoPart2(g GroundMap) Part2Result {
	poolOnly := func(c byte) bool {
		return c == Pool
	}
	return Part2Result(g.result(poolOnly))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
