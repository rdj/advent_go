package aoc_2018_20

import (
	"strings"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	egs := map[string]Part1Result{
		"^WNE$":                   3,
		"^ENWWW(NEEE|SSE(EE|N))$": 10,
		"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$":                         18,
		"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$":               23,
		"^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$": 31,
	}

	for in, want := range egs {
		got := DoPart1(ParseInput(strings.NewReader(in)))
		if got != want {
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
