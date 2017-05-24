package main

import (
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
	uuid "github.com/nu7hatch/gouuid"
)

func worker(handler *handlers.GameHandler, jobs <-chan int, results chan<- *game.Data) {
	for _ = range jobs {
		//generate corelation id !
		guid, guidGenErr := uuid.NewV4()
		if guidGenErr != nil {
			continue
		}
		game, gameError := playGame(handler, guid)

		if gameError != nil {
			//log
			continue
		}
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
