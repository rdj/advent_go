package aoc_2018_23

import (
	"strings"
	"testing"
)

const (
	Example1 = `pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1`

	Example2 = `pos=<10,12,12>, r=2
pos=<12,14,12>, r=2
pos=<16,12,12>, r=4
pos=<14,14,14>, r=6
pos=<50,50,50>, r=200
pos=<10,10,10>, r=5`
)

func TestPart1Examples(t *testing.T) {
	const in = Example1
	const want = 7
	got := DoPart1(ParseInput(strings.NewReader(in)))
	if got != want {
		t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
	}
}

func TestPart2Examples(t *testing.T) {
	const in = Example2
	const want = 36
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
