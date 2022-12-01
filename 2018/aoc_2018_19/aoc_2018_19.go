package aoc_2018_19

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
const Part1Want = 1326

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 14_562_240

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

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
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

func (p *Program) run(r Registers, maxSteps int) (_ Registers, killed bool) {
	steps := 0
	ip := &r[p.ipReg]
	for *ip = 0; *ip >= 0 && *ip < len(p.instructions); *ip++ {
		if steps > maxSteps {
			return r, true
		}
		if debug {
			fmt.Printf("ip=%d %v ", *ip, r)
		}
		instr := p.instructions[*ip]
		r = *operations[instr.op](&instr, r)
		if debug {
			fmt.Printf("%v %v\n", instr, r)
		}
		steps++
	}

	return r, false
}

func DoPart1(p *Program) Part1Result {
	r, killed := p.run(Registers{}, int(^uint(0)>>1))
	if killed {
		panic("part1 killed")
	}
	return Part1Result(r[0])
}

func DoPart2(p *Program) Part2Result {
	// Reading the input program, instructions 1-16 implement a
	// quadratic loop to sum all the divisors of the number in
	// register 1. Effectively:
	//
	//	n := SomeBigNumber               // $1: n
	// 	result := 0;                     // $0: result
	// 	for x := 1; x <= n; x++ {        // $5: x
	// 		for y := 1; y <= n; y++ {    // $3: y
	//			if x * y == n {
	//				result += x
	//			}
	//		}
	//	}
	//
	// SomeBigNumber is calculated by the second half of the program,
	// starting at instruction 17.
	//
	// I don't know how the inputs vary by person, but it's possible
	// that the registers are different, etc.
	//
	// I will run the program for 1000 cycles, then capture the value
	// of r[1] to get the big number. If it moved around, it would be
	// fine to just take the biggest number in any register.
	r := Registers{}
	r[0] = 1
	r, killed := p.run(r, 1000)
	if !killed {
		panic("part2 completed")
	}

	// Now do what the elves want in linear time instead. There are
	// even better algorithms, but this is fast enough and very
	// straightforward.
	n := r[1]
	result := 0
	for x := 1; x <= n; x++ {
		if n%x == 0 {
			result += x
		}
	}
	return Part2Result(result)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
