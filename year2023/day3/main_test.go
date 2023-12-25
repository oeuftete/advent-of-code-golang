package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ExampleOne string = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	ExampleOneAnswerA int = 4361
	ExampleOneAnswerB int = 467835
)

func TestParseSchematic(t *testing.T) {
	s := parseSchematic(ExampleOne)
	assert.Equal(t, 467, s.numbers[Coordinate{0, 0}])
	assert.Equal(t, 114, s.numbers[Coordinate{5, 0}])
	assert.Equal(t, 664, s.numbers[Coordinate{1, 9}])

	assert.Equal(t, string(`*`), string(s.symbols[Coordinate{3, 1}]))
	assert.Equal(t, string(`#`), string(s.symbols[Coordinate{6, 3}]))
	assert.Equal(t, string(`+`), string(s.symbols[Coordinate{5, 5}]))
	assert.Equal(t, string(`$`), string(s.symbols[Coordinate{3, 8}]))
}

func TestAnswers(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		answerA int
		answerB int
	}{
		{
			name:    "example one answer",
			input:   ExampleOne,
			answerA: ExampleOneAnswerA,
			answerB: ExampleOneAnswerB,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			answerA, answerB := answers(test.input)
			assert.Equal(t, test.answerA, answerA, "Answer A incorrect")
			assert.Equal(t, test.answerB, answerB, "Answer B incorrect")
		})
	}
}
