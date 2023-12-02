// Solve Advent of Code, 2023, day 2
package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type parseError struct {
	input string
	msg   string
}

func (e *parseError) Error() string {
	return fmt.Sprintf("%s: [%s]", e.msg, e.input)
}

type GameTurn struct {
	red   int
	green int
	blue  int
}

func (gt *GameTurn) isPossible(result GameTurn) bool {
	if gt.red <= result.red && gt.green <= result.green && gt.blue <= result.blue {
		return true
	}
	return false
}

type Game struct {
	id    int
	turns []GameTurn
}

func (g *Game) isPossible(contents GameTurn) bool {
	for _, turn := range g.turns {
		if !turn.isPossible(contents) {
			return false
		}
	}
	return true
}

func (g *Game) power() int {
	var power int
	var minimumSet GameTurn

	for _, turn := range g.turns {
		if turn.red > minimumSet.red {
			minimumSet.red = turn.red
		}
		if turn.green > minimumSet.green {
			minimumSet.green = turn.green
		}
		if turn.blue > minimumSet.blue {
			minimumSet.blue = turn.blue
		}
	}
	power = minimumSet.red * minimumSet.green * minimumSet.blue
	return power
}

func parseInputToGames(in string) ([]Game, error) {
	inputStrings := strings.Split(strings.TrimSpace(in), "\n")
	parsedGames := make([]Game, len(inputStrings))

	for idx, inputString := range inputStrings {
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		var g Game
		var err error

		re := regexp.MustCompile(`Game (\d+): (.*)`)
		m := re.FindSubmatch([]byte(inputString))
		if m == nil {
			return nil, &parseError{
				input: inputString,
				msg:   "Game not found",
			}
		}

		g.id, err = strconv.Atoi(string(m[1]))
		if err != nil {
			return nil, err
		}

		for _, turnString := range strings.Split(string(m[2]), "; ") {
			var turn GameTurn
			for _, colorCount := range strings.Split(turnString, ", ") {
				sp := strings.Fields(colorCount)
				n, _ := strconv.Atoi(sp[0])
				color := sp[1]

				switch color {
				case "red":
					turn.red = n
				case "green":
					turn.green = n
				case "blue":
					turn.blue = n
				default:
					return nil, &parseError{
						input: turnString,
						msg:   fmt.Sprintf("Unexpected color [%s]", color),
					}
				}
			}
			g.turns = append(g.turns, turn)
		}
		parsedGames[idx] = g
	}
	return parsedGames, nil
}

func answers(in string) (int, int) {
	var answerA int
	var answerB int

	const (
		RedResult   int = 12
		GreenResult int = 13
		BlueResult  int = 14
	)

	games, err := parseInputToGames(in)
	if err != nil {
		log.Fatal().Msg("Problem parsing input!")
	}

	gameResult := &GameTurn{
		red:   RedResult,
		green: GreenResult,
		blue:  BlueResult,
	}
	for _, g := range games {
		if g.isPossible(*gameResult) {
			answerA += g.id
		}
		answerB += g.power()
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
