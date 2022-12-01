package aoc_2022_01

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 74198

type Part2Result int

const Part2Want = 209914

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) [][]int {
	elves := make([][]int, 0)
	elf := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, elf)
			elf = make([]int, 0)
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		elf = append(elf, n)
	}
	if len(elf) > 0 {
		elves = append(elves, elf)
	}

	return elves
}

func sums(elves [][]int) []int {
	a := make([]int, 0, len(elves))
	for _, elf := range elves {
		n := 0
		for _, a := range elf {
			n += a
		}
		a = append(a, n)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	return a
}

func DoPart1(elves [][]int) Part1Result {
	return Part1Result(sums(elves)[0])
}

func DoPart2(elves [][]int) Part2Result {
	sums := sums(elves)
	a := 0
	for _, n := range sums[:3] {
		a += n
	}
	return Part2Result(a)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
