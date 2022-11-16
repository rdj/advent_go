package aoc_2018_09

import (
	"container/ring"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 412_959

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 3_333_662_986

type ParsedInput struct{ nplayers, lastMarble int }

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

	parsed := ParsedInput{}
	_, err = fmt.Sscanf(string(bytes), "%d players; last marble is worth %d points", &parsed.nplayers, &parsed.lastMarble)
	if err != nil {
		panic(err)
	}
	return parsed
}

func DoPart1(input ParsedInput) Part1Result {
	players := make([]int, input.nplayers)

	current := ring.New(1)
	current.Value = 0

	for m := 1; m <= input.lastMarble; m++ {
		if m%23 == 0 {
			p := (m - 1) % len(players)
			players[p] += m

			for i := 0; i < 8; i++ {
				current = current.Prev()
			}
			rem := current.Unlink(1)
			players[p] += rem.Value.(int)

			current = current.Next()
			continue
		}

		marb := ring.New(1)
		marb.Value = m

		skip := current.Next()
		skip.Link(marb)

		current = marb
	}

	best := 0
	for p := 0; p < len(players); p++ {
		if players[p] > best {
			best = players[p]
		}
	}

	return Part1Result(best)
}

func DoPart2(input ParsedInput) Part2Result {
	// LOL I thought I was going to have to analyze this into some
	// kind of large number modular arithmetic problem but it's no
	// problem to simulate with 64-bit integers
	input.lastMarble *= 100
	return Part2Result(DoPart1(input))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
