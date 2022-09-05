package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Game is the interface that declares methods for any game.
type Game interface {
	About() string
	Run() error
}

// quizgame provides basic properties for any quiz game
type quizgame struct {
	Questions [][]string // Questions that will be asked
	ExitWord  string     // Word to exit game.
}

func (*quizgame) About() string {
	return `
*** Quiz Game ***
Rules:
	1. give your numerical answers to the questions
	2. to stop the game type 'stop'
=========================================
`
}

func (g *quizgame) Run() error {
	score := 0
	scanner := bufio.NewScanner(os.Stdin)
	for i := range g.Questions {
		fmt.Printf("Question %d. %s = ", i+1, g.Questions[i][0])
		if scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				break
			}
			if line == g.ExitWord {
				break
			}
			if ok, e := processAnswer(g.Questions[i][1], line); e != nil {
				fmt.Println(e)
			} else if ok {
				score++
			}
		}
	}
	fmt.Printf("\nYou %d get of %d scores\n", score, len(g.Questions))
	return nil
}

func processAnswer(excpected, answer string) (bool, error) {
	if _, e := strconv.Atoi(answer); e != nil {
		return false, fmt.Errorf("invalid answer format - must be numeric input")
	}
	return excpected == answer, nil
}
