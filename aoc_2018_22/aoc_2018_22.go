package aoc_2018_22

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	xMultiplier = 16807 // 7⁵
	yMultiplier = 48271 // prime
	modulus     = 20183 // prime

	rocky  = 0
	wet    = 1
	narrow = 2
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 10204

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

type ParsedInput struct {
	depth uint
	x, y  int
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) ParsedInput {
	p := ParsedInput{}

	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		panic("bad input")
	}
	fmt.Sscanf(scanner.Text(), "depth: %d", &p.depth)

	if !scanner.Scan() {
		panic("bad input")
	}
	fmt.Sscanf(scanner.Text(), "target: %d,%d", &p.x, &p.y)

	return p
}

type Grid [][]uint

func (g Grid) String() string {
	sb := new(strings.Builder)
	for y, row := range g {
		if y > 0 {
			sb.WriteByte('\n')
		}
		for _, n := range row {
			fmt.Fprintf(sb, "%7d", n)
		}
	}
	return sb.String()
}

type Terrain [][]uint

func (g Terrain) String() string {
	sb := new(strings.Builder)
	for y, row := range g {
		if y > 0 {
			sb.WriteByte('\n')
		}
		for _, n := range row {
			var b byte
			switch n % 3 {
			case 0:
				b = '.'
			case 1:
				b = '='
			case 2:
				b = '|'
			}
			sb.WriteByte(b)
		}
	}
	return sb.String()
}

func DoPart1(input ParsedInput) Part1Result {
	geo := make([][]uint, input.y+1)
	ero := make([][]uint, input.y+1)

	for y := 0; y <= input.y; y++ {
		geo[y] = make([]uint, input.x+1)
		ero[y] = make([]uint, input.x+1)

		for x := 0; x <= input.x; x++ {
			switch {
			case x == 0 && y == 0:
				geo[y][x] = 0
			case x == input.x && y == input.y:
				geo[y][x] = 0
			case x == 1 && y == 0:
				geo[y][x] = xMultiplier
			case y == 0:
				geo[y][x] = geo[y][x-1] + xMultiplier
			case x == 0 && y == 1:
				geo[y][x] = yMultiplier
			case x == 0:
				geo[y][x] = geo[y-1][x] + yMultiplier
			default:
				geo[y][x] = ero[y][x-1] * ero[y-1][x]
			}
			geo[y][x] %= modulus
			ero[y][x] = (geo[y][x] + input.depth) % modulus
		}
	}

	// fmt.Println("Geological Index")
	// fmt.Println(Grid(geo))
	// fmt.Println()
	// fmt.Println("Erosion")
	// fmt.Println(Grid(ero))
	// fmt.Println()
	// fmt.Println("Terrain")
	// fmt.Println(Terrain(ero))
	// fmt.Println()

	risk := uint(0)
	for _, row := range ero {
		for _, e := range row {
			risk = (risk + e%3)
		}
	}

	return Part1Result(risk)
}

func DoPart2(input ParsedInput) Part2Result {
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
