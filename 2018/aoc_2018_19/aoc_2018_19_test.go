package aoc_2018_19

import (
	"strings"
	"testing"
)

const (
	Example1 = `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`
)

func TestPart1Examples(t *testing.T) {
	const in = Example1
	const want = 7
	got := DoPart1(ParseInput(strings.NewReader(in)))
	if got != want {
		t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
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
