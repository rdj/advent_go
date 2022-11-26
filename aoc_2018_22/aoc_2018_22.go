package aoc_2018_22

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
)

const (
	xMultiplier = 16807 // 7⁵
	yMultiplier = 48271 // prime
	modulus     = 20183 // prime

	//         N  W  R
	// None    1  1  0
	// Gear    0  1  1
	// Torch   1  0  1

	rocky  = 0b001
	wet    = 0b010
	narrow = 0b100

	noTool = 0b110
	gear   = 0b011
	torch  = 0b101

	toolCost = 7
)

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 10204

type Part2Result int

const Part2Want = 1004

func absdiff(n, m int) int {
	if n > m {
		return n - m
	}
	return m - n
}

type Point struct{ x, y int }

func (p Point) Neighbors() []Point {
	n := make([]Point, 4)
	if p.y > 0 {
		n = append(n, p.Up())
	}
	if p.x > 0 {
		n = append(n, p.Left())
	}
	n = append(n, p.Right(), p.Down())
	return n
}

func (a Point) Eq(b Point) bool {
	return a.x == b.x && a.y == b.y
}

func (a Point) Manhattan(b Point) int {
	return absdiff(a.x, b.x) + absdiff(a.y, b.y)
}

func (p Point) Up() Point {
	return Point{p.x, p.y - 1}
}

func (p Point) Left() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) Right() Point {
	return Point{p.x + 1, p.y}
}

func (p Point) Down() Point {
	return Point{p.x, p.y + 1}
}

type Cave struct {
	depth    int
	target   Point
	geo, ero map[Point]int
}

func NewCave() *Cave {
	c := new(Cave)
	c.geo = make(map[Point]int)
	c.ero = make(map[Point]int)
	return c
}

func (c *Cave) Erosion(p Point) int {
	if e, ok := c.ero[p]; ok {
		return e
	}

	e := (c.Geology(p) + c.depth) % modulus
	c.ero[p] = e
	return e
}

func (c *Cave) Risk(p Point) int {
	return c.Erosion(p) % 3
}

func (c *Cave) Terrain(p Point) int {
	return 1 << c.Risk(p)
}

func (c *Cave) TotalRisk() int {
	risk := 0
	p := Point{}
	for p.y = 0; p.y <= c.target.y; p.y++ {
		for p.x = 0; p.x <= c.target.x; p.x++ {
			risk += c.Risk(p)
		}
	}

	return risk
}

func (c *Cave) Geology(p Point) int {
	if g, ok := c.geo[p]; ok {
		return g
	}

	var g int
	switch {
	case p.Eq(Point{0, 0}) || p.Eq(c.target):
		g = 0
	case p.Eq(Point{1, 0}):
		g = xMultiplier
	case p.y == 0:
		g = c.Geology(p.Left()) + xMultiplier
	case p.Eq(Point{0, 1}):
		g = yMultiplier
	case p.x == 0:
		g = c.Geology(p.Up()) + yMultiplier
	default:
		g = c.Erosion(p.Left()) * c.Erosion(p.Up())
	}

	g %= modulus
	c.geo[p] = g
	return g
}

type Path struct {
	pos  Point
	cost int
	tool int
	heur int
}

func (p *Path) Branch(c *Cave, n Point) *Path {
	newp := *p
	newp.pos = n
	newp.cost++

	nter := c.Terrain(n)
	if p.tool&nter == 0 {
		newp.SetTool(nter | c.Terrain(p.pos))
	}

	newp.heur = n.Manhattan(c.target)
	if newp.tool != torch {
		newp.heur += toolCost
	}

	return &newp
}

func (p1 *Path) Less(p2 *Path) bool {
	h1 := p1.cost + p1.heur
	h2 := p2.cost + p2.heur

	if h1 == h2 {
		return p1.heur < p2.heur
	}
	return h1 < h2
}

func (p *Path) Seen() Seen {
	return Seen{p.pos, p.tool}
}

func (p *Path) SetTool(tool int) {
	p.tool = tool
	p.cost += toolCost
}

type Paths []*Path

func (p Paths) Len() int           { return len(p) }
func (p Paths) Less(i, j int) bool { return p[i].Less(p[j]) }
func (p Paths) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *Paths) Push(x any)        { *p = append(*p, x.(*Path)) }

func (p *Paths) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*p = old[0 : n-1]
	return item
}

type Seen struct {
	p    Point
	tool int
}

func (c *Cave) ShortestPathLength() int {
	// Started with Dijkstra, then added an A-Star heuristic cost for
	// big efficiency gain.
	//
	// Given our Part 2 input, Dijkstra visits ~798k nodes (runtime
	// ~1.3s), but A-star cuts that to ~142k (runtime ~0.4s).
	cost := map[Seen]int{}
	visited := map[Seen]bool{}
	toVisit := new(Paths)
	heap.Init(toVisit)

	p := new(Path)
	p.tool = torch
	p.heur = p.pos.Manhattan(c.target)
	heap.Push(toVisit, p)

	for {
		p = heap.Pop(toVisit).(*Path)
		ps := p.Seen()
		if visited[ps] {
			continue
		}
		visited[ps] = true

		if p.pos.Eq(c.target) && p.tool == torch {
			return p.cost
		}

		branches := make([]*Path, 0, 4)
		for _, n := range p.pos.Neighbors() {
			branches = append(branches, p.Branch(c, n))
		}
		// May need a final step that has no movement, only a tool
		// change
		if p.pos.Eq(c.target) {
			b := *p
			b.SetTool(torch)
			b.heur = 0
			branches = append(branches, &b)
		}

		for _, b := range branches {
			bs := b.Seen()
			if visited[bs] {
				continue
			}

			c, found := cost[bs]
			if !found || b.cost < c {
				cost[bs] = b.cost
				heap.Push(toVisit, b)
			}
		}
	}
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) *Cave {
	c := NewCave()

	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		panic("bad input")
	}
	fmt.Sscanf(scanner.Text(), "depth: %d", &c.depth)

	if !scanner.Scan() {
		panic("bad input")
	}
	fmt.Sscanf(scanner.Text(), "target: %d,%d", &c.target.x, &c.target.y)

	return c
}

func DoPart1(c *Cave) Part1Result {
	return Part1Result(c.TotalRisk())
}

func DoPart2(c *Cave) Part2Result {
	return Part2Result(c.ShortestPathLength())
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
