// Game factory
package game

import (
	"fmt"
	"strings"

	"alukart32.com/quizgame/internal/resource"
)

const (
	CsvQuizGame = iota
)

const (
	GAME_RULES_CSV_FILE_PATH = "../assets/game-rules-csv.txt"
)

var (
	csvReader = resource.GetCsvRuleReader(';', '#')
)

func GetGame(t int, problemsFile string) (Game, error) {
	fmt.Println("loading game rules ...")
	switch t {
	case CsvQuizGame:
		lines, e := csvReader.ReadAll(problemsFile)
		if e != nil {
			panic(e)
		}
		gameOjct := &quizgame{
			Problems: convertCSVLines(lines),
			ExitWord: "stop",
		}
		return gameOjct, nil
	default:
		return nil, fmt.Errorf("not implemented yet")
	}
}

func convertCSVLines(lines [][]string) []problem {
	if len(lines) == 0 {
		return nil
	}

	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
