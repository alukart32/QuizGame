package main

import (
	"fmt"

	"alukart32.com/quizgame/internal/game"
)

func main() {
	g, e := game.GetGame(game.CsvQuizGame)
	if e != nil {
		panic(e)
	}
	fmt.Println(g.About())
	g.Run()
}
