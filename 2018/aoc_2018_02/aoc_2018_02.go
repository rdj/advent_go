package aoc_2018_02

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var _ = fmt.Println
var _ = strconv.Atoi

const inputFile = "input.txt"

type AdventResult int

func openInput() io.Reader {
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

func ParseInput(input io.Reader) io.Reader {
	return input
}

func DoPart1(input io.Reader) AdventResult {
	scanner := bufio.NewScanner(input)
	twos := 0
	threes := 0
	for scanner.Scan() {
		counts := make(map[byte]int)
		for _, b := range scanner.Bytes() {
			counts[b] += 1
		}
		two := false
		three := false
		for _, n := range counts {
			switch n {
			case 2:
				two = true
			case 3:
				three = true
			}
		}
		if two {
			twos += 1
		}
		if three {
			threes += 1
		}
	}
	return AdventResult(twos * threes)
}

func DoPart2(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	words := make([][]byte, 0, 20)
	for scanner.Scan() {
		word := append([]byte{}, scanner.Bytes()...)
		words = append(words, word)
	}

	for i, a := range words {
	Word:
		for _, b := range words[i+1:] {
			diff := -1
			for x := 0; x < len(a); x++ {
				if a[x] != b[x] {
					if diff == -1 {
						diff = x
					} else {
						continue Word
					}
				}
			}
			if diff != -1 {
				new := make([]byte, 0, len(a)-1)
				new = append(new, a[0:diff]...)
				new = append(new, a[diff+1:]...)
				return string(new)
			}
		}
	}

	panic("criteria not met")
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() string {
	return DoPart2(ParseInput(openInput()))
}
