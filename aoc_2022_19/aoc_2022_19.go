package aoc_2022_19

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 1177

type Part2Result int

const Part2Want = 62_744

const (
	Ore = iota
	Clay
	Obsidian
	Geode
	ResourceCount
)

type Resources [ResourceCount]int

type Blueprint struct {
	id   int
	cost [ResourceCount]Resources
}

type States struct {
	a []*State
	b *Blueprint
	t int
}

func (s States) Len() int           { return len(s.a) }
func (s States) Less(i, j int) bool { return s.a[j].MaxYield(s.b, s.t) < s.a[i].MaxYield(s.b, s.t) }
func (s States) Swap(i, j int)      { s.a[i], s.a[j] = s.a[j], s.a[i] }
func (s *States) Push(x any)        { s.a = append(s.a, x.(*State)) }

func (s *States) Pop() any {
	old := s.a
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	s.a = old[0 : n-1]
	return item
}

func (b *Blueprint) BestGeodeProduction(totalMinutes int) int {
	visited := map[State]bool{}
	q := &States{a: []*State{InitialState()}, b: b, t: totalMinutes}

	for q.Len() > 0 {
		s := heap.Pop(q).(*State)

		if visited[*s] {
			continue
		}

		visited[*s] = true

		if s.time == totalMinutes {
			return s.inventory[Geode]
		}

		for _, b := range s.Branches(b) {
			if visited[*b] {
				continue
			}
			// prune more probably
			heap.Push(q, b)
		}
	}
	panic("search failed")
}

func (b *Blueprint) MaxWanted(r int) int {
	if r == Geode {
		return int(^uint(0) >> 1)
	}
	m := 0
	for c := 0; c < ResourceCount; c++ {
		if b.cost[c][r] > m {
			m = b.cost[c][r]
		}
	}
	return m
}

type State struct {
	time       int
	inventory  Resources
	production Resources
}

func InitialState() *State {
	initial := new(State)
	initial.production[Ore] = 1
	return initial
}

func (s *State) MaxYield(b *Blueprint, totalMinutes int) int {
	// Uhhh ok what's a dumb heuristic we can come up with that
	// doesn't require that much math. Ignoring all constraints, the
	// best possible outcome is to increase geode production by one
	// each minute.
	p := s.production[Geode]
	i := s.inventory[Geode]
	// surely we could do this with math but the loop requires no thought
	for t := s.time; t < totalMinutes; t++ {
		i += p
		p += 1
	}
	return i
}

func (s *State) Advance() {
	s.time++
	for r := 0; r < ResourceCount; r++ {
		s.inventory[r] += s.production[r]
	}
}

func (before *State) Branches(b *Blueprint) []*State {
	branches := []*State{}

	base := *before
	base.Advance()
	branches = append(branches, &base)

	for r := 0; r < ResourceCount; r++ {
		// Prune? - Some purchases must be bad. Best if this is a
		// priori and doesn't require analysis of the previous state.
		if before.production[r] >= b.MaxWanted(r) {
			continue
		}

		if before.CanPurchase(b, r) {
			buy := base
			buy.Purchase(b, r)
			branches = append(branches, &buy)
		}
	}

	return branches
}

func (s *State) CanPurchase(b *Blueprint, want int) bool {
	for r := 0; r < ResourceCount; r++ {
		if s.inventory[r] < b.cost[want][r] {
			return false
		}
	}
	return true
}

func (s *State) Purchase(b *Blueprint, want int) {
	for r := 0; r < ResourceCount; r++ {
		s.inventory[r] -= b.cost[want][r]
	}
	s.production[want]++
}

func ParseInput(input io.Reader) []*Blueprint {
	bps := make([]*Blueprint, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		bp := new(Blueprint)
		_, err := fmt.Sscanf(scanner.Text(), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.id, &bp.cost[Ore][Ore], &bp.cost[Clay][Ore], &bp.cost[Obsidian][Ore], &bp.cost[Obsidian][Clay], &bp.cost[Geode][Ore], &bp.cost[Geode][Obsidian])
		if err != nil {
			panic(err)
		}
		bps = append(bps, bp)
	}
	return bps
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func DoPart1(bps []*Blueprint) Part1Result {
	total := 0
	for _, b := range bps {
		v := b.BestGeodeProduction(24)
		total += b.id * v
	}

	return Part1Result(total)
}

func DoPart2(bps []*Blueprint) Part2Result {
	if len(bps) > 3 {
		bps = bps[:3]
	}
	total := 1
	for _, b := range bps {
		v := b.BestGeodeProduction(32)
		total *= v
	}
	return Part2Result(total)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
