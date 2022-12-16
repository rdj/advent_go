package aoc_2022_16

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 1751

type Part2Result int

const Part2Want = 2207

type Valve struct {
	name    string
	flow    int
	tunnels []string
}

func NewValve() Valve {
	return Valve{tunnels: []string{}}
}

type Valves map[string]Valve

// Computes the pairwise shortest paths in a weighted digraph. A little bit overkill, but simple to code.
//
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
func (valves Valves) distances() map[string]map[string]int {
	dist := map[string]map[string]int{}

	for v := range valves {
		dist[v] = map[string]int{v: 0}
		for _, t := range valves[v].tunnels {
			dist[v][t] = 1
		}
	}

	for k := range valves {
		for i := range valves {
			for j := range valves {
				ik, ok := dist[i][k]
				if !ok {
					continue
				}
				kj, ok := dist[k][j]
				if !ok {
					continue
				}

				ij, ok := dist[i][j]
				if !ok || ij > ik+kj {
					dist[i][j] = ik + kj
				}
			}
		}
	}

	return dist
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) Valves {
	valves := Valves{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		v := Valve{}
		_, err := fmt.Sscanf(scanner.Text(), "Valve %s has flow rate=%d;", &v.name, &v.flow)
		if err != nil {
			panic(err)
		}

		tuns := strings.Split(line, ", ")
		for _, t := range tuns {
			t := string(t[len(t)-2:])
			v.tunnels = append(v.tunnels, t)
		}

		valves[v.name] = v
	}

	return valves
}

type Actor struct {
	pos  string
	time int
}

type Partial struct {
	actors [2]Actor
	flow   int
	open   int
}

func (valves Valves) BestFlow(timeLimit int, twoActors bool) int {
	toOpen := map[string]int{}
	for _, v := range valves {
		if v.flow > 0 {
			toOpen[v.name] = len(toOpen)
		}
	}

	type Seen struct {
		pos, open int
	}
	best := map[Seen]int{}

	dist := valves.distances()

	q := []Partial{Partial{
		[2]Actor{
			Actor{"AA", 1},
			Actor{"AA", 1},
		},
		0,
		0,
	}}

	for len(q) > 0 {
		p := q[len(q)-1]
		q = q[:len(q)-1]

		a := 0
		if twoActors && p.actors[1].time < p.actors[0].time {
			a = 1
		}

		for v, vno := range toOpen {
			vflag := 1 << vno

			if 0 != p.open&vflag {
				continue
			}

			np := Partial{
				actors: p.actors,
				flow:   p.flow,
				open:   p.open,
			}

			np.actors[a].pos = v
			np.actors[a].time += dist[p.actors[a].pos][v] + 1
			np.flow += (timeLimit + 1 - np.actors[a].time) * valves[v].flow
			np.open |= vflag

			seen := Seen{pos: 2*toOpen[np.actors[0].pos] + 3*toOpen[np.actors[1].pos], open: np.open}
			if np.flow < best[seen] {
				continue
			}
			best[seen] = np.flow

			//fmt.Printf("len=%d; t=%d; flow=%d; path=%v\n", len(np.path), np.time, np.flow, np.path)

			q = append(q, np)
		}
	}

	return lo.Max(lo.Values(best))
}

func DoPart1(valves Valves) Part1Result {
	return Part1Result(valves.BestFlow(30, false))
}

func DoPart2(valves Valves) Part2Result {
	return Part2Result(valves.BestFlow(26, true))
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
