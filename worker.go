package main

import (
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
)

func worker(handler *handlers.GameHandler, jobs <-chan int, results chan<- *game.Data) {
	for _ = range jobs {
		game := playGame(handler)
		results <- game
	}
}

func createWorkers(quantity int, handler *handlers.GameHandler, jobs <-chan int, results chan<- *game.Data) {
	for w := 1; w <= quantity; w++ {
		go worker(handler, jobs, results)
	}
}

func queJobs(quantity int, jobs chan<- int) {
	for j := 1; j <= quantity; j++ {
		jobs <- j
	}
}
