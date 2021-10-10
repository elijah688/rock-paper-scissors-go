package main

import (
	"rps/util"
	"sync"
)

func main() {
	var playerChannel chan string = make(chan string)
	var endChan chan bool = make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go util.HumanInput(playerChannel, &wg, endChan)
	wg.Add(1)
	go util.Game(playerChannel, &wg, endChan)
	wg.Wait()
}
