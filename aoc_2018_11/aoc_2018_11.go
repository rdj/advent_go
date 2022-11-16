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

func subValue(x, y, sub int, singles [][]int) int {
	value := 0
	for y0 := y; y0 < y+sub; y0++ {
		for x0 := x; x0 < x+sub; x0++ {
			value += singles[y0][x0]
		}
	}
	return value
}

func subValues(sub int, singles [][]int) [][]int {
	subDim := len(singles) + 1 - sub
	subs := make([][]int, subDim)
	for i := range subs {
		subs[i] = make([]int, subDim)
	}

	// for _, row := range singles {
	// 	for _, val := range row {
	// 		fmt.Printf("% d ", val)
	// 	}
	// 	fmt.Printf("\n")
	// }

	for y := range subs {
		for x := range subs[y] {
			subs[y][x] = subValue(x, y, sub, singles)
		}
	}

	return subs
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

func bestSubSquare(d int, singles [][]int) (best, xbest, ybest int) {
	values := subValues(d, singles)
	best, xbest, ybest = math.MinInt, math.MinInt, math.MinInt

	for y, row := range values {
		for x, val := range row {
			if val > best {
				xbest, ybest = x, y
				best = val
			}
		}
	}

	return
}

func DoPart1(input ParsedInput) Part1Result {
	singles := singleValues(300, int(input))
	_, xbest, ybest := bestSubSquare(3, singles)
	return Part1Result(fmt.Sprintf("%d,%d", xbest+1, ybest+1))
}

func DoPart2(input ParsedInput) Part2Result {
	singles := singleValues(300, int(input))
	best, xbest, ybest, dbest := math.MinInt, 0, 0, 0

	for d := 1; d < 300; d++ {
		val, x, y := bestSubSquare(d, singles)
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
