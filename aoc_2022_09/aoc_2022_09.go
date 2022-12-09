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

func DoPart1(input []Move) Part1Result {
	seen := map[Point]bool{}

	head := Point{}
	tail := head
	seen[tail] = true

	for _, m := range input {
		for i := 0; i < m.n; i++ {
			head = head.Move(m.d)
			tail = tail.Follow(head)
			seen[tail] = true
		}
	}

	return Part1Result(len(seen))
}

func DoPart2(input []Move) Part2Result {
	const snakeLen = 10
	snake := make([]Point, snakeLen)
	seen := map[Point]bool{}

	for _, m := range input {
		for i := 0; i < m.n; i++ {
			snake[0] = snake[0].Move(m.d)

			for s := 1; s < len(snake); s++ {
				snake[s] = snake[s].Follow(snake[s-1])
			}

			seen[snake[len(snake)-1]] = true
		}
	}

	return Part2Result(len(seen))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
