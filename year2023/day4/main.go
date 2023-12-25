// Solve Advent of Code, 2023, day 4
package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
)

// TODO: extract to helper.
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

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

func (gc *GameCard) nWinners() (nWinners int) {
	nWinners = 0

	mine := maps.Keys(gc.mine)
	winners := maps.Keys(gc.winners)

	for _, n := range mine {
		if slices.Contains(winners, n) {
			nWinners++
		}
	}
	return nWinners
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

	answerA = exponentialGame(gameCards)
	answerB = moreCardsGame(gameCards)
	return answerA, answerB
}

func moreCardsGame(gameCards []GameCard) (score int) {
	score = 0

	type GameCardSlot struct {
		gc     GameCard
		copies int
	}

	gameCardMap := make(map[int]GameCardSlot)

	for _, gameCard := range gameCards {
		gameCardMap[gameCard.id] = GameCardSlot{gameCard, 1}
	}

	for _, gameCard := range gameCards {
		nWinners := gameCard.nWinners()
		for i := 0; i < gameCardMap[gameCard.id].copies; i++ {
			for _, j := range makeRange(gameCard.id+1, gameCard.id+nWinners) {
				nCopies := gameCardMap[j].copies + 1
				gameCardMap[j] = GameCardSlot{gameCardMap[j].gc, nCopies}
			}
		}
	}

	for _, slot := range gameCardMap {
		score += slot.copies
	}

	return score
}

func exponentialGame(gameCards []GameCard) (score int) {
	score = 0

	for _, gameCard := range gameCards {
		cardScore := uint(0)
		nWinners := gameCard.nWinners()
		if nWinners > 0 {
			cardScore = 1 << (nWinners - 1)
		}
		score += int(cardScore)
		log.Debug().Int("answer", score).Msg("Current running score")
	}
	return score
}

//go:embed input.txt
var aocInput string

func main() {
	answerA, answerB := answers(aocInput)
	fmt.Printf("Answer A: %d\n", answerA)
	fmt.Printf("Answer B: %d\n", answerB)
}
