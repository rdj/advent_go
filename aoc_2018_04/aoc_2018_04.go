package aoc_2018_04

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type AdventResult int

func openInput() io.Reader {
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

func ParseInput(input io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

var guardPattern *regexp.Regexp
var minutePattern *regexp.Regexp

func init() {
	guardPattern = regexp.MustCompile(`Guard #(\d+)`)
	minutePattern = regexp.MustCompile(`:(\d+)\]`)
}

func DoPart1(input []string) AdventResult {
	sort.Strings(input)

	sleepy := map[int]map[int]int{}

	guard := 0
	sleeping := -1
	for _, line := range input {
		m := minutePattern.FindStringSubmatch(line)
		if m == nil {
			panic("bad input")
		}

		minute, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}

		if m = guardPattern.FindStringSubmatch(line); m != nil {
			if guard, err = strconv.Atoi(m[1]); err != nil {
				panic(err)
			}
			if sleeping >= 0 {
				panic("guard change while sleeping")
			}
			if _, ok := sleepy[guard]; !ok {
				sleepy[guard] = make(map[int]int)
			}
			continue
		}
		if strings.Contains(line, "falls") {
			sleeping = minute
			continue
		}
		if strings.Contains(line, "wakes") {
			if sleeping > minute {
				panic("crossed day boundary")
			}
			for x := sleeping; x < minute; x++ {
				sleepy[guard][x] += 1
			}
			sleeping = -1
			continue
		}
	}

	type Guard struct{ guard, sleepTime, bestMinute, bestValue int }
	var max Guard

	for guard, mins := range sleepy {
		m := Guard{guard: guard}

		for minute, n := range mins {
			m.sleepTime += n
			if n > m.bestValue {
				m.bestMinute = minute
				m.bestValue = n
			}
		}
		if m.sleepTime > max.sleepTime {
			max = m
		}
	}

	return AdventResult(max.guard * max.bestMinute)
}

func DoPart2(input []string) AdventResult {
	return AdventResult(0)
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() AdventResult {
	return DoPart2(ParseInput(openInput()))
}
