package aoc_2018_14

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	egs := map[string]Part1Result{
		"9":    "5158916779",
		"5":    "0124515891",
		"18":   "9251071085",
		"2018": "5941429882",
	}

	for in, want := range egs {
		got := DoPart1(ParseInput1(strings.NewReader(in)))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart2Examples(t *testing.T) {
	egs := map[string]Part2Result{
		"51589": 9,
		"01245": 5,
		"92510": 18,
		"59414": 2018,
	}

	for in, want := range egs {
		got := DoPart2(ParseInput2(strings.NewReader(in)))
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
