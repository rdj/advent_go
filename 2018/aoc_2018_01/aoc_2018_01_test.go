package aoc_2018_01

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	egs := map[string][]int{
		"+1\n+1\n+1": {1, 1, 1},
		"+1\n+1\n-2": {1, 1, -2},
		"-1\n-2\n-3": {-1, -2, -3},
	}

	for in, want := range egs {
		got := ParseInput(strings.NewReader(in))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ParseInput(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart1Examples(t *testing.T) {
	egs := map[string]AdventResult{
		"+1\n-2\n+3\n+1": 3,
		"+1\n+1\n+1":     3,
		"+1\n+1\n-2":     0,
		"-1\n-2\n-3":     -6,
	}

	for in, want := range egs {
		got := DoPart1(ParseInput(strings.NewReader(in)))
		if got != want {
			t.Errorf("DoPart1(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart2Examples(t *testing.T) {
	egs := map[string]AdventResult{
		"+1\n-2\n+3\n+1":     2,
		"+1\n-1":             0,
		"+3\n+3\n+4\n-2\n-4": 10,
		"-6\n+3\n+8\n+5\n-6": 5,
		"+7\n+7\n-2\n-7\n-4": 14,
	}

	for in, want := range egs {
		got := DoPart2(ParseInput(strings.NewReader(in)))
		if got != want {
			t.Errorf("DoPart2(%q) got %v, wanted %v", in, got, want)
		}
	}
}

func TestPart1(t *testing.T) {
	const want = 502
	got := Part1()
	if got != want {
		t.Errorf("Part1() got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	const want = 71_961
	got := Part2()
	if got != want {
		t.Errorf("Part2() got %v, wanted %v", got, want)
	}
}
