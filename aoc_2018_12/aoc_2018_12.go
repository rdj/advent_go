package aoc_2018_12

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 2349

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 2_100_000_001_168

const HasPlant = '#'
const LacksPlant = '.'

type Pot struct {
	id    int
	value bool
}

type Pots struct {
	state *list.List
	rules []bool
}

func (pots *Pots) String() string {
	var sb strings.Builder
	for e := pots.state.Front(); e != nil; e = e.Next() {
		if e.Value.(*Pot).value {
			sb.WriteRune(HasPlant)
		} else {
			sb.WriteRune(LacksPlant)
		}
	}
	return sb.String()
}

func (pots *Pots) advance(n int) {
	for i := 0; i < n; i++ {
		pots.resize()

		windowMask := 0b1_1111
		window := 0

		var cur, up1, up2 *list.Element
		cur = pots.state.Front()
		up1 = cur.Next()
		up2 = up1.Next()

		for cur != nil {
			window = (window << 1) & windowMask
			if up2 != nil && up2.Value.(*Pot).value {
				window |= 1
			}

			cur.Value.(*Pot).value = pots.rules[window]

			cur = up1
			up1 = up2
			if up1 != nil {
				up2 = up1.Next()
			}
		}
	}
}

// To function, we need a buffer of two empty pots at the beginning at the end
func (pots *Pots) resize() {
	// Add pots if needed
	for pots.state.Front().Value.(*Pot).value || pots.state.Front().Next().Value.(*Pot).value {
		pots.state.PushFront(&Pot{pots.state.Front().Value.(*Pot).id - 1, false})
	}
	for pots.state.Back().Value.(*Pot).value || pots.state.Back().Prev().Value.(*Pot).value {
		pots.state.PushBack(&Pot{pots.state.Back().Value.(*Pot).id + 1, false})
	}

	// This code sloppily assumes that the state list never gets very small

	back2 := pots.state.Front()
	back1 := back2.Next()
	cur := back1.Next()
	for !cur.Value.(*Pot).value {
		pots.state.Remove(back2)
		back2 = back1
		back1 = cur
		cur = cur.Next()
	}

	up2 := pots.state.Back()
	up1 := up2.Prev()
	cur = up1.Prev()
	for !cur.Value.(*Pot).value {
		pots.state.Remove(up2)
		up2 = up1
		up1 = cur
		cur = cur.Prev()
	}
}

func (pots *Pots) sum(offset int) int {
	sum := 0
	for e := pots.state.Front(); e != nil; e = e.Next() {
		pot := e.Value.(*Pot)
		if pot.value {
			sum += pot.id + offset
		}
	}
	return sum
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) Pots {
	pots := Pots{}

	scanner := bufio.NewScanner(input)

	if !scanner.Scan() {
		panic("no first line")
	}

	var initial string
	_, err := fmt.Sscanf(scanner.Text(), "initial state: %s", &initial)
	if err != nil {
		panic(err)
	}
	pots.state = list.New()

	state := bufio.NewScanner(strings.NewReader(initial))
	state.Split(bufio.ScanBytes)
	for id := 0; state.Scan(); id++ {
		pots.state.PushBack(&Pot{id, state.Bytes()[0] == HasPlant})
	}

	pots.rules = make([]bool, 32)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var inrule string
		var outrule rune
		_, err := fmt.Sscanf(scanner.Text(), "%s => %c", &inrule, &outrule)
		if err != nil {
			panic(err)
		}

		inrule = strings.Map(func(r rune) rune {
			switch r {
			case LacksPlant:
				return '0'
			case HasPlant:
				return '1'
			default:
				return -1
			}
		}, inrule)

		in, err := strconv.ParseUint(inrule, 2, 5)
		if err != nil {
			panic(err)
		}

		pots.rules[in] = outrule == HasPlant
	}
	return pots
}

func DoPart1(pots Pots) Part1Result {
	pots.advance(20)
	return Part1Result(pots.sum(0))
}

// Eventually applying the rules reaches equilibrium where the number
// and relative positions of the pots do not change, the whole thing
// just shifts with each iteration.
func DoPart2(pots Pots) Part2Result {
	front := 0
	state := pots.String()
	for i := 1; true; i++ {
		pots.advance(1)
		state0 := pots.String()
		front0 := pots.state.Front().Value.(*Pot).id
		if state == state0 {
			delta := front0 - front
			gensRemain := 50_000_000_000 - i

			return Part2Result(pots.sum(delta * gensRemain))
		}
		state = state0
		front = front0
	}

	panic("unreachable")
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
