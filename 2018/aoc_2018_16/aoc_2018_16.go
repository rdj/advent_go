package aoc_2018_16

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
const Part1Want = 563

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 629

type ParsedInput struct {
	samples []*Sample
	program []Instruction
}

type Instruction struct {
	opcode, a, b, c int
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

type Registers [4]int

func (a *Registers) Eq(b *Registers) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3]
}

type Sample struct {
	before, after Registers
	instr         Instruction
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(reader io.Reader) ParsedInput {
	input := ParsedInput{[]*Sample{}, []Instruction{}}
	samples := &input.samples
	sample := &Sample{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}

		_, err := fmt.Sscanf(scanner.Text(),
			"Before: [%d, %d, %d, %d]",
			&sample.before[0], &sample.before[1], &sample.before[2], &sample.before[3])
		if err != nil {
			panic(err)
		}

		if !scanner.Scan() {
			panic("truncated input")
		}
		_, err = fmt.Sscanf(scanner.Text(),
			"%d %d %d %d",
			&sample.instr.opcode, &sample.instr.a, &sample.instr.b, &sample.instr.c)
		if err != nil {
			panic(err)
		}

		if !scanner.Scan() {
			panic("truncated input")
		}
		_, err = fmt.Sscanf(scanner.Text(),
			"After: [%d, %d, %d, %d]",
			&sample.after[0], &sample.after[1], &sample.after[2], &sample.after[3])
		if err != nil {
			panic(err)
		}

		*samples = append(*samples, sample)
		sample = &Sample{}

		if !scanner.Scan() {
			panic("truncated input")
		}
		if len(scanner.Text()) > 0 {
			panic("malformed input")
		}
	}

	scanner.Scan()
	for scanner.Scan() {
		instr := Instruction{}
		_, err := fmt.Sscanf(scanner.Text(),
			"%d %d %d %d",
			&instr.opcode, &instr.a, &instr.b, &instr.c)
		if err != nil {
			panic(err)
		}
		input.program = append(input.program, instr)
	}

	return input
}

func DoPart1(input ParsedInput) Part1Result {
	cum := 0
	for _, sample := range input.samples {
		n := 0
		for _, op := range operations {
			if op(&sample.instr, sample.before).Eq(&sample.after) {
				n++
			}
		}
		if n >= 3 {
			cum++
		}
	}
	return Part1Result(cum)
}

type opmap []map[string]bool

func (m opmap) eliminate(opcode int, op string) {
	delete(m[opcode], op)
	if len(m[opcode]) == 1 {
		var found string
		for n := range m[opcode] {
			found = n
			break
		}
		for i := range m {
			if i != opcode && m[i][found] {
				m.eliminate(i, found)
			}
		}
	}
}

func (m opmap) opcodeMap() map[int]string {
	opcodes := map[int]string{}
opcode:
	for opcode, mapping := range m {
		if len(mapping) != 1 {
			panic("deduction failure")
		}
		for opname := range mapping {
			opcodes[opcode] = opname
			continue opcode
		}
	}
	return opcodes
}

func deduceOpcodes(input ParsedInput) map[int]string {
	ops := make(opmap, 16)
	for n := range ops {
		ops[n] = make(map[string]bool, len(operations))
		for op := range operations {
			ops[n][op] = true
		}
	}

	for _, sample := range input.samples {
		opcode := sample.instr.opcode
		for op := range ops[opcode] {
			if !operations[op](&sample.instr, sample.before).Eq(&sample.after) {
				ops.eliminate(opcode, op)
			}
		}
	}

	return ops.opcodeMap()
}

func DoPart2(input ParsedInput) Part2Result {
	ops := deduceOpcodes(input)

	r := Registers{}
	for _, instr := range input.program {
		r = *operations[ops[instr.opcode]](&instr, r)
	}

	return Part2Result(r[0])
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
