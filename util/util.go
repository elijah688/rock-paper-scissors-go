package util

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
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
func handleHumanInput(input string) (int, error) {
	for index, legalMove := range rps {
		if legalMove == input {
			return index, nil
		}
	}
	return -1, errors.New("illegal input")
}

func CpuMove() int {
	rand.Seed(time.Now().UTC().UnixNano())
	var cpuMove int = rand.Intn(3)
	return cpuMove
}

func HumanInput(
	playerInChan chan int,
	wg *sync.WaitGroup,
	endChan chan bool,
) {
	defer wg.Done()
	for {
		var humanMove string
		fmt.Scanln(&humanMove)
		if humanMove == "q" {
			endChan <- true
			return
		} else if rpsValue, err := handleHumanInput(humanMove); err == nil {
			playerInChan <- rpsValue
		} else {
			fmt.Println(err)
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

func produceWinner(humanInput, cpuInput int) {
	if humanInput == cpuInput {
		draw()
	} else if (cpuInput+1)%3 == humanInput {
		youWin()
	} else {
		cpuWins()
	}
}
func Game(
	playerInChan chan int,
	wg *sync.WaitGroup,
	endChan chan bool,
) {
	defer wg.Done()
	for {
		initialPrompt()
		select {
		case playerMove := <-playerInChan:
			var cpuMove int = CpuMove()
			fmt.Println("You Chose: ", rps[playerMove])
			fmt.Println("CPU Chose: ", rps[cpuMove])
			produceWinner(playerMove, cpuMove)
			clearScreen()
		case <-endChan:
			return
		}
	}
}
