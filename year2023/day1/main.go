package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func nConvert(s string) int {
	var n int
	var ok bool

	if n, ok = digitMap[s]; ok {
		return n
	}

	n, _ = strconv.Atoi(s)
	return n
}

func parseInputToNumbers(in string, useWords bool) ([]int, error) {
	inputStrings := strings.Split(strings.TrimSpace(in), "\n")
	parsedInts := make([]int, len(inputStrings))

	for idx, inputString := range inputStrings {
		var first int
		var last int

		log.Info().Msgf("Input: %s", inputString)

		// Find first
		var re *regexp.Regexp
		if useWords {
			re = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
		} else {
			re = regexp.MustCompile(`\d`)
		}
		firstString := re.Find([]byte(inputString))
		first = nConvert(string(firstString))

		// Find last
		var lastString string
		maxIndex := -1

		digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		if useWords {
			digits = append(digits, "one", "two", "three", "four", "five", "six", "seven", "eight", "nine")
		}

		for _, d := range digits {
			i := strings.LastIndex(inputString, d)
			if i > maxIndex {
				maxIndex = i
				lastString = d
			}
		}
		last = nConvert(lastString)

		log.Info().Msgf("  Parsed first and last: (%d, %d)", first, last)
		parsedInts[idx] = 10*first + last
	}
	log.Info().Msgf("All values: %d", parsedInts)
	return parsedInts, nil
}

func answers(in string) (int, int) {
	var answerA int
	var answerB int
	var calibrationValues []int

	calibrationValues, _ = parseInputToNumbers(in, false)
	for _, i := range calibrationValues {
		answerA += i
	}

	calibrationValues, _ = parseInputToNumbers(in, true)
	for _, i := range calibrationValues {
		answerB += i
	}

	return answerA, answerB
}

//go:embed input.txt
var aocInput string

func main() {
	answerA, answerB := answers(aocInput)
	fmt.Printf("Answer A: %d\n", answerA)
	fmt.Printf("Answer B: %d\n", answerB)
}
