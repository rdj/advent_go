package aoc_2018_11

import (
	"reflect"
	"strings"
	"testing"
)

func TestSingleValue(t *testing.T) {
	egs := []struct{ x, y, n, want int }{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}

	for _, eg := range egs {
		eg.x -= 1
		eg.y -= 1
		got := singleValue(eg.x, eg.y, eg.n)
		want := eg.want
		if got != want {
			t.Errorf("singleValue(%d, %d, %d) got %v, wanted %v", eg.x, eg.y, eg.n, got, want)
		}
	}
}

func TestPart1Examples(t *testing.T) {
	egs := map[string]Part1Result{
		"18": "33,45",
		"42": "21,61",
	}

	for in, want := range egs {
		got := DoPart1(ParseInput(strings.NewReader(in)))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart2Examples(t *testing.T) {
	egs := map[string]Part2Result{
		"18": "90,269,16",
		"42": "232,251,12",
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
