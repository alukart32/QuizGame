// Game factory
package game

import (
	"fmt"

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

func GetGame(t int) (Game, error) {
	fmt.Println("loading game rules ...")
	switch t {
	case CsvQuizGame:
		data, e := csvReader.ReadAll(GAME_RULES_CSV_FILE_PATH)
		if e != nil {
			panic(e)
		}
		gameOjct := &quizgame{
			Questions: data,
			ExitWord:  "stop",
		}
		return gameOjct, nil
	default:
		return nil, fmt.Errorf("not implemented yet")
	}
}
