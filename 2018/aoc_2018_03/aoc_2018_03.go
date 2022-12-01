package aoc_2018_03

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

type Coord struct {
	row, col int
}

type Claim struct {
	id, width, height int
	Coord
}

func (c *Claim) Scan(s string) error {
	_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &c.id, &c.col, &c.row, &c.width, &c.height)
	return err
}

func openInput() io.Reader {
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

func ParseInput(input io.Reader) chan Claim {
	ch := make(chan Claim)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			var c Claim
			if err := c.Scan(scanner.Text()); err != nil {
				panic(err)
			}
			ch <- c
		}
	}()
	return ch
}

func buildMap(ch chan Claim) map[Coord]int {
	claimed := map[Coord]int{}
	for e := range ch {
		for r := e.row; r < e.row+e.height; r++ {
			for c := e.col; c < e.col+e.width; c++ {
				claimed[Coord{r, c}] += 1
			}
		}
	}
	return claimed
}

func DoPart1(ch chan Claim) AdventResult {
	claimed := buildMap(ch)

	count := 0
	for _, n := range claimed {
		if n > 1 {
			count++
		}
	}
	return AdventResult(count)
}

func DoPart2(ch chan Claim) AdventResult {
	// Capture the Claims from ch before passing them to buildMap
	claims := make([]Claim, 0, 16)
	build := make(chan Claim)
	go func() {
		defer close(build)
		for c := range ch {
			claims = append(claims, c)
			build <- c
		}
	}()
	claimed := buildMap(build)

Entry:
	for _, e := range claims {
		for r := e.row; r < e.row+e.height; r++ {
			for c := e.col; c < e.col+e.width; c++ {
				if claimed[Coord{r, c}] > 1 {
					continue Entry
				}
			}
		}
		return AdventResult(e.id)
	}

	return AdventResult(0)
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() AdventResult {
	return DoPart2(ParseInput(openInput()))
}
