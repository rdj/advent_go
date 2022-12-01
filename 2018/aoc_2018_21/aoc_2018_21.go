package aoc_2018_21

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 3_941_014

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 13_775_890

const registerCount = 6

type Instruction struct {
	op      string
	a, b, c int
}

type Operation func(i *Instruction, r Registers) *Registers
type Operations map[string]Operation

var operations Operations

func init() {
	operations = Operations{
		"addr": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a] + r[i.b]
			return &r
		},
		"addi": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a] + i.b
			return &r
		},
		"mulr": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a] * r[i.b]
			return &r
		},
		"muli": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a] * i.b
			return &r
		},
		"banr": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a] & r[i.b]
			return &r
		},
		"bani": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a] & i.b
			return &r
		},
		"borr": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a] | r[i.b]
			return &r
		},
		"bori": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a] | i.b
			return &r
		},
		"setr": func(i *Instruction, r Registers) *Registers {
			r[i.c] = r[i.a]
			return &r
		},
		"seti": func(i *Instruction, r Registers) *Registers {
			r[i.c] = i.a
			return &r
		},
		"gtir": func(i *Instruction, r Registers) *Registers {
			if i.a > r[i.b] {
				r[i.c] = 1
			} else {
				r[i.c] = 0
			}
			return &r
		},
		"gtri": func(i *Instruction, r Registers) *Registers {
			if r[i.a] > i.b {
				r[i.c] = 1
			} else {
				r[i.c] = 0
			}
			return &r
		},
		"gtrr": func(i *Instruction, r Registers) *Registers {
			if r[i.a] > r[i.b] {
				r[i.c] = 1
			} else {
				r[i.c] = 0
			}
			return &r
		},
		"eqir": func(i *Instruction, r Registers) *Registers {
			if i.a == r[i.b] {
				r[i.c] = 1
			} else {
				r[i.c] = 0
			}
			return &r
		},
		"eqri": func(i *Instruction, r Registers) *Registers {
			if r[i.a] == i.b {
				r[i.c] = 1
			} else {
				r[i.c] = 0
			}
			return &r
		},
		"eqrr": func(i *Instruction, r Registers) *Registers {
			if r[i.a] == r[i.b] {
				r[i.c] = 1
			} else {
				r[i.c] = 0
			}
			return &r
		},
	}
}

type Registers [registerCount]int

type Program struct {
	ipReg        int
	instructions []Instruction
}

func (a *Registers) Eq(b *Registers) bool {
	for i := 0; i < registerCount; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ParseInput(input io.Reader) *Program {
	program := new(Program)

	scanner := bufio.NewScanner(input)

	if !scanner.Scan() {
		panic("bad input")
	}

	_, err := fmt.Sscanf(scanner.Text(), "#ip %d", &program.ipReg)
	if err != nil {
		panic("bad input")
	}

	for scanner.Scan() {
		instr := Instruction{}
		_, err = fmt.Sscanf(scanner.Text(), "%s %d %d %d", &instr.op, &instr.a, &instr.b, &instr.c)
		if err != nil {
			panic("bad input")
		}
		program.instructions = append(program.instructions, instr)
	}
	return program
}

var debug bool = false

func (p *Program) run(r Registers, toInstr int, ch chan Registers) Registers {
	ip := &r[p.ipReg]
	for *ip = 0; *ip >= 0 && *ip < len(p.instructions); *ip++ {
		if *ip == toInstr {
			ch <- r
		}
		if debug {
			fmt.Printf("ip=%d %v ", *ip, r)
		}
		instr := p.instructions[*ip]
		r = *operations[instr.op](&instr, r)
		if debug {
			fmt.Printf("%v %v\n", instr, r)
		}
	}

	return r
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func DoPart1(prog *Program) Part1Result {
	ch := make(chan Registers)
	go prog.run(Registers{}, 28, ch)
	r := <-ch
	return Part1Result(r[5])
}

// This takes forever but I'm not like having very much fun so I'd say
// good enough for now.
//
// Checked reddit and nobody really had a better approach. Some people
// transpiled the elves' code to something else so they could run
// part2 faster, but stil just looked for a cycle as here.
func DoPart2(prog *Program) Part2Result {
	ch := make(chan Registers)
	go prog.run(Registers{}, 28, ch)

	seen := map[int]bool{}
	last := 0

	for r := range ch {
		if seen[r[5]] {
			return Part2Result(last)
		}
		last = r[5]
		seen[last] = true
	}
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
