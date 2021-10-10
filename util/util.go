package util

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const (
	ROCK     = "rock"
	PAPER    = "paper"
	SCISSORS = "scissors"
)

var rps []string = []string{"rock", "paper", "scissors"}

func initialPrompt() {
	fmt.Println("[ROCK, PAPER, SCISORS]")
	fmt.Print("-> ")
}

func clearScreen() {
	time.Sleep(1 * time.Second)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
func validHumanInput(input string) bool {
	for _, legalMove := range rps {
		if legalMove == input {
			return true
		}
	}
	return false
}

func CpuMove() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var cpuMove int = rand.Intn(3)
	return rps[cpuMove]
}

func HumanInput(
	playerInChan chan string,
	wg *sync.WaitGroup,
	endChan chan bool,
) {
	defer wg.Done()
	for {
		var humanMove string
		fmt.Scanln(&humanMove)
		humanMove = strings.ToLower(humanMove)
		if humanMove == "q" {
			endChan <- true
			return
		} else if validHumanInput(humanMove) {
			playerInChan <- humanMove
		} else {
			fmt.Println("Illegal input! Try Again")
			clearScreen()
			initialPrompt()
		}
	}
}

func youWin() {
	fmt.Println("\n==============")
	fmt.Println("You Win!!")
	fmt.Println("==============")
}

func cpuWins() {
	fmt.Println("\n==============")
	fmt.Println("CPU Wins!!")
	fmt.Println("==============")
}

func draw() {
	fmt.Println("\n==============")
	fmt.Println("Draw!!")
	fmt.Println("==============")
}

func produceWinner(humanInput, cpuInput string) {
	switch humanInput {
	case cpuInput:
		draw()
	case ROCK:
		if cpuInput == SCISSORS {
			youWin()
		} else {
			cpuWins()
		}
	case SCISSORS:
		if cpuInput == PAPER {
			youWin()
		} else {
			cpuWins()
		}
	case PAPER:
		if cpuInput == ROCK {
			youWin()
		} else {
			cpuWins()
		}
	}
}

func Game(
	playerInChan chan string,
	wg *sync.WaitGroup,
	endChan chan bool,
) {
	defer wg.Done()
	for {
		initialPrompt()
		select {
		case playerMove := <-playerInChan:
			var cpuMove string = CpuMove()
			fmt.Println("You Chose: ", playerMove)
			fmt.Println("CPU Chose: ", cpuMove)
			produceWinner(playerMove, cpuMove)
			clearScreen()
		case <-endChan:
			return
		}
	}
}
