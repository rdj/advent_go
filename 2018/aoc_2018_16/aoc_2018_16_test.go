package aoc_2018_16

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	egs := map[string]Part1Result{
		"Before: [3, 2, 1, 1]\n9 2 1 2\nAfter:  [3, 2, 2, 1]\n\n": 1,
	}

	for in, want := range egs {
		got := DoPart1(ParseInput(strings.NewReader(in)))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
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
