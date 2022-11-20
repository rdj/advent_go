package aoc_2018_15

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 269_430

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 0xBAAD_F00D

const MAX_HP = 200
const ATTACK_POWER = 3

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type Point struct{ x, y int }

func (p1 Point) Less(p2 Point) bool {
	if p1.y == p2.y {
		return p1.x < p2.x
	}
	return p1.y < p2.y
}

type ReadingOrder []Point

func (a ReadingOrder) Len() int      { return len(a) }
func (a ReadingOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ReadingOrder) Less(i, j int) bool {
	return a[i].Less(a[j])
}

func (a Point) Manhattan(b Point) int {
	return absdiff(a.x, b.x) + absdiff(a.y, b.y)
}

func (p Point) Neighbors() []Point {
	return []Point{
		Point{p.x, p.y - 1},
		Point{p.x - 1, p.y},
		Point{p.x + 1, p.y},
		Point{p.x, p.y + 1},
	}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Unit struct {
	pos  Point
	team byte
	hp   int
}

type Arena struct {
	walls  map[Point]bool
	units  map[Point]*Unit
	noMove map[Point]bool
}

func (a *Arena) adjacentTarget(u *Unit) *Unit {
	var target *Unit
	minHP := int(^uint(0) >> 1)
	for _, n := range u.pos.Neighbors() {
		o := a.units[n]
		if o != nil && o.team != u.team {
			if o.hp < minHP {
				target = o
				minHP = target.hp
			}
		}
	}
	return target
}

func (a *Arena) attack(attacker *Unit, target *Unit) {
	target.hp -= ATTACK_POWER
	if target.hp < 0 {
		delete(a.units, target.pos)
		a.noMove = map[Point]bool{}
	}
}

func (a *Arena) done() bool {
	teams := map[byte]bool{}
	for _, u := range a.units {
		teams[u.team] = true
	}
	return len(teams) == 1
}

func (a *Arena) openNeighborsOfEnemies(u *Unit) []Point {
	added := map[Point]bool{}
	dests := make([]Point, 0)
	for _, e := range a.units {
		if e.team == u.team {
			continue
		}
		for _, xy := range a.openNeighbors(e.pos) {
			if !added[xy] {
				added[xy] = true
				dests = append(dests, xy)
			}
		}
	}
	return dests
}

type Partial struct {
	dsts []Point
	path []Point
}

func (p Partial) complete() bool {
	if len(p.path) == 0 {
		return false
	}
	pos := p.pos()
	for _, dst := range p.dsts {
		if pos == dst {
			return true
		}
	}
	return false
}

func (p Partial) heur() int {
	cost := int(^uint(0) >> 1)

	pos := p.pos()
	for _, dst := range p.dsts {
		man := pos.Manhattan(dst)
		if man < cost {
			cost = man
		}
	}

	cost += len(p.path)
	return cost
}

func (p Partial) pos() Point {
	return p.path[len(p.path)-1]
}

func (p Partial) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s heur=%d : %s", p.pos(), p.heur(), p.path)
	return sb.String()
}

type Partials []Partial

func (p Partials) Len() int { return len(p) }
func (p Partials) Less(i, j int) bool {
	if p[i].heur() == p[j].heur() {
		if len(p[i].path) < 2 || len(p[j].path) < 2 {
			return false
		}
		return p[i].path[1].Less(p[j].path[1])
	}
	return p[i].heur() < p[j].heur()
}
func (p Partials) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *Partials) Push(x any) {
	*p = append(*p, x.(Partial))
}

func (p *Partials) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

var debugMove bool = false

func (a *Arena) move(u *Unit) {
	if a.noMove[u.pos] {
		return
	}

	dests := a.openNeighborsOfEnemies(u)
	if len(dests) == 0 {
		return
	}
	if debugMove {
		fmt.Println("Destinations", dests)
	}

	parts := &Partials{}
	heap.Init(parts)
	heap.Push(parts, Partial{dests, []Point{u.pos}})

	type Best struct {
		cost  int
		first Point
	}
	bestSeen := map[Point]Best{}

	for len(*parts) > 0 {
		part := heap.Pop(parts).(Partial)
		// debugMove = debugMove || (part.path[0].x == 26 && part.path[0].y == 16)

		if debugMove {
			fmt.Println("Exploring", part)
		}
		if part.complete() {
			// for _, foo := range *parts {
			// 	fmt.Println(foo)
			// }
			// fmt.Println("Complete with ", part)
			u.pos = part.path[1]
			a.noMove = map[Point]bool{}
			return
		}

		newLen := len(part.path) + 1
		for _, next := range a.openNeighbors(part.pos()) {
			// for _, seen := range part.path {
			// 	if seen == next {
			// 		if debugMove {
			// 			fmt.Println("Avoiding loopback to ", next, " given path ", part.path)
			// 		}
			// 		continue Next
			// 	}
			// }

			var first Point
			if len(part.path) > 1 {
				first = part.path[1]
			} else {
				first = next
			}

			if best, ok := bestSeen[next]; ok {
				if best.cost < newLen {
					continue // prune, beaten on pure cost
				}
				if best.cost == newLen && best.first.Less(first) {
					continue // prune, beaten on first step reading order
				}
			}

			bestSeen[next] = Best{newLen, first}
			path := make([]Point, len(part.path), newLen)
			copy(path, part.path)
			path = append(path, next)
			nextPart := Partial{part.dsts, path}
			if debugMove {
				fmt.Println("Adding", nextPart)
			}
			heap.Push(parts, nextPart)
		}
	}

	a.noMove[u.pos] = true
}

func (a *Arena) openNeighbors(o Point) []Point {
	points := make([]Point, 0, 4)
	for _, p := range o.Neighbors() {
		// if p.x == 21 && p.y == 14 {
		// 	if a.walls[p] {
		// 		fmt.Println("Dropping", p, "due to wall")
		// 	} else if a.units[p] != nil {
		// 		fmt.Println("Dropping", p, "due to unit")
		// 		fmt.Println(a)
		// 	}
		// }
		if !a.walls[p] && a.units[p] == nil {
			points = append(points, p)
		}
	}
	return points
}

func (a *Arena) Run() int {
	rounds := 0
	// fmt.Println("Round", rounds)
	// fmt.Println(a)
	for {
		//debugMove = rounds == 24
		for n, pos := range a.turnOrder() {
			if debugMove {
				fmt.Println("Round", rounds, "Turn", n, pos)
			}
			if a.done() {
				return rounds
			}

			u := a.units[pos]
			delete(a.units, pos)
			if u == nil {
				// unit died before its turn this round
				continue
			}
			a.takeTurn(u)
			a.units[u.pos] = u
		}

		rounds++
		fmt.Println("Round", rounds)
		//fmt.Println(a)
	}
}

func (a *Arena) String() string {
	var sb strings.Builder

	walls := []Point{}
	for p, _ := range a.walls {
		walls = append(walls, p)
	}
	sort.Sort(ReadingOrder(walls))
	end := walls[len(walls)-1]
	units := []*Unit{}

	for y := 0; y <= end.y; y++ {
		if y > 0 {
			for i, u := range units {
				if i == 0 {
					sb.WriteString("  ")
				} else {
					sb.WriteString(", ")
				}
				fmt.Fprintf(&sb, "%c(%d)", u.team, u.hp)
			}
			units = []*Unit{}
			sb.WriteByte('\n')
		}
		for x := 0; x <= end.x; x++ {
			p := Point{x, y}
			if u := a.units[p]; u != nil {
				sb.WriteByte(u.team)
				units = append(units, u)
			} else if a.walls[p] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
	}
	return sb.String()
}

func (a *Arena) SumHP() int {
	sum := 0
	for _, u := range a.units {
		sum += u.hp
	}
	return sum
}

func (a *Arena) takeTurn(u *Unit) {
	target := a.adjacentTarget(u)
	if target == nil {
		a.move(u)
		target = a.adjacentTarget(u)
	}
	if target != nil {
		a.attack(u, target)
	}
}

func (a *Arena) turnOrder() []Point {
	points := make([]Point, 0, len(a.units))
	for p := range a.units {
		points = append(points, p)
	}
	sort.Sort(ReadingOrder(points))
	return points
}

// Board with Walls, Spaces, Units
// Units
// - have hitpoints
// - have attack power
// - must be able to enumerate all enemies
// - run same combat logic
// Pathfinding, collision avoidance, shortest path, tiebroken on first step reading order

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) *Arena {
	a := Arena{
		walls:  map[Point]bool{},
		units:  map[Point]*Unit{},
		noMove: map[Point]bool{},
	}
	lines := bufio.NewScanner(input)
	for y := 0; lines.Scan(); y++ {
		for x, c := range lines.Text() {
			switch c {
			case '#':
				a.walls[Point{x, y}] = true
			case 'G', 'E':
				p := Point{x, y}
				a.units[p] = &Unit{p, byte(c), MAX_HP}
			}
		}
	}
	return &a
}

func DoPart1(arena *Arena) Part1Result {
	rounds := arena.Run()
	hp := arena.SumHP()
	result := rounds * hp
	fmt.Println("Rounds", rounds, "HP", hp, "Result", result)
	fmt.Println(arena)
	// 222852 is too low
	return Part1Result(result)
}

func DoPart2(arena *Arena) Part2Result {
	return Part2Fake
}

func Part1() Part1Result {
	//return Part1Fake
	//debugMove = true
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
