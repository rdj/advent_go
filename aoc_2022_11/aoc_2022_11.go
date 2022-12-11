package aoc_2022_11

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 66124

type Part2Result int

const Part2Want = 19_309_892_877

type Monkey struct {
	number                      int
	items                       []int
	op                          func(n int) int
	factor, isFactor, notFactor int
	inspections                 int
}

type Monkeys struct {
	monkeys     []*Monkey
	manageWorry func(int) int
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func square(n int) int {
	return n * n
}

func parseOp(op string) func(n int) int {
	if op == "old * old" {
		return square
	}

	var arg int
	var infix byte
	lo.Must(fmt.Sscanf(op, "old %c %d", &infix, &arg))
	switch infix {
	case '*':
		return func(n int) int {
			return arg * n
		}

	case '+':
		return func(n int) int {
			return arg + n
		}
	}

	panic("bad op")
}

func ParseInput(input io.Reader) Monkeys {
	monkeys := Monkeys{}
	monkeys.monkeys = make([]*Monkey, 0)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		m := new(Monkey)
		monkeys.monkeys = append(monkeys.monkeys, m)

		_, err := fmt.Sscanf(scanner.Text(), "Monkey %d:", &m.number)
		if err != nil {
			panic(fmt.Sprintf("%s: %s", err, scanner.Text()))
		}

		if !scanner.Scan() {
			panic("eof")
		}

		items := scanner.Text()[len("  Starting items: "):]
		m.items = lo.Map(strings.Split(items, ", "), func(s string, _ int) int { return lo.Must(strconv.Atoi(s)) })

		if !scanner.Scan() {
			panic("eof")
		}

		op := lo.Must(lo.Last(strings.Split(scanner.Text(), " = ")))
		m.op = parseOp(op)

		if !scanner.Scan() {
			panic("eof")
		}

		_, err = fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &m.factor)
		if err != nil {
			panic("bad factor")
		}

		if !scanner.Scan() {
			panic("eof")
		}

		_, err = fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &m.isFactor)
		if err != nil {
			panic("bad factor")
		}

		if !scanner.Scan() {
			panic("eof")
		}

		_, err = fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &m.notFactor)
		if err != nil {
			panic("bad factor")
		}
	}
	return monkeys
}

func (m *Monkey) Catch(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) ProcessOne(monkeys Monkeys) bool {
	if len(m.items) == 0 {
		return false
	}

	item := m.items[0]
	m.items = m.items[1:]

	item = m.op(item)
	m.inspections++

	item = monkeys.manageWorry(item)

	if item%m.factor == 0 {
		monkeys.monkeys[m.isFactor].Catch(item)
	} else {
		monkeys.monkeys[m.notFactor].Catch(item)
	}

	return true
}

func (m *Monkey) Turn(monkeys Monkeys) {
	for m.ProcessOne(monkeys) {
	}
}

func (monkeys Monkeys) Round() {
	for _, m := range monkeys.monkeys {
		m.Turn(monkeys)
	}
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %d (%d)", m.number, m.items, m.inspections)
}

func (monkeys Monkeys) Business(rounds int) int {
	for i := 0; i < rounds; i++ {
		monkeys.Round()
	}

	a := lo.Map(monkeys.monkeys, func(m *Monkey, _ int) int { return m.inspections })
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	return a[0] * a[1]
}

func DoPart1(monkeys Monkeys) Part1Result {
	monkeys.manageWorry = func(n int) int {
		return n / 3
	}
	return Part1Result(monkeys.Business(20))
}

func DoPart2(monkeys Monkeys) Part2Result {
	modulus := 1
	for _, m := range monkeys.monkeys {
		modulus *= m.factor
	}
	monkeys.manageWorry = func(n int) int {
		return n % modulus
	}

	return Part2Result(monkeys.Business(10_000))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
