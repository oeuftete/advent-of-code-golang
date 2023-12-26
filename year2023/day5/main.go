package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	InitialSource string = "seed"
)

type ConversionMap struct {
	source      string
	destination string
	conversions []func(int) int
}

type Almanac struct {
	seeds          []int
	conversionMaps map[string]ConversionMap
}

func NewAlmanac() (a Almanac) {
	a.conversionMaps = make(map[string]ConversionMap)
	return a
}

func parseAlmanac(in string) Almanac {
	inputStrings := strings.Split(strings.TrimSpace(in), "\n")
	almanac := NewAlmanac()

	// Parse the seed line
	seedLine, inputStrings := strings.TrimSpace(inputStrings[0]), inputStrings[2:]
	reWS := regexp.MustCompile(`\s+`)
	_, seeds, _ := strings.Cut(seedLine, ": ")

	for _, s := range reWS.Split(seeds, -1) {
		seed, _ := strconv.Atoi(s)
		almanac.seeds = append(almanac.seeds, seed)
	}

	reMapStart := regexp.MustCompile(`(\w+)-to-(\w+) map:`)
	reMapRules := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)
	var lastSource string
	for _, row := range inputStrings {
		row := strings.TrimSpace(row)
		if row == "" {
			continue
		}
		if mapSubmatches := reMapStart.FindAllSubmatch([]byte(row), -1); mapSubmatches != nil {
			var cm ConversionMap
			cm.source = string(mapSubmatches[0][1])
			cm.destination = string(mapSubmatches[0][2])
			almanac.conversionMaps[cm.source] = cm
			lastSource = cm.source
			continue
		}
		if mapSubmatches := reMapRules.FindAllSubmatch([]byte(row), -1); mapSubmatches != nil {
			destRangeStart, _ := strconv.Atoi(string(mapSubmatches[0][1]))
			sourceRangeStart, _ := strconv.Atoi(string(mapSubmatches[0][2]))
			rangeLength, _ := strconv.Atoi(string(mapSubmatches[0][3]))
			f := func(i int) int {
				r := -1
				if i >= sourceRangeStart && i < (sourceRangeStart+rangeLength) {
					r = destRangeStart + (i - sourceRangeStart)
				}
				return r
			}
			conversionMap := almanac.conversionMaps[lastSource]
			conversionMap.conversions = append(conversionMap.conversions, f)
			almanac.conversionMaps[lastSource] = conversionMap
		}
	}
	return almanac
}

func (a *Almanac) answerA() int {
	locations := make([]int, 0, len(a.seeds))
	source := InitialSource

	for _, n := range a.seeds {
		log.Debug().Msgf("Generating value for seed %d", n)
		currentN := n
		for source != "" {
			// Look up the conversion map for the current source.  If there isn't
			// one, we must be at the end of the list.
			cm, ok := a.conversionMaps[source]
			if !ok {
				source = InitialSource
				break
			}

			// Apply the functions for each potential conversionMap
			//   - If there's a hit:
			//     - set currentN
			//     - break
			for _, f := range cm.conversions {
				if conversion := f(currentN); conversion != -1 {
					log.Debug().Msgf("Converted [%s:%d] to [%s:%d]", source, currentN, cm.destination, conversion)
					currentN = conversion
					break
				}
			}

			// No hit, leave currentN unchanged, move to next
			source = cm.destination
		}
		locations = append(locations, currentN)
	}
	return slices.Min(locations)
}

//go:embed input.txt
var aocInput string

func answers(in string) (answerA, answerB int) {
	almanac := parseAlmanac(in)

	answerA = almanac.answerA()
	answerB = 0
	return answerA, answerB
}

func main() {
	answerA, answerB := answers(aocInput)
	fmt.Printf("Answer A: %d\n", answerA)
	fmt.Printf("Answer B: %d\n", answerB)
}
