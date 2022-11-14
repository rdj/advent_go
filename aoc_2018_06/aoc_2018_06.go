package aoc_2018_06

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var _ = fmt.Println
var _ = strconv.Atoi

const inputFile = "input.txt"

type AdventResult int

type point struct{ x, y int }

func openInput() io.Reader {
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

func ParseInput(input io.Reader) []point {
	points := make([]point, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var p point
		fmt.Sscanf(scanner.Text(), "%d, %d", &p.x, &p.y)
		points = append(points, p)
	}
	return points
}

func dist(a, b point) int {
	dist := 0
	if a.x > b.x {
		dist += a.x - b.x
	} else {
		dist += b.x - a.x
	}
	if a.y > b.y {
		dist += a.y - b.y
	} else {
		dist += b.y - a.y
	}
	return dist
}

func DoPart1(points []point) AdventResult {
	type Bounds struct{ xmin, xmax, ymin, ymax int }
	bounds := Bounds{int(^uint(0) >> 1), -1, int(^uint(0) >> 1), -1}
	for _, p := range points {
		if p.x < bounds.xmin {
			bounds.xmin = p.x
		}
		if p.x > bounds.xmax {
			bounds.xmax = p.x
		}
		if p.y < bounds.ymin {
			bounds.ymin = p.y
		}
		if p.y > bounds.ymax {
			bounds.ymax = p.y
		}
	}

	type Tile struct {
		id   byte
		dist uint16
	}
	tiles := make([][]Tile, bounds.ymax+bounds.ymin)
	for i := range tiles {
		tiles[i] = make([]Tile, bounds.xmax+bounds.xmin+1)
	}

	for id, p := range points {
		for y, row := range tiles {
			for x, tile := range row {
				dist := uint16(dist(p, point{x, y}))

				switch {
				case tile.id == 0 && tile.dist == 0, dist < tile.dist: // unclaimed or closest
					tiles[y][x] = Tile{byte(id + 1), uint16(dist)}

				case tile.dist == dist: // tie
					tiles[y][x] = Tile{0, dist}
				}
			}
		}
	}

	// I mistakenly thought that only points with a maximal x or y
	// coordinate would have infinite area. That's works for the
	// example but not for the real input. A better test is whether
	// any "edge" tile is closest to that point.
	edges := map[byte]bool{}
	for _, t := range tiles[0] {
		edges[t.id] = true
	}
	for _, t := range tiles[len(tiles)-1] {
		edges[t.id] = true
	}
	for _, row := range tiles {
		edges[row[0].id] = true
		edges[row[len(row)-1].id] = true
	}

	type Max struct {
		id    byte
		count int
	}

	max := Max{}

	for id := range points {
		m := Max{byte(id + 1), 0}
		if edges[m.id] {
			continue
		}
		for _, row := range tiles {
			for _, tile := range row {
				if tile.id == m.id {
					m.count += 1
				}
			}
		}
		if m.count > max.count {
			max = m
		}
	}

	return AdventResult(max.count)
}

func DoPart2(points []point) AdventResult {
	return AdventResult(0)
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() AdventResult {
	return DoPart2(ParseInput(openInput()))
}
