package aoc_2018_12

import (
	"strings"
	"testing"
)

const Example1 = `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`

func TestPart1Examples(t *testing.T) {
	const want = 325
	got := DoPart1(ParseInput(strings.NewReader(Example1)))
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
