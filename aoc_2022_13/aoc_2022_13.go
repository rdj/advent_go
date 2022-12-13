package aoc_2022_13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 5529

type Part2Result int

const Part2Want = 27690

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) NumLists {
	numlists := make(NumLists, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		if len(line) == 0 {
			continue
		}

		var nl interface{}
		err := json.Unmarshal(line, &nl)
		if nil != err {
			panic(fmt.Sprintf("JSON error processing %q: %v", line, err))
		}
		numlists = append(numlists, nl)
	}

	return numlists
}

const (
	InOrder = iota
	OutOfOrder
	Inconclusive
)

func order(a, b interface{}) int {
	switch aa := a.(type) {
	case float64:
		switch bb := b.(type) {
		case float64:
			//fmt.Println("Compare", aa, "vs", bb)
			switch {
			case aa < bb:
				//fmt.Println("Left side is smaller, so inputs are in the right order")
				return InOrder

			case bb < aa:
				//fmt.Println("Right side is smaller, so inputs are not in the right order")
				return OutOfOrder

			default:
				return Inconclusive
			}

		case []interface{}:
			//fmt.Println("Mismatch, listifying a:", aa, "vs", bb)
			return order([]interface{}{a}, b)
		}

	case []interface{}:
		switch bb := b.(type) {
		case float64:
			//fmt.Println("Mismatch, listifying b:", aa, "vs", bb)
			return order(a, []interface{}{b})

		case []interface{}:
			for i := range aa {
				if i >= len(bb) {
					//fmt.Println("Right side ran out of items, so inputs are not in the right order")
					return OutOfOrder
				}

				r := order(aa[i], bb[i])
				switch r {
				case Inconclusive:
					continue

				default:
					return r
				}
			}
			if len(aa) == len(bb) {
				return Inconclusive
			}
			//fmt.Println("Left side ran out of items, so inputs are in the right order")
			return InOrder
		}

	default:
		fmt.Printf("%T\n", a)
		panic("bad json type")
	}

	panic("they're equal")
}

func DoPart1(in NumLists) Part1Result {
	sum := 0

	for i := 0; i < len(in); i += 2 {
		// fmt.Println("== Pair", i/2+1, " ==")
		// fmt.Println("- Compare", in[i].nl, "vs", in[i+1])
		r := order(in[i], in[i+1])
		if r == InOrder {
			sum += i/2 + 1
		}
	}

	return Part1Result(sum)
}

type NumLists []interface{}

func (a NumLists) Len() int           { return len(a) }
func (a NumLists) Less(i, j int) bool { return InOrder == order(a[i], a[j]) }
func (a NumLists) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func DoPart2(in NumLists) Part2Result {
	in = append(in, []interface{}{[]interface{}{2.0}})
	in = append(in, []interface{}{[]interface{}{6.0}})
	sort.Sort(in)

	two := 0
	for i, n := range in {
		if string(lo.Must(json.Marshal(n))) == "[[2]]" {
			two = i + 1
		}
		if string(lo.Must(json.Marshal(n))) == "[[6]]" {
			return Part2Result(two * (i + 1))
		}
	}
	panic("did not find")
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
