package aoc_2018_09

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	egs := map[string]Part1Result{
		"9 players; last marble is worth 25 points":    32,
		"10 players; last marble is worth 1618 points": 8317,
		"13 players; last marble is worth 7999 points": 146373,
		"17 players; last marble is worth 1104 points": 2764,
		"21 players; last marble is worth 6111 points": 54718,
		"30 players; last marble is worth 5807 points": 37305,
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
