package aoc_2018_20

import (
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 4432

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 8681

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) []byte {
	data, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}
	return data
}

type Point struct{ x, y int }

func (p Point) move(b byte) Point {
	switch b {
	case 'N':
		return Point{p.x, p.y - 1}
	case 'E':
		return Point{p.x + 1, p.y}
	case 'S':
		return Point{p.x, p.y + 1}
	case 'W':
		return Point{p.x - 1, p.y}
	}
	panic(fmt.Sprintf("bad direction: %c", b))
}

func buildCosts(in []byte) map[Point]int {
	if in[0] != '^' {
		panic("bad input")
	}
	in = in[1:]

	costs := map[Point]int{
		Point{}: 0,
	}

	locs := []Point{
		Point{},
	}

	type Branch struct{ start, end []Point }
	branches := []Branch{{[]Point{Point{}}, []Point{}}}

Reading:
	for _, b := range in {
		switch b {
		default:
			for i := range locs {
				old := locs[i]
				cost := costs[old] + 1

				locs[i] = old.move(b)
				prevCost, dupe := costs[locs[i]]
				if !dupe || cost < prevCost {
					costs[locs[i]] = cost
				}
			}

		case '(':
			branch := Branch{
				start: make([]Point, len(locs)),
				end:   []Point{},
			}
			copy(branch.start, locs)
			branches = append(branches, branch)

		case '|':
			branch := &branches[len(branches)-1]
			dedupe := map[Point]bool{}
			for _, p := range branch.end {
				dedupe[p] = true
			}
			for _, p := range locs {
				if !dedupe[p] {
					branch.end = append(branch.end, p)
				}
			}
			locs = locs[:len(branch.start)]
			copy(locs, branch.start)

		case ')':
			branch := &branches[len(branches)-1]
			dedupe := map[Point]bool{}
			for _, p := range locs {
				dedupe[p] = true
			}
			for _, p := range branch.end {
				if !dedupe[p] {
					locs = append(locs, p)
				}
			}
			branches = branches[:len(branches)-1]

		case '$':
			break Reading
		}
	}

	return costs
}

func longest(in []byte) int {
	costs := buildCosts(in)

	max := 0
	for _, dist := range costs {
		if dist > max {
			max = dist
		}
	}
	return max
}

func count(in []byte, minCost int) int {
	costs := buildCosts(in)

	count := 0
	for _, dist := range costs {
		if dist >= minCost {
			count++
		}
	}
	return count
}

func DoPart1(input []byte) Part1Result {
	longest := longest(input)
	return Part1Result(longest)
}

func DoPart2(input []byte) Part2Result {
	return Part2Result(count(input, 1000))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
