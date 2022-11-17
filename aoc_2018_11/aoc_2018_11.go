package aoc_2018_11

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result string

const Part1Fake = "0xDEAD_BEEF"
const Part1Want = "235,14"

type Part2Result string

const Part2Fake = "0xDEAD_BEEF"
const Part2Want = "237,227,14"

type ParsedInput int

// x,y coordinate (1-indexed)
// n = 1133
// x..x+2, y..y+2
// (x+10)*y + n)*(x+10) / 100 % 10 - 5
func singleValue(x, y, n int) int {
	x += 1
	y += 1

	id := x + 10
	value := id * y
	value += n
	value *= id
	value /= 100
	value %= 10
	value -= 5

	return value
}

func singleValues(d, n int) [][]int {
	values := make([][]int, d)
	for i := range values {
		values[i] = make([]int, d)
	}
	for y := range values {
		for x := range values[y] {
			values[y][x] = singleValue(x, y, n)
		}
	}

	return values
}

// Calculates a Summed-area table, aka integral image. In a single
// pass, each position is transformed into the sum of all positions at
// lower coordinates.
//
// See https://en.wikipedia.org/wiki/Summed-area_table
func partials(d, n int) [][]int {
	singles := singleValues(d, n)
	for y := range singles {
		for x := range singles[y] {
			n := &singles[y][x]
			if x > 0 {
				if y > 0 {
					*n -= singles[y-1][x-1]
				}
				*n += singles[y][x-1]
			}
			if y > 0 {
				*n += singles[y-1][x]
			}
		}
	}
	return singles
}

// From the table calculated by `partials`, any arbitrary sub-area
// sum can be calculated by simple combination of four coordinates.
// The linked wikipedia article has an illustration.
func bestSubSquare(d int, partials [][]int) (best, xbest, ybest int) {
	best, xbest, ybest = math.MinInt, math.MinInt, math.MinInt

	// easier to track the bottom right corner for the calculations ...
	for y := d - 1; y < len(partials); y++ {
		for x := d - 1; x < len(partials[y]); x++ {
			n := partials[y][x]
			if x-d >= 0 {
				n -= partials[y][x-d]
				if y-d >= 0 {
					n += partials[y-d][x-d]
				}
			}
			if y-d >= 0 {
				n -= partials[y-d][x]
			}

			if n > best {
				best = n
				xbest = x
				ybest = y
			}
		}
	}

	// ... but the top-left corner is returned to the caller
	return best, xbest + 1 - d, ybest + 1 - d
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) ParsedInput {
	bytes, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}
	for bytes[len(bytes)-1] == '\n' {
		bytes = bytes[:len(bytes)-1]
	}
	value, err := strconv.Atoi(string(bytes))
	if err != nil {
		panic(err)
	}
	return ParsedInput(value)
}

func DoPart1(input ParsedInput) Part1Result {
	partials := partials(300, int(input))
	_, xbest, ybest := bestSubSquare(3, partials)
	return Part1Result(fmt.Sprintf("%d,%d", xbest+1, ybest+1))
}

func DoPart2(input ParsedInput) Part2Result {
	partials := partials(300, int(input))
	best, xbest, ybest, dbest := math.MinInt, 0, 0, 0

	for d := 1; d < 300; d++ {
		val, x, y := bestSubSquare(d, partials)
		if val > best {
			best, xbest, ybest, dbest = val, x, y, d
		}
	}

	return Part2Result(fmt.Sprintf("%d,%d,%d", xbest+1, ybest+1, dbest))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
