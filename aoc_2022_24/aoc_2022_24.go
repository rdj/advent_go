package aoc_2022_24

import (
	"bufio"
	"container/heap"
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

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 242

type Part2Result int

const Part2Want = 720

func absdiff(n, m int) int {
	if n > m {
		return n - m
	}
	return m - n
}

func lcm(a, b int) int {
	a_ := big.NewInt(int64(a))
	b_ := big.NewInt(int64(b))
	var gcd_ big.Int
	gcd_.GCD(nil, nil, a_, b_)
	gcd := int(gcd_.Uint64())
	return (a * b) / gcd
}

// mod_euclid(a, b) returns the remainder of the "euclidian division"
// a / b. This result always has the same sign as b. This is how the %
// operator works in perl, ruby, and python. (In go, as in C, a % b
// matches the sign of a.)
//
// See also: https://torstencurdt.com/tech/posts/modulo-of-negative-numbers/
func mod_euclid(a, b int) int {
	return (a%b + b) % b
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

func (a Point) Manhattan(b Point) int {
	return absdiff(a.x, b.x) + absdiff(a.y, b.y)
}

func (p Point) Neighbors() []Point {
	return []Point{
		p.North(),
		p.East(),
		p.South(),
		p.West(),
		p, // in this problem, we can stay in place
	}
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
	start, goal             Point
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

	for c := 0; c < b.cols; c++ {
		start := Point{c, 0}
		goal := Point{c, b.rows - 1}
		if b.get(0, start) != WallTile {
			b.start = start
		}
		if b.get(0, goal) != WallTile {
			b.goal = goal
		}
	}

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

type Partial struct {
	time int
	pos  Point
}

type WorkItem struct {
	p    Partial
	heur int
}

func (a *WorkItem) Less(b *WorkItem) bool {
	ha := a.p.time + a.heur
	hb := b.p.time + b.heur

	if ha == hb {
		return a.heur < b.heur
	}
	return ha < hb
}

type Work []*WorkItem

func (a Work) Len() int           { return len(a) }
func (a Work) Less(i, j int) bool { return a[i].Less(a[j]) }
func (a Work) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *Work) Push(x any)        { *a = append(*a, x.(*WorkItem)) }

func (a *Work) Pop() any {
	old := *a
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*a = old[0 : n-1]
	return item
}

func (b *Blizzards) shortestPathLength(start, goal Point, t0 int) int {
	cost := map[Partial]int{}
	visited := map[Partial]bool{}
	toVisit := new(Work)
	heap.Init(toVisit)

	{
		w := WorkItem{}
		w.p.pos = start
		w.p.time = t0
		heap.Push(toVisit, &w)
	}

	for toVisit.Len() > 0 {
		entry := heap.Pop(toVisit).(*WorkItem)
		p := entry.p
		if visited[p] {
			continue
		}
		visited[p] = true

		if p.pos == goal {
			return p.time
		}

		t := p.time + 1
		for _, n := range p.pos.Neighbors() {
			if n.x < 0 || n.y < 0 || n.x >= b.cols || n.y >= b.rows {
				continue
			}
			if b.get(t, n) != EmptyTile {
				continue
			}
			part := Partial{t, n}
			if visited[part] {
				continue
			}
			c, found := cost[part]
			if !found || part.time < c {
				cost[part] = part.time
				w := WorkItem{part, 0} // n.Manhattan(goal)}
				heap.Push(toVisit, &w)
			}
		}
	}

	panic("no path found")
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
	return Part1Result(b.shortestPathLength(b.start, b.goal, 0))
}

func DoPart2(b *Blizzards) Part2Result {
	there := int(DoPart1(b))
	back := b.shortestPathLength(b.goal, b.start, there)
	again := b.shortestPathLength(b.start, b.goal, back)
	return Part2Result(again)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
