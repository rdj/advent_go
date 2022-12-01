package aoc_2018_15

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

const inputFile = "input.txt"

type Part1Result int

const Part1Fake = 0xDEAD_BEEF
const Part1Want = 269_430

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 55_160

const maxHP = 200
const defaultAttackPower = 3
const TeamElf = 'E'

type Point struct{ x, y int }

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

type Unit struct {
	pos  Point
	team byte
	hp   int
}

func (u *Unit) isElf() bool {
	return u.team == TeamElf
}

type Arena struct {
	walls map[Point]bool
	units map[Point]*Unit
	elfAP int
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
	attackPower := defaultAttackPower
	if attacker.isElf() {
		attackPower = a.elfAP
	}

	target.hp -= attackPower
	if target.hp <= 0 {
		delete(a.units, target.pos)
	}
}

func (a *Arena) done() bool {
	teams := map[byte]bool{}
	for _, u := range a.units {
		teams[u.team] = true
	}
	return len(teams) == 1
}

func (a *Arena) elfCount() int {
	n := 0
	for _, u := range a.units {
		if u.isElf() {
			n++
		}
	}
	return n
}

func getBestMove(paths [][]Point, dests map[Point]bool) Point {
	// The path selection conditions are excrutiatingly
	// tedious. Based on reddit threads, the minute details
	// here matter for some people's inputs but not others.
	//
	// Once all shortest length paths to all destinations have
	// been found, the winning path is selected first based on
	// the reading order of the destination square, then based
	// on the reading order of the first step in the path.
	//
	// For all of the examples and Part 1, it is sufficient to
	// use the first shortest path found by seeding BFS with
	// the initial steps in reading order.
	//
	// However, for part 2 of my input, the destination square
	// tiebreaker makes a difference, and causes a unit to
	// take a later-in-reading-order first step in order to
	// move towards an earlier-in-reading-order destination
	// square.

	bestPath := paths[0]
	paths = paths[1:]
	bestLen := len(bestPath)
	bestPos := bestPath[bestLen-1]

	for _, p := range paths {
		if len(p) > bestLen {
			break
		}

		pd := p[len(p)-1]
		switch {
		case !dests[pd]: // not a dest
			continue
		case bestPos.Less(pd): // worse dest
			continue
		case pd.Less(bestPos): // better dest
			fallthrough
		case p[1].Less(bestPath[1]): // better first step (same dest)
			bestPos = pd
			bestPath = p
		default:
			continue
		}
	}

	return bestPath[1]
}

func (a *Arena) move(u *Unit) {
	dests := a.openNeighborsOfEnemies(u)
	if len(dests) == 0 {
		return
	}

	// BFS. Current work item is always shortest and earliest in
	// reading order.
	paths := [][]Point{[]Point{u.pos}}
	seen := map[Point]bool{}

	for len(paths) > 0 {
		path := paths[0]

		pos := path[len(path)-1]
		if dests[pos] {
			u.pos = getBestMove(paths, dests)
			return
		}

		paths[0] = nil
		paths = paths[1:]

		newLen := len(path) + 1
		for _, next := range a.openNeighbors(pos) {
			if seen[next] {
				continue
			}
			seen[next] = true

			newPath := make([]Point, newLen-1, newLen)
			copy(newPath, path)
			newPath = append(newPath, next)
			paths = append(paths, newPath)
		}
	}
}

func (a *Arena) openNeighbors(o Point) []Point {
	points := make([]Point, 0, 4)
	for _, p := range o.Neighbors() {
		if !a.walls[p] && a.units[p] == nil {
			points = append(points, p)
		}
	}
	return points
}

func (a *Arena) openNeighborsOfEnemies(u *Unit) map[Point]bool {
	set := map[Point]bool{}
	for _, e := range a.units {
		if e.team == u.team {
			continue
		}
		for _, xy := range a.openNeighbors(e.pos) {
			set[xy] = true
		}
	}
	return set
}

func (a *Arena) Run() int {
	rounds := 0
	for {
		for _, pos := range a.turnOrder() {
			u := a.units[pos]
			if u == nil {
				// unit died before its turn this round
				continue
			}

			if a.done() {
				return rounds
			}

			delete(a.units, pos)

			a.takeTurn(u)
			a.units[u.pos] = u
		}

		rounds++
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

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) *Arena {
	a := Arena{
		walls: map[Point]bool{},
		units: map[Point]*Unit{},
		elfAP: defaultAttackPower,
	}
	lines := bufio.NewScanner(input)
	for y := 0; lines.Scan(); y++ {
		for x, c := range lines.Text() {
			switch c {
			case '#':
				a.walls[Point{x, y}] = true
			case 'G', 'E':
				p := Point{x, y}
				a.units[p] = &Unit{p, byte(c), maxHP}
			}
		}
	}
	return &a
}

func DoPart1(arena *Arena) Part1Result {
	rounds := arena.Run()
	hp := arena.SumHP()
	result := rounds * hp
	return Part1Result(result)
}

func DoPart2(reader io.Reader) Part2Result {
	input, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	// I originally used binary search, but reddit says for some
	// inputs there are failing cases with higher AP values than
	// success cases. It's fast enough to just count up.
	var arena *Arena
	var rounds int
	deadElves := 1
	for ap := defaultAttackPower + 1; deadElves != 0; ap++ {
		arena = ParseInput(bytes.NewReader(input))
		arena.elfAP = ap
		elvesBefore := arena.elfCount()
		rounds = arena.Run()
		deadElves = elvesBefore - arena.elfCount()
	}

	outcome := arena.SumHP() * rounds
	return Part2Result(outcome)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(openInput())
}
