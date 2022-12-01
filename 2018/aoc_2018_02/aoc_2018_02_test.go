package aoc_2018_02

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	egs := map[string]AdventResult{
		"abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab\n": 12,
	}

	for in, want := range egs {
		got := DoPart1(ParseInput(strings.NewReader(in)))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart2Examples(t *testing.T) {
	egs := map[string]string{
		"abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz": "fgij",
	}

	for in, want := range egs {
		got := DoPart2(ParseInput(strings.NewReader(in)))
		if got != want {
			t.Errorf("DoPart2(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart1(t *testing.T) {
	const want = 6175
	got := Part1()
	if got != want {
		t.Errorf("Part1() got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	const want = "asgwjcmzredihqoutcylvzinx"
	got := Part2()
	if got != want {
		t.Errorf("Part2() got %q, wanted %q", got, want)
	}
}
