package aoc_2022_24

import (
	"strings"
	"testing"
)

const (
	Example1 = `#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#`

//	Example1 = `#.######
//
// #>>.<^<#
// #.<..<<#
// #>v.><>#
// #<^v^^>#
// ######.#`
)

func TestPart1Examples(t *testing.T) {
	egs := map[string]Part1Result{
		Example1: 18,
	}

	for in, want := range egs {
		got := DoPart1(ParseInput(strings.NewReader(in)))
		if got != want {
			t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart2Examples(t *testing.T) {
	egs := map[string]Part2Result{
		Example1: Part2Want,
	}

	for in, want := range egs {
		got := DoPart2(ParseInput(strings.NewReader(in)))
		if got != want {
			t.Errorf("DoPart2(%q) got %v, wanted %v", in, got, want)
		}
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
