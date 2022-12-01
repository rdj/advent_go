package aoc_2018_06

import (
	"reflect"
	"strings"
	"testing"
)

const example1 = "1, 1\n1, 6\n8, 3\n3, 4\n5, 5\n8, 9"

func TestPart1Examples(t *testing.T) {
	egs := map[string]AdventResult{
		example1: 17,
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
		example1: 16,
	}

	for in, want := range egs {
		got := DoPart2(ParseInput(strings.NewReader(in)), 32)
		if got != want {
			t.Errorf("DoPart2(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart1(t *testing.T) {
	const want = 3722
	got := Part1()
	if got != want {
		t.Errorf("Part1() got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	const want = 44_634
	got := Part2()
	if got != want {
		t.Errorf("Part2() got %v, wanted %v", got, want)
	}
}
