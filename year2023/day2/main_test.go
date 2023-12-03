package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ExampleOne string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	ExampleOneAnswerA int = 8
	ExampleOneAnswerB int = 2286
)

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
