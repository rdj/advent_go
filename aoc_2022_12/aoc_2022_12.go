package aoc_2022_12

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 420

type Part2Result int

const Part2Want = 414

func absdiff(n, m int) int {
	if n > m {
		return n - m
	}
	return m - n
}

type Point struct{ x, y int }

func (p Point) Neighbors() []Point {
	n := make([]Point, 0, 4)
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

const startElevation = 'a'
const destElevation = 'z'

type HeightMap struct {
	start, dest Point
	heights     map[Point]byte
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) HeightMap {
	hm := HeightMap{}
	hm.heights = make(map[Point]byte)
	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		for x, b := range scanner.Text() {
			p := Point{x, y}
			switch b {
			case 'S':
				hm.start = p
				hm.heights[p] = startElevation

			case 'E':
				hm.dest = p
				hm.heights[p] = destElevation

			default:
				hm.heights[p] = byte(b)
			}
		}
	}
	return hm
}

type Path struct {
	pos  Point
	cost int
	heur int
}

func (p *Path) Branch(hm *HeightMap, n Point) *Path {
	newp := *p
	newp.pos = n
	newp.cost++
	newp.heur = n.Manhattan(hm.dest)
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

func (hm *HeightMap) ShortestPathLength() int {
	cost := map[Point]int{}
	visited := map[Point]bool{}
	toVisit := new(Paths)
	heap.Init(toVisit)

	p := new(Path)
	p.pos = hm.start
	p.heur = p.pos.Manhattan(hm.dest)
	heap.Push(toVisit, p)

	for toVisit.Len() > 0 {
		p = heap.Pop(toVisit).(*Path)
		if visited[p.pos] {
			continue
		}
		visited[p.pos] = true

		if p.pos.Eq(hm.dest) {
			return p.cost
		}

		h := hm.heights[p.pos]

		branches := make([]*Path, 0, 4)
		for _, n := range p.pos.Neighbors() {
			hn, ok := hm.heights[n]
			if !ok {
				continue
			}
			if hn > h+1 {
				continue
			}
			branches = append(branches, p.Branch(hm, n))
		}

		for _, b := range branches {
			if visited[b.pos] {
				continue
			}

			c, found := cost[b.pos]
			if !found || b.cost < c {
				cost[b.pos] = b.cost
				heap.Push(toVisit, b)
			}
		}
	}

	return int(^uint(0) >> 1)
}

func DoPart1(hm HeightMap) Part1Result {
	return Part1Result(hm.ShortestPathLength())
}

func DoPart2(hm HeightMap) Part2Result {
	min := 10000
	for p, h := range hm.heights {
		if h == 'a' {
			hm.start = p
			s := hm.ShortestPathLength()
			if s < min {
				min = s
			}
		}
	}
	return Part2Result(min)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
