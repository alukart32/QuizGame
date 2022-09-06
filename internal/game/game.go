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

type problem struct {
	q string
	a string
}

// quizgame provides basic properties for any quiz game
type quizgame struct {
	Problems []problem // Questions that will be asked
	ExitWord string    // Word to exit game.
}

func (*quizgame) About() string {
	return `
*** Quiz Game ***
Rules:
	1. give numerical answers to the questions
	2. to stop the game type 'stop'
=========================================
`
}

func (g *quizgame) Run() error {
	score := 0
	scanner := bufio.NewScanner(os.Stdin)
	for i, p := range g.Problems {
		fmt.Printf("Question %d. %s = ", i+1, p.q)
		if scanner.Scan() {
			line := scanner.Text()
			if err := scanner.Err(); err != nil {
				return err
			}
			if line == g.ExitWord {
				break
			}
			if ok, err := processAnswer(&p, line); err != nil {
				fmt.Println(err)
			} else if ok {
				score++
			}
		}
	}
	fmt.Printf("\nYou %d get of %d scores\n", score, len(g.Problems))
	return nil
}

func processAnswer(p *problem, answer string) (bool, error) {
	if _, e := strconv.Atoi(answer); e != nil {
		return false, fmt.Errorf("invalid answer format - must be numeric input")
	}
	return p.a == answer, nil
}
