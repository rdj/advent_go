package aoc_2018_08

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	egs := map[string]Part1Result{
		"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2": 138,
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
		"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2": 66,
	}

	for in, want := range egs {
		got := DoPart2(ParseInput(strings.NewReader(in)))
		if got != want {
			t.Errorf("DoPart2(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart1(t *testing.T) {
	var want Part1Result = 49_180
	got := Part1()
	if got != want {
		t.Errorf("Part1() got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	var want Part2Result = 20_611
	got := Part2()
	if got != want {
		t.Errorf("Part2() got %v, wanted %v", got, want)
	}
}
