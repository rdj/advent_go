package aoc_2018_13

import (
	"strings"
	"testing"
)

const Example1 = `/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

const Example2 = `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`

func TestPart1Examples(t *testing.T) {
	const want = "7,3"
	got := DoPart1(ParseInput(strings.NewReader(Example1)))
	if got != want {
		t.Errorf("DoPart1(Example1) got %v, wanted %v", got, want)
	}
}

func TestPart2Examples(t *testing.T) {
	const want = "6,4"
	got := DoPart2(ParseInput(strings.NewReader(Example2)))
	if got != want {
		t.Errorf("DoPart2(Example2) got %v, wanted %v", got, want)
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
