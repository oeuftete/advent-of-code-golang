package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseInputToGroups(in string) ([][]int, error) {
	stringGroups := strings.Split(strings.TrimSpace(in), "\n\n")
	intGroups := make([][]int, len(stringGroups))

	for idx, stringGroup := range stringGroups {
		group := strings.Split(stringGroup, "\n")
		intGroup := make([]int, len(group))

		for gIdx, s := range group {
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			intGroup[gIdx] = i
		}
		intGroups[idx] = intGroup
	}
	return intGroups, nil
}

func answers(in string) (int, int) {
	intGroups, _ := parseInputToGroups(in)
	intGroupSums := make([]int, len(intGroups))

	for idx, intGroup := range intGroups {
		var sum int
		for _, calorieCount := range intGroup {
			sum += calorieCount
		}
		intGroupSums[idx] = sum
	}

	slices.Sort(intGroupSums)
	slices.Reverse(intGroupSums)

	answerA := intGroupSums[0]
	answerB := intGroupSums[0] + intGroupSums[1] + intGroupSums[2]
	return answerA, answerB
}

//go:embed input.txt
var aocInput string

func main() {
	answerA, answerB := answers(aocInput)
	fmt.Printf("Answer A: %d\n", answerA)
	fmt.Printf("Answer B: %d\n", answerB)
}
