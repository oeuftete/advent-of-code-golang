package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ExampleOne string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	ExampleTwo string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	ExampleOneAnswer int = 142
	ExampleTwoAnswer int = 281
)

func TestParseInputToGroups(t *testing.T) {
	parsedInput, err := parseInputToNumbers(ExampleOne, false)
	if assert.NoError(t, err, "input parsing failed") {
		assert.ElementsMatch(t, []int{12, 38, 15, 77}, parsedInput, "Elements in parsed input were unexpected")
	}
}

func TestAnswers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		answer   int
		useWords bool
	}{
		{
			name:     "example one answer",
			input:    ExampleOne,
			answer:   ExampleOneAnswer,
			useWords: false,
		},
		{
			name:     "example two answer",
			input:    ExampleTwo,
			answer:   ExampleTwoAnswer,
			useWords: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var answer int
			answerA, answerB := answers(test.input)
			if test.useWords {
				answer = answerB
			} else {
				answer = answerA
			}
			assert.Equal(t, answer, test.answer, "Answer A incorrect")
		})
	}
}
