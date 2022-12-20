package aoc_2022_20

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 3700

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 10_626_948_369_382

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) []int {
	lines := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, n)
	}

	return lines
}

func dump(r *ring.Ring) {
	fmt.Printf("%v", r.Value)
	for e := r.Next(); e != r; e = e.Next() {
		fmt.Printf(" %v", e.Value)
	}
	fmt.Println()
}

func mod_euclid(a, b int) int {
	return (a%b + b) % b
}

func run(input []int, multiplier, iterations int) int {
	var zero *ring.Ring = nil

	order := make([]*ring.Ring, 0, len(input))
	r := ring.New(len(input))
	for _, n := range input {
		r.Value = n * multiplier
		order = append(order, r)
		if n == 0 {
			zero = r
		}
		r = r.Next()
	}

	for i := 0; i < iterations; i++ {
		for _, t := range order {
			//dump(r)
			n := t.Value.(int)
			if n == 0 {
				continue
			}

			after := t.Next()
			t.Prev().Unlink(1)

			n = mod_euclid(n-1, len(input)-1)

			pos := after.Move(n)
			//fmt.Println(t.Value, "moves between", pos.Value, "and", pos.Next().Value)
			pos.Link(t)
		}
		//dump(r)
	}

	x := zero.Move(1000).Value.(int)
	y := zero.Move(2000).Value.(int)
	z := zero.Move(3000).Value.(int)

	return x + y + z
}

func DoPart1(input []int) Part1Result {
	return Part1Result(run(input, 1, 1))
}

func DoPart2(input []int) Part2Result {
	return Part2Result(run(input, 811589153, 10))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
