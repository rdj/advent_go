package aoc_2022_15

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/samber/lo"
)

var _ = fmt.Println
var _ = lo.Max[int]

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 4_961_647

type Part2Result int

const Part2Want = 12_274_327_017_867

type ParsedInput []string

type Point struct{ x, y int }

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (a Point) Manhattan(b Point) int {
	return absdiff(a.x, b.x) + absdiff(a.y, b.y)
}

type Sensor struct {
	pos, beacon Point
}

type Sensors []Sensor

type Range struct{ a, b int }
type Ranges []Range

type ByA Ranges

func (a ByA) Len() int           { return len(a) }
func (a ByA) Less(i, j int) bool { return a[i].a < a[j].a }
func (a ByA) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByB Ranges

func (a ByB) Len() int           { return len(a) }
func (a ByB) Less(i, j int) bool { return a[i].b < a[j].b }
func (a ByB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (ranges *Ranges) compact() {
	sort.Sort(ByA(*ranges))

	after := Ranges{}
	cum := (*ranges)[0]
	for _, r := range (*ranges)[1:] {
		if cum.b+1 >= r.a {
			cum.b = max(cum.b, r.b)
		} else {
			after = append(after, cum)
			cum = r
		}
	}
	after = append(after, cum)

	*ranges = after
}

func (sens Sensors) LineRanges(y int) Ranges {
	ranges := Ranges{}

	for _, s := range sens {
		d := s.pos.Manhattan(s.beacon)
		d -= absdiff(s.pos.y, y)
		if d >= 0 {
			ranges = append(ranges, Range{s.pos.x - d, s.pos.x + d})
		}
	}

	ranges.compact()
	return ranges
}

func (sens Sensors) LineCoverage(y int) int {
	ranges := sens.LineRanges(y)

	sum := 0
	for _, r := range ranges {
		sum += r.b - r.a
	}

	return sum
}

func (sens Sensors) FindBeacon(maxXY int) Point {
	for y := 0; y <= maxXY; y++ {
		ranges := sens.LineRanges(y)
		if len(ranges) > 1 {
			return Point{ranges[0].b + 1, y}
		}
	}
	panic("found no sensor")
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) Sensors {
	sens := Sensors{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		s := Sensor{}
		fmt.Sscanf(
			scanner.Text(),
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.pos.x, &s.pos.y, &s.beacon.x, &s.beacon.y,
		)
		sens = append(sens, s)
	}
	return sens
}

func DoPart1(sens Sensors, y int) Part1Result {
	return Part1Result(sens.LineCoverage(y))
}

func DoPart2(sens Sensors, maxXY int) Part2Result {
	p := sens.FindBeacon(maxXY)
	return Part2Result(p.x*4_000_000 + p.y)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()), 2_000_000)
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()), 4_000_000)
}
