package aoc_2018_18

import (
	"strings"
	"testing"
)

const (
	Example1 = `.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`
)

func TestPart1Example1(t *testing.T) {
	const in = Example1
	const want = 1147
	got := DoPart1(ParseInput(strings.NewReader(in)))
	if got != want {
		t.Errorf("DoPart1(Example1) got %v, wanted %v", got, want)
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
