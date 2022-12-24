package aoc_2022_24

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"math/bits"
	"os"
	"strings"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

// mod_euclid(a, b) returns the remainder of the "euclidian division"
// a / b. This result always has the same sign as b. This is how the %
// operator works in perl, ruby, and python. (In go, as in C, a % b
// matches the sign of a.)
//
// See also: https://torstencurdt.com/tech/posts/modulo-of-negative-numbers/
func mod_euclid(a, b int) int {
	return (a%b + b) % b
}

func lcm(a, b int) int {
	a_ := big.NewInt(int64(a))
	b_ := big.NewInt(int64(b))
	var gcd_ big.Int
	gcd_.GCD(nil, nil, a_, b_)
	gcd := int(gcd_.Uint64())
	return (a * b) / gcd
}

type Point struct{ x, y int }

func (a Point) Plus(b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func (a Point) Mod(b Point) Point {
	return Point{mod_euclid(a.x, b.x), mod_euclid(a.y, b.y)}
}

func (p Point) North() Point {
	return p.Plus(Point{0, -1})
}

func (p Point) East() Point {
	return p.Plus(Point{1, 0})
}

func (p Point) South() Point {
	return p.Plus(Point{0, 1})
}

func (p Point) West() Point {
	return p.Plus(Point{-1, 0})
}

func (p Point) OffsetByFlag(d byte) Point {
	switch d {
	case NorthBliz:
		return p.North()
	case EastBliz:
		return p.East()
	case SouthBliz:
		return p.South()
	case WestBliz:
		return p.West()
	default:
		panic("bad direction flag")
	}
}

const (
	EmptyTile = byte(0)
	NorthBliz = byte(0b00001)
	EastBliz  = byte(0b00010)
	SouthBliz = byte(0b00100)
	WestBliz  = byte(0b01000)
	WallTile  = byte(0b10000)
)

type Blizzards struct {
	rows, cols              int
	state                   [][]byte
	timeExplored, timeCycle int
}

func NewBlizzard(initial []byte, rows, cols int) *Blizzards {
	b := new(Blizzards)
	b.rows, b.cols = rows, cols
	// Note: If blizzards can enter the start/end tiles, the cycle is
	// much more complicated to find.
	b.timeCycle = lcm(rows-2, cols-2)
	if len(initial) != rows*cols {
		panic("length mismatch")
	}
	b.state = make([][]byte, b.timeCycle)
	b.state[0] = make([]byte, b.rows*b.cols)
	copy(b.state[0], initial)
	return b
}

func (b *Blizzards) advance(t0 int) {
	if t0 >= b.timeCycle {
		panic("bad t0")
	}
	for t := b.timeExplored; t < t0; t++ {
		wrap := Point{b.cols, b.rows}

		cur := b.state[t]
		next := make([]byte, b.rows*b.cols)
		b.state[t+1] = next

		for i, v := range cur {
			if v == WallTile {
				next[i] = WallTile
				continue
			}
			p := b.point(i)
			for d := 0; d < 4; d++ {
				flag := v & (1 << d)
				if flag != 0 {
					n := p
					ni := 0
					wall := WallTile
					for wall == WallTile {
						n = n.OffsetByFlag(flag).Mod(wrap)
						ni = b.index(n)
						// Don't check for walls in next since they're not copied ahead of time
						wall = cur[ni]
					}
					if n.y == 0 || n.y == b.rows-1 {
						panic("blizzard entered start/goal tile")
					}
					next[ni] |= flag
				}
			}
		}
	}
	b.timeExplored = t0
}

func (b *Blizzards) get(t int, p Point) byte {
	s := b.stateAt(t)
	i := b.index(p)
	return s[i]
}

func (b *Blizzards) point(i int) Point {
	return Point{i % b.cols, i / b.cols}
}

func (b *Blizzards) index(p Point) int {
	return p.y*b.cols + p.x
}

func (b *Blizzards) stateAt(t int) []byte {
	t = t % b.timeCycle
	b.advance(t)
	return b.state[t]
}

func (b *Blizzards) String(t int) string {
	s := b.stateAt(t)
	sb := new(strings.Builder)
	for i, v := range s {
		if i%b.cols == 0 && i > 0 {
			sb.WriteByte('\n')
		}

		switch v {
		case EmptyTile:
			sb.WriteByte('.')
		case WallTile:
			sb.WriteByte('#')
		case NorthBliz:
			sb.WriteByte('^')
		case EastBliz:
			sb.WriteByte('>')
		case SouthBliz:
			sb.WriteByte('v')
		case WestBliz:
			sb.WriteByte('<')
		default:
			sb.WriteString(fmt.Sprintf("%d", bits.OnesCount8(v)))
		}
	}

	return sb.String()
}

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 0xBAAD_F00D

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) *Blizzards {
	initial := []byte{}
	rows, cols := 0, 0

	scanner := bufio.NewScanner(input)
	for ; scanner.Scan(); rows++ {
		line := scanner.Text()
		cols = len(line)

		for _, r := range line {
			b := EmptyTile
			switch r {
			case '^':
				b = NorthBliz
			case '>':
				b = EastBliz
			case 'v':
				b = SouthBliz
			case '<':
				b = WestBliz
			case '#':
				b = WallTile
			case '.':
				b = EmptyTile
			default:
				panic(fmt.Sprintf("bad tile: %q", r))
			}
			initial = append(initial, b)
		}
	}

	return NewBlizzard(initial, rows, cols)
}

func DoPart1(b *Blizzards) Part1Result {
	return Part1Fake
}

func DoPart2(b *Blizzards) Part2Result {
	return Part2Fake
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
