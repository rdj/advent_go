package aoc_2018_17

import (
	"strings"
	"testing"
)

const (
	Example1 = `x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504`
)

func TestPart1Example1(t *testing.T) {
	const in = Example1
	const want = 57
	got := DoPart1(ParseInput(strings.NewReader(in)))
	if got != want {
		t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
	}
}

func TestPart2Example1(t *testing.T) {
	const in = Example1
	const want = 29
	got := DoPart2(ParseInput(strings.NewReader(in)))
	if got != want {
		t.Errorf("DoPart2(%q) got %v, wanted %v", in, got, want)
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
