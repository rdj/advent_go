package aoc_2018_03

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	egs := map[string]AdventResult{
		"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2": 4,
	}

	for in, want := range egs {
		got := DoPart1(ParseInput(strings.NewReader(in)))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart2Examples(t *testing.T) {
	egs := map[string]AdventResult{
		"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2": 3,
	}

	for in, want := range egs {
		got := DoPart2(ParseInput(strings.NewReader(in)))
		if got != want {
			t.Errorf("DoPart2(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart1(t *testing.T) {
	const want = 103_482
	got := Part1()
	if got != want {
		t.Errorf("Part1() got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	const want = 686
	got := Part2()
	if got != want {
		t.Errorf("Part2() got %v, wanted %v", got, want)
	}
}
