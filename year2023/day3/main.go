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
	type PossibleGear struct {
		numbers map[Coordinate]int
	}

	const GearLimit = 2
	answerB := 0
	possibleGears := make(map[Coordinate]PossibleGear)

	for cn, n := range s.numbers {
		for _, x := range makeRange(cn.x-1, cn.x+len(strconv.Itoa(n))) {
			for _, y := range makeRange(cn.y-1, cn.y+1) {
				log.Debug().Msgf("Checking for symbol at (%d,%d)", x, y)
				if symbol, ok := s.symbols[Coordinate{x, y}]; ok && symbol == '*' {
					log.Debug().Msgf("Found possible gear at (%d,%d)", x, y)
					if gear, ok := possibleGears[Coordinate{x, y}]; ok {
						log.Debug().Msgf("Adding [%d] to existing gear at (%d,%d)", n, x, y)
						gear.numbers[cn] = n
					} else {
						log.Debug().Msgf("Adding [%d] to new gear at (%d,%d)", n, x, y)
						possibleGears[Coordinate{x, y}] = PossibleGear{
							numbers: map[Coordinate]int{cn: n},
						}
					}
				}
			}
		}
	}

	log.Debug().Msgf("Found %d possible gears...", len(possibleGears))

	for c, g := range possibleGears {
		numbers := g.numbers
		if len(numbers) == GearLimit {
			log.Debug().Msgf("Found gear with two numbers at (%d, %d): %v", c.x, c.y, numbers)
			gearRatio := 1
			for _, v := range numbers {
				gearRatio *= v
			}
			answerB += gearRatio
		} else {
			log.Debug().Msgf("Found gear with lt or gt two numbers at (%d, %d): %v", c.x, c.y, numbers)
		}
	}

	return answerB
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
