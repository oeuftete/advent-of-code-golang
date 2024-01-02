package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ExampleOne string = `
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
	ExampleOneAnswerA int = 35
	ExampleOneAnswerB int = 46
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
