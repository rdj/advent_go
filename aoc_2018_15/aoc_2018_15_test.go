package aoc_2018_15

import (
	"strings"
	"testing"
)

const Example1 = `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

func TestPart1Example1(t *testing.T) {
	const want = 27730
	got := DoPart1(ParseInput(strings.NewReader(Example1)))
	if got != want {
		t.Errorf("DoPart1(Example1) got %v, wanted %v", got, want)
	}
}

const Example2 = `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`

func TestPart1Example2(t *testing.T) {
	const want = 36334
	got := DoPart1(ParseInput(strings.NewReader(Example2)))
	if got != want {
		t.Errorf("DoPart1(Example2) got %v, wanted %v", got, want)
	}
}

const Example3 = `#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`

func TestPart1Example3(t *testing.T) {
	const want = 39514
	got := DoPart1(ParseInput(strings.NewReader(Example3)))
	if got != want {
		t.Errorf("DoPart1(Example3) got %v, wanted %v", got, want)
	}
}

const Example4 = `#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`

func TestPart1Example4(t *testing.T) {
	const want = 27755
	got := DoPart1(ParseInput(strings.NewReader(Example4)))
	if got != want {
		t.Errorf("DoPart1(Example4) got %v, wanted %v", got, want)
	}
}

const Example5 = `#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`

func TestPart1Example5(t *testing.T) {
	const want = 28944
	got := DoPart1(ParseInput(strings.NewReader(Example5)))
	if got != want {
		t.Errorf("DoPart1(Example5) got %v, wanted %v", got, want)
	}
}

const Example6 = `#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`

func TestPart1Example6(t *testing.T) {
	const want = 18740
	got := DoPart1(ParseInput(strings.NewReader(Example6)))
	if got != want {
		t.Errorf("DoPart1(Example6) got %v, wanted %v", got, want)
	}
}

func TestPart2Examples(t *testing.T) {
	const want = Part2Want
	got := DoPart2(ParseInput(strings.NewReader(Example1)))
	if got != want {
		t.Errorf("DoPart2(Example1) got %v, wanted %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	const want = Part1Want
	got := Part1()
	if got != want {
		t.Errorf("Part1() got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	const want = Part2Want
	got := Part2()
	if got != want {
		t.Errorf("Part2() got %v, wanted %v", got, want)
	}
}
