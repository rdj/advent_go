package aoc_2022_25

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

type Part1Result string

const Part1Want = "2-0==21--=0==2201==2"

type ParsedInput []string

func openInput() io.Reader {
	reader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return reader
}

func ParseInput(input io.Reader) ParsedInput {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func DoPart1(input ParsedInput) Part1Result {
	total := 0
	for _, line := range input {
		n := 0
		for _, r := range line {
			n *= 5
			switch r {
			case '=':
				n -= 2
			case '-':
				n -= 1
			case '0':
				//
			case '1':
				n += 1
			case '2':
				n += 2
			}
		}
		//fmt.Println(line, n)
		total += n
	}

	//fmt.Println("Total", total)
	digits := []int{}
	for total > 0 {
		digits = append(digits, total%5)
		total /= 5
	}
	//fmt.Println(digits)
	runes := []rune{}
	for i := range digits {
		r := '0'
		switch digits[i] {
		case 5:
			r = '0'
			digits[i+1]++
		case 4:
			r = '-'
			digits[i+1]++
		case 3:
			r = '='
			digits[i+1]++
		default:
			r += rune(digits[i])
		}
		runes = append(runes, r)
	}
	//fmt.Println(string(runes))
	sb := new(strings.Builder)
	for i := range runes {
		sb.WriteRune(runes[len(runes)-1-i])
	}
	return Part1Result(sb.String())
}

func Part1() Part1Result {
	return DoPart1(ParseInput(openInput()))
}
