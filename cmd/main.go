package main

import (
	"flag"
	"fmt"
	"log"

	"alukart32.com/quizgame/internal/game"
)

func main() {
	problemsFile := flag.String("filepath", "../assets/problems.csv", "a file in the format 'question:answer'")
	flag.Parse()

	g, err := game.GetGame(game.CsvQuizGame, *problemsFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.About())
	if err = g.Run(); err != nil {
		log.Fatal(err)
	}
}
