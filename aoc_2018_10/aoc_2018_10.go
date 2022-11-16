package aoc_2018_10

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result string

const Part1Fake = "0xDEAD_BEEF"
const Part1Want = `######....##....######..######..#####...#.......######...####.
.....#...#..#...#............#..#....#..#............#..#....#
.....#..#....#..#............#..#....#..#............#..#.....
....#...#....#..#...........#...#....#..#...........#...#.....
...#....#....#..#####......#....#####...#..........#....#.....
..#.....######..#.........#.....#..#....#.........#.....#..###
.#......#....#..#........#......#...#...#........#......#....#
#.......#....#..#.......#.......#...#...#.......#.......#....#
#.......#....#..#.......#.......#....#..#.......#.......#...##
######..#....#..######..######..#....#..######..######...###.#`

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 10105

type Bodies []Body

type Point struct{ x, y int }

func (bs Bodies) advance(t int) {
	for i := range bs {
		bs[i].advance(t)
	}
}

func (bs Bodies) limits() (xmin, ymin, xmax, ymax int) {
	xmin, xmax = math.MaxInt, math.MinInt
	ymin, ymax = math.MaxInt, math.MinInt

	for _, b := range bs {
		if b.x < xmin {
			xmin = b.x
		}
		if b.x > xmax {
			xmax = b.x
		}
		if b.y < ymin {
			ymin = b.y
		}
		if b.y > ymax {
			ymax = b.y
		}
	}

	return
}

func (bs Bodies) Size() int {
	xmin, ymin, xmax, ymax := bs.limits()
	return (ymax - ymin) + (xmax - xmin)
}

func (bs Bodies) String() string {
	points := map[Point]bool{}
	for _, b := range bs {
		points[Point{b.x, b.y}] = true
	}

	xmin, ymin, xmax, ymax := bs.limits()
	sb := strings.Builder{}

	size := (ymax - ymin) + (xmax - xmin)
	//fmt.Fprintf(&sb, "Size=%d : (%d, %d) - (%d, %d)\n", size, xmin, ymin, xmax, ymax)

	if size < 100 {
		for y := ymin; y <= ymax; y++ {
			if y > ymin {
				sb.WriteRune('\n')
			}
			for x := xmin; x <= xmax; x++ {
				r := '.'
				if points[Point{x, y}] {
					r = '#'
				}
				sb.WriteRune(r)
			}
		}
	}

	return sb.String()
}

type Body struct {
	x, y, dx, dy int
}

func (b *Body) advance(t int) {
	b.x += b.dx * t
	b.y += b.dy * t
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) Bodies {
	bodies := make([]Body, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		body := Body{}
		fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%d, %d>", &body.x, &body.y, &body.dx, &body.dy)
		bodies = append(bodies, body)
	}
	return bodies
}

func simulate(bodies Bodies) (string, int) {
	var prevSize int
	t := 0
	for {
		prevSize = bodies.Size()

		bodies.advance(1)
		size := bodies.Size()

		if size > prevSize {
			bodies.advance(-1)
			// fmt.Println("T = ", t)
			// fmt.Println(bodies)
			return bodies.String(), t
		}

		t += 1
	}

	panic("diverged")
}

func DoPart1(bodies Bodies) Part1Result {
	msg, _ := simulate(bodies)
	return Part1Result(msg)
}

func DoPart2(bodies Bodies) Part2Result {
	_, t := simulate(bodies)
	return Part2Result(t)
}

func Part1() Part1Result {
	//return Part1Fake
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
