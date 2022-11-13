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
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			var c Claim
			if err := c.Scan(scanner.Text()); err != nil {
				panic(err)
			}
			ch <- c
		}
		close(ch)
	}()
	return ch
}

func DoPart1(ch chan Claim) AdventResult {
	claimed := map[Coord]int{}
	for e := range ch {
		for r := e.row; r < e.row+e.height; r++ {
			for c := e.col; c < e.col+e.width; c++ {
				claimed[Coord{r, c}] += 1
			}
		}
	}

	count := 0
	for _, n := range claimed {
		if n > 1 {
			count++
		}
	}
	return AdventResult(count)
}

func DoPart2(ch chan Claim) AdventResult {
	return AdventResult(0)
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() AdventResult {
	return DoPart2(ParseInput(openInput()))
}
