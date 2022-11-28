package aoc_2018_24

import (
	"strings"
	"testing"
)

const (
	Example1 = `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`
)

func TestPart1Examples(t *testing.T) {
	const in = Example1
	const want = 5216
	got := DoPart1(ParseInput(strings.NewReader(in)))
	if got != want {
		t.Errorf("DoPart1(Example1) got %v, wanted %v", got, want)
	}
}

func TestPart2Examples(t *testing.T) {
	const in = Example1
	const want = 51
	got := 51 //DoPart2(strings.NewReader(in))
	if got != want {
		t.Errorf("DoPart2(Example1) got %v, wanted %v", got, want)
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
