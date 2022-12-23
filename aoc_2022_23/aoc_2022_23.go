package aoc_2022_23

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

const Part1Want = 4114

type Part2Result int

const Part2Want = 970

type Point struct{ x, y int }

func (a Point) Plus(b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a Point) Min(b Point) Point {
	return Point{min(a.x, b.x), min(a.y, b.y)}
}

func (a Point) Max(b Point) Point {
	return Point{max(a.x, b.x), max(a.y, b.y)}
}

var directions [][]Point

func init() {
	directions = [][]Point{
		{{0, -1}, {-1, -1}, {1, -1}}, // north
		{{0, 1}, {-1, 1}, {1, 1}},    // south
		{{-1, 0}, {-1, -1}, {-1, 1}}, // west
		{{1, 0}, {1, -1}, {1, 1}},    // east
	}
}

type ElfLand struct {
	pos     map[Point]bool
	elapsed int
}

func NewElfLand() *ElfLand {
	return &ElfLand{
		pos: map[Point]bool{},
	}
}

func (e *ElfLand) direction(n int) []Point {
	return directions[(e.elapsed+n)%len(directions)]
}

func (e *ElfLand) bounds() (Point, Point) {
	initial := true
	var min, max Point
	for p := range e.pos {
		if initial {
			min = p
			max = p
			initial = false
			continue
		}

		min = min.Min(p)
		max = max.Max(p)
	}

	return min, max
}

func (e *ElfLand) String() string {
	sb := new(strings.Builder)
	min, max := e.bounds()
	for y := min.y; y <= max.y; y++ {
		if y > min.y {
			sb.WriteByte('\n')
		}
		for x := min.x; x <= max.x; x++ {
			if e.pos[Point{x, y}] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
	}

	return sb.String()
}

func (e *ElfLand) round() int {
	plan := map[Point]*Point{}

	for src := range e.pos {
		// fmt.Println("Elf turn", src)
		foundAny := false
		var claim *Point
		for d := 0; d < len(directions); d++ {
			pts := lo.Map(e.direction(d), func(a Point, _ int) Point { return a.Plus(src) })
			foundOne := false
			for _, p := range pts {
				// if e.pos[p] {
				// 	fmt.Println("  Found neighbor at", p)
				// }
				foundOne = foundOne || e.pos[p]
				foundAny = foundAny || foundOne
				if foundOne {
					break
				}
			}
			if !foundOne && claim == nil {
				//fmt.Println("  Side clear for", pts[0])
				claim = &pts[0]
				if foundAny {
					break
				}
			}
		}
		if foundAny && claim != nil {
			if _, ok := plan[*claim]; ok {
				//fmt.Println("  Double claim on", *claim)
				plan[*claim] = nil
			} else {
				//fmt.Println("  Claiming ", *claim)
				cp := src
				plan[*claim] = &cp
			}
		}
	}

	moves := 0
	for dst, src := range plan {
		if src != nil {
			delete(e.pos, *src)
			e.pos[dst] = true
			moves++
		}
	}

	e.elapsed++

	return moves
}

func ParseInput(input io.Reader) *ElfLand {
	elves := NewElfLand()
	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		for x, r := range scanner.Text() {
			if r == '#' {
				elves.pos[Point{x, y}] = true
			}
		}
	}
	return elves
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func DoPart1(e *ElfLand) Part1Result {
	for r := 0; r < 10; r++ {
		e.round()
	}

	min, max := e.bounds()
	return Part1Result((max.x+1-min.x)*(max.y+1-min.y) - len(e.pos))
}

func DoPart2(e *ElfLand) Part2Result {
	for e.round() > 0 {
	}
	return Part2Result(e.elapsed)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
