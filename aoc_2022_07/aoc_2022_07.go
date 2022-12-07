package aoc_2022_07

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result int

const Part1Want = 1_743_217

type Part2Result int

const Part2Want = 8_319_096

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) []*File {
	scanner := bufio.NewScanner(input)

	fs := make([]*File, 0)

	root := new(File)
	root.name = "/"

	fs = append(fs, root)

	cur := root

Line:
	for scanner.Scan() {
		line := scanner.Text()

		if line == "$ cd /" {
			cur = root
			continue
		}

		if line == "$ ls" {
			continue
		}

		if line == "$ cd .." {
			if cur != root {
				cur = cur.parent
			}
			continue
		}

		var name string

		_, err := fmt.Sscanf(line, "$ cd %s", &name)
		if err == nil {
			for _, c := range cur.children {
				if c.name == name {
					cur = c
					continue Line
				}
			}
			panic(fmt.Sprintf("dir not found: %s", name))
		}

		_, err = fmt.Sscanf(line, "dir %s", &name)
		if err == nil {
			f := new(File)
			fs = append(fs, f)
			f.name = name
			f.parent = cur
			cur.children = append(cur.children, f)
			continue Line
		}

		f := new(File)
		fs = append(fs, f)
		f.parent = cur
		cur.children = append(cur.children, f)
		_, err = fmt.Sscanf(line, "%d %s", &f.size, &f.name)
		if err != nil {
			panic(fmt.Sprintf("unrecognized line: %s", line))
		}
	}

	return fs
}

type File struct {
	parent   *File
	name     string
	size     int
	children []*File
}

func (f *File) computeSize() int {
	if f.size > 0 {
		return f.size
	}

	size := 0
	for _, c := range f.children {
		size += c.computeSize()
	}
	f.size = size
	return size
}

func (f *File) isDir() bool {
	return 0 != len(f.children)
}

func DoPart1(fs []*File) Part1Result {
	const maxDirSize = 100000

	fs[0].computeSize()

	sum := 0
	for _, f := range fs {
		if f.isDir() && f.size < maxDirSize {
			sum += f.size
		}
	}

	return Part1Result(sum)
}

func DoPart2(fs []*File) Part2Result {
	const totalSpace = 70000000
	const needSpace = 30000000

	fs[0].computeSize()

	freeSpace := totalSpace - fs[0].size
	minToFree := needSpace - freeSpace

	best := fs[0]
	for _, f := range fs {
		if f.isDir() && f.size < best.size && f.size >= minToFree {
			best = f
		}
	}

	return Part2Result(best.size)
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput(openInput()))
}
