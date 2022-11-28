package aoc_2018_23

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 580

type Part2Result int

const Part2Want = 97816347

type Point3D struct{ x, y, z int64 }

func absdiff(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}

func (a *Point3D) Manhattan(b *Point3D) int64 {
	return absdiff(a.x, b.x) + absdiff(a.y, b.y) + absdiff(a.z, b.z)
}

type Nanobot struct {
	pos    Point3D
	radius int64
}

func (b *Nanobot) minDistanceToOrigin() int64 {
	d, r := b.pos.Manhattan(&Point3D{}), b.radius
	if r < d {
		return d - r
	}
	return 0
}

func (b *Nanobot) maxDistanceToOrigin() int64 {
	return b.pos.Manhattan(&Point3D{}) + b.radius
}

func (a *Nanobot) rangeIncludes(b *Nanobot) bool {
	return a.pos.Manhattan(&b.pos) <= a.radius
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) []*Nanobot {
	bots := make([]*Nanobot, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		bot := new(Nanobot)
		fmt.Sscanf(scanner.Text(), "pos=<%d,%d,%d>, r=%d", &bot.pos.x, &bot.pos.y, &bot.pos.z, &bot.radius)
		bots = append(bots, bot)
	}
	return bots
}

func DoPart1(bots []*Nanobot) Part1Result {
	strongest := bots[0]
	for _, b := range bots {
		if b.radius > strongest.radius {
			strongest = b
		}
	}

	inRange := 0
	for _, b := range bots {
		if strongest.rangeIncludes(b) {
			inRange++
		}
	}

	return Part1Result(inRange)
}

type Scope struct {
	d     int64
	delta int
}

type Scopes []Scope

func (s Scopes) Len() int           { return len(s) }
func (s Scopes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Scopes) Less(i, j int) bool { return s[i].d < s[j].d }

// I feel pretty skeptical that this solution works in general for all
// possible inputs, but it worked for the example and for my input.
func DoPart2(bots []*Nanobot) Part2Result {
	const entering = 1
	const leaving = -1

	scopes := make(Scopes, 0, 2*len(bots))
	for _, b := range bots {
		scopes = append(scopes,
			Scope{b.minDistanceToOrigin(), entering},
			Scope{b.maxDistanceToOrigin() + 1, leaving})
	}

	sort.Sort(scopes)

	cum := 0
	max := 0
	dist := int64(0)
	for _, s := range scopes {
		cum += s.delta
		if cum > max {
			max = cum
			dist = s.d
		}
	}

	return Part2Result(dist)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
