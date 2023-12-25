// Solve Advent of Code, 2023, day 4
package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
)

type GameCard struct {
	id      int
	winners map[int]bool
	mine    map[int]bool
}

func NewGameCard() GameCard {
	var gc GameCard
	gc.winners = make(map[int]bool)
	gc.mine = make(map[int]bool)
	return gc
}

func parseCards(in string) (gameCards []GameCard) {
	inputStrings := strings.Split(strings.TrimSpace(in), "\n")
	reCard := regexp.MustCompile(`^Card\s*(\d+):`)
	reWS := regexp.MustCompile(`\s+`)

	for _, row := range inputStrings {
		gc := NewGameCard()

		reCardMatches := reCard.FindSubmatch([]byte(row))
		gc.id, _ = strconv.Atoi(string(reCardMatches[1]))

		_, numberString, _ := strings.Cut(row, ": ")
		winnersString, mineString, _ := strings.Cut(numberString, " | ")

		for _, w := range reWS.Split(strings.TrimSpace(winnersString), -1) {
			iw, _ := strconv.Atoi(w)
			gc.winners[iw] = true
		}

		for _, m := range reWS.Split(strings.TrimSpace(mineString), -1) {
			im, _ := strconv.Atoi(m)
			gc.mine[im] = true
		}
		gameCards = append(gameCards, gc)
	}
	return gameCards
}

func answers(in string) (answerA, answerB int) {
	gameCards := parseCards(in)

	answerA = 0

	for _, gameCard := range gameCards {
		cardScore := uint(0)

		mine := maps.Keys(gameCard.mine)
		sort.Ints(mine)

		winners := maps.Keys(gameCard.winners)
		sort.Ints(winners)

		for _, n := range mine {
			if slices.Contains(winners, n) {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore <<= 1
				}
			}
		}
		answerA += int(cardScore)
		log.Debug().Int("answer", answerA).Msg("Current running score")
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
