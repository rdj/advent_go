package aoc_2022_21

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 78342931359552

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

type Monkey struct {
	id string
	n  int
	a  string
	b  string
	op string
}

func (m Monkey) Value(monkeys Monkeys) int {
	if m.n != NoNumber {
		return m.n
	}

	a := monkeys[m.a].Value(monkeys)
	b := monkeys[m.b].Value(monkeys)

	val := 0
	switch m.op {
	case "+":
		val = a + b

	case "-":
		val = a - b

	case "*":
		val = a * b

	case "/":
		val = a / b
	}
	m.n = val
	return val
}

func (m Monkey) Expression(monkeys Monkeys) string {
	if m.id == "humn" {
		return "(humn)"
	}
	if m.n != NoNumber {
		return fmt.Sprintf("%d", m.n)
	}

	if m.id == "root" {
		m.op = "="
	}

	a := monkeys[m.a].Expression(monkeys)
	b := monkeys[m.b].Expression(monkeys)

	if a[0] == '(' || b[0] == '(' {
		return fmt.Sprintf("(%s %s %s)", a, m.op, b)
	}

	return fmt.Sprintf("%d", m.Value(monkeys))
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

type Monkeys map[string]Monkey

const NoNumber = int(^uint(0) >> 1)

func ParseInput(input io.Reader) Monkeys {
	lines := make(Monkeys)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		m := Monkey{}

		strs := strings.Split(line, ": ")
		m.id = strs[0]

		strs = strings.Split(strs[1], " ")
		if len(strs) == 1 {
			m.n = lo.Must(strconv.Atoi(strs[0]))
		} else {
			m.n = NoNumber
			m.a = strs[0]
			m.op = strs[1]
			m.b = strs[2]
		}

		lines[m.id] = m
	}

	return lines
}

func DoPart1(monkeys Monkeys) Part1Result {
	return Part1Result(monkeys["root"].Value(monkeys))
}

func DoPart2(monkeys Monkeys) Part2Result {
	// Then I solved it by hand too long for Wolfram Alfa and I
	// couldn't immediately remember how to do it in sympy. Now that I
	// did that, I could probably write the code to generate the
	// expression pretty easily, but then I have to also evaluate the
	// expression which seems annoying. Probably you could do this all
	// with the monkeys from part one since they're an operation tree
	// anyway. The tricky part is making sure to handle e.g. (5 -
	// (humn_expr)), cause you need to subtract five and then multiply
	// by -1. All the other operations are trivially reversible.
	fmt.Println(monkeys["root"].Expression(monkeys))
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
