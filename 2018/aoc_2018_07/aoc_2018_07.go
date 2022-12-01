package aoc_2018_07

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var _ = fmt.Println

const inputFile = "input.txt"

type AdventResult string

func openInput() io.Reader {
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return data
}

type Depend struct{ dependsOn, id int }

func ParseInput(input io.Reader) []Depend {
	depends := make([]Depend, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var before, after rune
		fmt.Sscanf(scanner.Text(), "Step %c must be finished before step %c can begin.", &before, &after)
		depends = append(depends, Depend{int(before - 'A'), int(after - 'A')})
	}
	return depends
}

type Todo struct {
	tasks map[int]uint
	ids   []int
}

func buildTodo(inputs []Depend) Todo {
	todo := Todo{map[int]uint{}, []int{}}

	for _, spec := range inputs {
		todo.tasks[spec.id] |= 1 << spec.dependsOn

		// one id only appears as a dependency and we for sure need it in the map
		todo.tasks[spec.dependsOn] |= 0
	}

	// if multiple tasks are satisfied, lower ids get preference
	for id := range todo.tasks {
		todo.ids = append(todo.ids, id)
	}
	sort.Ints(todo.ids)

	return todo
}

func (todo *Todo) next(satisfied uint) (int, bool) {
	for id := range todo.ids {
		deps, present := todo.tasks[id]
		if !present {
			continue
		}

		if deps&^satisfied == 0 {
			delete(todo.tasks, id)
			return id, true
		}
	}

	return -1, false
}

func DoPart1(inputs []Depend) AdventResult {
	todo := buildTodo(inputs)

	order := make([]byte, 0, len(todo.ids))
	satisfied := uint(0)

	for len(order) < len(todo.ids) {
		id, ok := todo.next(satisfied)
		if !ok {
			panic("no task satisfied")
		}

		satisfied |= 1 << id
		order = append(order, byte('A'+id))
	}

	return AdventResult(order)
}

func DoPart2(input []Depend, nworkers, baseCost int) int {
	todo := buildTodo(input)

	order := make([]byte, 0, len(todo.ids))
	satisfied := uint(0)

	const noTask = -1
	type job struct {
		id, timeLeft int
	}
	jobs := make([]job, nworkers)
	for i := 0; i < nworkers; i++ {
		jobs[i].id = noTask
	}

	elapsed := 0
	for len(order) < len(todo.ids) {
		step := int(^uint(0) >> 1)

		// assign tasks
		for i := range jobs {
			if jobs[i].id == noTask {
				id, ok := todo.next(satisfied)
				if !ok {
					continue
				}
				jobs[i] = job{id, id + baseCost}
			}
			if jobs[i].timeLeft < step {
				step = jobs[i].timeLeft
			}
		}

		// advance clock, complete jobs
		for i := range jobs {
			if jobs[i].id == noTask {
				continue
			}

			jobs[i].timeLeft -= step
			if jobs[i].timeLeft > 0 {
				continue
			}

			satisfied |= 1 << jobs[i].id
			order = append(order, byte('A'+jobs[i].id))
			jobs[i].id = noTask
		}
		elapsed += step
	}

	return elapsed
}

func Part1() AdventResult {
	return DoPart1(ParseInput(openInput()))
}

func Part2() int {
	return DoPart2(ParseInput(openInput()), 5, 61)
}
