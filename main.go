package main

import (
	"rps/util"
	"sync"
)

func main() {
	var playerChannel chan int = make(chan int)
	var endChan chan bool = make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go util.HumanInput(playerChannel, &wg, endChan)
	wg.Add(1)
	go util.Game(playerChannel, &wg, endChan)
	wg.Wait()
}
