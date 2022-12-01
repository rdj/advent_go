package aoc_2018_14

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var _ = fmt.Println

const inputFile = "input.txt"

type Part1Result string

const Part1Fake = "0xDEAD_BEEF"
const Part1Want = "1132413111"

type Part2Result int

const Part2Fake = 0xDEAD_BEEF
const Part2Want = 20_340_232

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		panic("no input")
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func ParseInput2(input io.Reader) []byte {
	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		panic("no input")
	}
	text := scanner.Text()
	seq := make([]byte, len(text))
	for i := range seq {
		seq[i] = byte(text[i] - '0')
	}
	return seq
}

func bake(n int) string {
	a := []byte{3, 7}
	i := 0
	j := 1

	for len(a) < n+10 {
		sum := a[i] + a[j]
		tens := sum / 10
		if tens > 0 {
			a = append(a, tens)
		}
		ones := sum % 10
		a = append(a, ones)
		i = (i + int(a[i]) + 1) % len(a)
		j = (j + int(a[j]) + 1) % len(a)
	}

	var sb strings.Builder
	for _, b := range a[n : n+10] {
		fmt.Fprintf(&sb, "%d", b)
	}
	return sb.String()
}

func checkSeq(a []byte, seq []byte) (int, bool) {
	if len(a) < len(seq) {
		return 0, false
	}

	skipped := len(a) - len(seq)

	for i := range seq {
		if seq[i] != a[skipped+i] {
			return skipped, false
		}
	}
	return skipped, true
}

func bake2(seq []byte) int {
	a := []byte{3, 7}
	i := 0
	j := 1

	for n := 0; ; n++ {
		sum := a[i] + a[j]
		tens := sum / 10
		if tens > 0 {
			a = append(a, tens)
			if r, ok := checkSeq(a, seq); ok {
				return r
			}
		}
		ones := sum % 10
		a = append(a, ones)
		if r, ok := checkSeq(a, seq); ok {
			return r
		}

		i = (i + int(a[i]) + 1) % len(a)
		j = (j + int(a[j]) + 1) % len(a)
	}

}

func DoPart1(input int) Part1Result {
	return Part1Result(bake(input))
}

func DoPart2(input []byte) Part2Result {
	return Part2Result(bake2(input))
}

func Part1() Part1Result {
	return DoPart1(ParseInput1(openInput()))
}

func Part2() Part2Result {
	return DoPart2(ParseInput2(openInput()))
}
