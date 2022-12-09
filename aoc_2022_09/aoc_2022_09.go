package aoc_2022_09

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 6011

type Part2Result int

const Part2Want = 2419

type Move struct {
	d byte
	n int
}

type Point struct {
	x, y int
}

func (p Point) Up() Point {
	return Point{p.x, p.y - 1}
}

func (p Point) Down() Point {
	return Point{p.x, p.y + 1}
}

func (p Point) Left() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) Right() Point {
	return Point{p.x + 1, p.y}
}

func (p Point) Move(d byte) Point {
	switch d {
	case 'U':
		p = p.Up()
	case 'D':
		p = p.Down()
	case 'L':
		p = p.Left()
	case 'R':
		p = p.Right()
	}

	return p
}

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (a Point) Manhattan(b Point) int {
	return absdiff(a.x, b.x) + absdiff(a.y, b.y)
}

func (a Point) Eq(b Point) bool {
	return a.x == b.x && a.y == b.y
}

func (tail Point) Follow(head Point) Point {
	switch {
	case head.Manhattan(tail) < 2:
		// no op

	case head.x == tail.x && head.y > tail.y:
		tail = tail.Down()

	case head.x == tail.x && head.y < tail.y:
		tail = tail.Up()

	case head.y == tail.y && head.x > tail.x:
		tail = tail.Right()

	case head.y == tail.y && head.x < tail.x:
		tail = tail.Left()

	case head.Manhattan(tail) == 2:
		// no op

	default:
		if head.x > tail.x {
			tail = tail.Right()
		} else {
			tail = tail.Left()
		}
		if head.y > tail.y {
			tail = tail.Down()
		} else {
			tail = tail.Up()
		}
	}

	return tail
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) []Move {
	lines := make([]Move, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		m := Move{}
		_, err := fmt.Sscanf(scanner.Text(), "%c %d", &m.d, &m.n)
		if nil != err {
			panic(err)
		}
		lines = append(lines, m)
	}
	return lines
}

func DoSnake(moves []Move, snakeLen int) int {
	snake := make([]Point, snakeLen)
	seen := map[Point]bool{}

	for _, m := range moves {
		for i := 0; i < m.n; i++ {
			snake[0] = snake[0].Move(m.d)

			for s := 1; s < len(snake); s++ {
				snake[s] = snake[s].Follow(snake[s-1])
			}

			seen[snake[len(snake)-1]] = true
		}
	}

	return len(seen)
}

func DoPart1(input []Move) Part1Result {
	r := DoSnake(input, 2)
	return Part1Result(r)
}

func DoPart2(input []Move) Part2Result {
	r := DoSnake(input, 10)
	return Part2Result(r)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
