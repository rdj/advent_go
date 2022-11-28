package aoc_2018_24

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 18346

type Part2Result int

const Part2Want = 8698

type CombatGroup struct {
	team        int
	id          int
	units       int
	hp          int
	initiative  int
	attackPower int
	attackType  string
	immunities  map[string]bool
	weaknesses  map[string]bool
}

func NewCombatGroup() *CombatGroup {
	g := new(CombatGroup)
	g.immunities = make(map[string]bool)
	g.weaknesses = make(map[string]bool)
	return g
}

func (a *CombatGroup) damageTo(d *CombatGroup) int {
	if d.immunities[a.attackType] {
		return 0
	}
	pow := a.effectivePower()
	if d.weaknesses[a.attackType] {
		pow *= 2
	}
	return pow
}

func (g *CombatGroup) effectivePower() int {
	return g.units * g.attackPower
}

func teamName(n int) string {
	switch n {
	case 0:
		return "Immune System"
	case 1:
		return "Infection"
	default:
		return "Unknown"
	}
}

func (g *CombatGroup) String() string {
	return fmt.Sprintf("%s group %d", teamName(g.team), g.id)
}

func (d *CombatGroup) takeDamage(a *CombatGroup) int {
	killed := a.damageTo(d) / d.hp
	if killed > d.units {
		killed = d.units
	}
	d.units -= killed
	return killed
}

type ByPower []*CombatGroup

func (s ByPower) Len() int      { return len(s) }
func (s ByPower) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPower) Less(i, j int) bool {
	ip := s[i].effectivePower()
	jp := s[j].effectivePower()
	if ip == jp {
		return s[j].initiative < s[i].initiative
	}
	return jp < ip // higher power first
}

type ByInitiative []*CombatGroup

func (s ByInitiative) Len() int           { return len(s) }
func (s ByInitiative) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByInitiative) Less(i, j int) bool { return s[j].initiative < s[i].initiative }

type ById []*CombatGroup

func (s ById) Len() int           { return len(s) }
func (s ById) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ById) Less(i, j int) bool { return s[i].id < s[j].id }

type CombatGroups []*CombatGroup

func (c CombatGroups) boost(n int) {
	for _, g := range c {
		if g.team == 0 {
			g.attackPower += n
		}
	}
}

func (c CombatGroups) finished() (bool, int) {
	teams := map[int]bool{}
	for _, g := range c {
		if g.units == 0 {
			continue
		}
		teams[g.team] = true
		if len(teams) > 1 {
			return false, -1
		}
	}

	for t := range teams {
		return true, t
	}

	return false, 0
}

func (c CombatGroups) totalUnits() int {
	units := 0
	for _, g := range c {
		units += g.units
	}
	return units
}

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) CombatGroups {
	inputPattern := regexp.MustCompile(`(\d+) units each with (\d+) hit points (?:\(([^)]+)\) )?with an attack that does (\d+) (\S+) damage at initiative (\d+)`)
	modPattern := regexp.MustCompile(`(weak|immune) to ([^;]+)`)

	groups := make(CombatGroups, 0)
	scanner := bufio.NewScanner(input)
	team := 0
	id := 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line == "Immune System:" {
			continue
		}
		if line == "Infection:" {
			team++
			id = 1
			continue
		}

		m := inputPattern.FindStringSubmatch(line)
		if m == nil {
			panic(fmt.Sprintf("bad input %q", line))
		}

		var err error

		g := NewCombatGroup()
		g.team = team
		g.id = id
		id++
		g.units, err = strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}

		g.hp, err = strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}

		mps := modPattern.FindAllStringSubmatch(m[3], -1)
		if mps != nil {
			for _, mod := range mps {
				types := strings.Split(mod[2], ", ")
				for _, t := range types {
					switch mod[1] {
					case "weak":
						g.weaknesses[t] = true
					case "immune":
						g.immunities[t] = true
					default:
						panic("bad mod")
					}
				}
			}
		}

		g.attackPower, err = strconv.Atoi(m[4])
		if err != nil {
			panic(err)
		}

		g.attackType = m[5]

		g.initiative, err = strconv.Atoi(m[6])
		if err != nil {
			panic(err)
		}

		groups = append(groups, g)
	}
	return groups
}

func simulateRound(groups CombatGroups) {
	selected := map[*CombatGroup]bool{}
	targets := map[*CombatGroup]*CombatGroup{}

	sort.Sort(ByPower(groups))

	for _, attacker := range groups {
		if attacker.units == 0 {
			continue
		}
		var bestTarget *CombatGroup = nil
		bestDamage := 0
		for _, defender := range groups {
			if selected[defender] || attacker == defender || attacker.team == defender.team || defender.units == 0 {
				continue
			}

			dmg := attacker.damageTo(defender)
			if debugCombat {
				fmt.Printf("%s would deal %s %d damage\n", attacker, defender, dmg)
			}
			if dmg == 0 || dmg < bestDamage {
				continue
			}

			if dmg == bestDamage {
				bestPower := bestTarget.effectivePower()
				pow := defender.effectivePower()
				if bestPower > pow {
					continue
				}

				if bestPower == pow {
					if bestTarget.initiative > defender.initiative {
						continue
					}
				}
			}

			bestTarget = defender
			bestDamage = dmg
		}
		if bestTarget != nil {
			selected[bestTarget] = true
			targets[attacker] = bestTarget
		}
	}

	if debugCombat {
		fmt.Println()
	}

	sort.Sort(ByInitiative(groups))
	for _, attacker := range groups {
		if attacker.units == 0 {
			continue
		}
		defender := targets[attacker]
		if defender == nil {
			continue
		}
		killed := defender.takeDamage(attacker)
		if debugCombat {
			fmt.Printf("%s attacked %s, killing %d units\n", attacker, defender, killed)
		}
	}
}

const maxRounds = 10_000

func (groups CombatGroups) simulateCombat() {
	for round := 0; round < maxRounds; round++ {
		if done, _ := groups.finished(); done {
			break
		}
		if debugCombat {
			sort.Sort(ById(groups))
			fmt.Println("## Round", round)
			for team := 0; team <= 1; team++ {
				fmt.Printf("%s:\n", teamName(team))
				for _, g := range groups {
					if g.team == team && g.units > 0 {
						fmt.Printf("Group %d contains %d units\n", g.id, g.units)
					}
				}
			}
			fmt.Println()
		}
		simulateRound(groups)
		if debugCombat {
			fmt.Println()
		}
	}
}

var debugCombat bool = false

func DoPart1(groups CombatGroups) Part1Result {
	groups.simulateCombat()
	return Part1Result(groups.totalUnits())
}

func DoPart2(reader io.Reader) Part2Result {
	input, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	// Similar to a previous puzzle, binary search does not always
	// work out, we just need to count up from the bottom. With my
	// input, there were a couple of boost values before the target
	// value that deadlocked, so I added a maximum number of rounds to
	// simulate.

	for boost := 1; ; boost++ {
		groups := ParseInput(bytes.NewReader(input))
		groups.boost(boost)
		groups.simulateCombat()
		_, winner := groups.finished()

		if winner == 0 {
			return Part2Result(groups.totalUnits())
		}
	}
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(openInput())
}
