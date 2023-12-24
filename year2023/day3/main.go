// Solve Advent of Code, 2023, day 3
package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

type Coordinate struct {
	x int
	y int
}

type Schematic struct {
	numbers map[Coordinate]int
	symbols map[Coordinate]rune
}

func NewSchematic() *Schematic {
	var s Schematic
	s.numbers = make(map[Coordinate]int)
	s.symbols = make(map[Coordinate]rune)

	return &s
}

func (s *Schematic) answerA() int {
	answerA := 0
	for cn, n := range s.numbers {
		log.Debug().Msgf("Looking at [%d] (%d,%d)", n, cn.x, cn.y)
		// Look at all the coordinates from cn.x-1 to cn.x+len(n)
	symbol_search:
		for _, x := range makeRange(cn.x-1, cn.x+len(strconv.Itoa(n))) {
			for _, y := range makeRange(cn.y-1, cn.y+1) {
				log.Debug().Msgf("Checking for symbol at (%d,%d)", x, y)
				symbol, ok := s.symbols[Coordinate{x, y}]
				if ok {
					log.Debug().Msgf("Found symbol %q at (%d,%d)", symbol, x, y)
					answerA += n
					break symbol_search
				}
			}
		}
	}
	return answerA
}

func (s *Schematic) answerB() int {
	return 0
}

func parseSchematic(in string) Schematic {
	inputStrings := strings.Split(strings.TrimSpace(in), "\n")
	s := NewSchematic()

	for y, row := range inputStrings {
		// Look for numbers with regexp
		reDigit := regexp.MustCompile(`\d+`)
		for _, ds := range reDigit.FindAllIndex([]byte(row), -1) {
			s.numbers[Coordinate{ds[0], y}], _ = strconv.Atoi(row[ds[0]:ds[1]])
		}

		// Look for symbols with regexp
		reSymbol := regexp.MustCompile(`[^\d\.]`)
		for _, ss := range reSymbol.FindAllIndex([]byte(row), -1) {
			s.symbols[Coordinate{ss[0], y}], _, _, _ = strconv.UnquoteChar(row[ss[0]:ss[1]], '"')
		}
	}
	return *s
}

func answers(in string) (int, int) {
	s := parseSchematic(in)
	return s.answerA(), s.answerB()
}

//go:embed input.txt
var aocInput string

func main() {
	answerA, answerB := answers(aocInput)
	fmt.Printf("Answer A: %d\n", answerA)
	fmt.Printf("Answer B: %d\n", answerB)
}
