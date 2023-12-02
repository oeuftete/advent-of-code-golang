// Solve Advent of Code, 2023, day 1
package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var digitWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type notFoundError struct {
	input string
	msg   string
}

func (e *notFoundError) Error() string {
	return fmt.Sprintf("%s: [%s]", e.msg, e.input)
}

func nConvert(s string) int {
	var n int
	var ok bool
	digitMap := make(map[string]int)

	for idx, w := range digitWords {
		digitMap[w] = idx + 1
	}

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
			re = regexp.MustCompile(`\d|` + strings.Join(digitWords, "|"))
		} else {
			re = regexp.MustCompile(`\d`)
		}
		firstString := re.Find([]byte(inputString))
		if firstString == nil {
			return nil, &notFoundError{
				input: inputString,
				msg:   "No first string found",
			}
		}
		first = nConvert(string(firstString))

		// Find last
		var lastString string
		maxIndex := -1

		digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		if useWords {
			digits = append(digits, digitWords...)
		}

		for _, d := range digits {
			i := strings.LastIndex(inputString, d)
			if i > maxIndex {
				maxIndex = i
				lastString = d
			}
		}
		if lastString == "" {
			return nil, &notFoundError{
				input: inputString,
				msg:   "No last string found",
			}
		}
		last = nConvert(lastString)

		log.Info().Msgf("  Parsed first and last: (%d, %d)", first, last)
		parsedInts[idx] = int(math.Pow10(1))*first + last
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
