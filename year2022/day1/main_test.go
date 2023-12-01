package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	EXAMPLE string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	EXAMPLE_ANSWER_A int = 24000
	EXAMPLE_ANSWER_B int = 45000
)

func TestParseInputToGroups(t *testing.T) {
	parsedGroups, _ := parseInputToGroups(EXAMPLE)
	assert.Equal(t, 5, len(parsedGroups), "Unexpected number of parsed groups")
	assert.ElementsMatch(t, []int{1000, 2000, 3000}, parsedGroups[0], "Elements in first parsed group were unexpected")
}

func TestAnswers(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		answerA int
		answerB int
	}{
		{
			name:    "part a example answer",
			input:   EXAMPLE,
			answerA: EXAMPLE_ANSWER_A,
			answerB: EXAMPLE_ANSWER_B,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			answerA, answerB := answers(test.input)
			assert.Equal(t, answerA, test.answerA, "Answer A incorrect")
			assert.Equal(t, answerB, test.answerB, "Answer B incorrect")
		})
	}
}
