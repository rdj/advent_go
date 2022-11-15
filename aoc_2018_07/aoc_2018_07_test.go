package aoc_2018_07

import (
	"strings"
	"testing"
)

const example1 = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

func TestPart1Examples(t *testing.T) {
	const want = "CABDFE"

	got := DoPart1(ParseInput(strings.NewReader(example1)))
	if got != want {
		t.Errorf("DoPart1(example1) got %v, wanted %v", got, want)
	}
}

func TestPart2Examples(t *testing.T) {
	const want = 15

	got := DoPart2(ParseInput(strings.NewReader(example1)), 2, 1)
	if got != want {
		t.Errorf("DoPart2(example1) got %v, wanted %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	const want = "BFKEGNOVATIHXYZRMCJDLSUPWQ"
	got := Part1()
	if got != want {
		t.Errorf("Part1() got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	const want = 1020
	got := Part2()
	if got != want {
		t.Errorf("Part2() got %v, wanted %v", got, want)
	}
}
