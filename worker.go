package main

import (
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
	uuid "github.com/nu7hatch/gouuid"
	"go.uber.org/zap"
)

func worker(handler *handlers.GameHandler, jobs <-chan int, results chan<- *game.Data, logger *zap.Logger) {
	for _ = range jobs {
		guid, guidGenErr := uuid.NewV4()
		if guidGenErr != nil {
			continue
		}
		// game, gameError := playGame(handler)
		game, gameError := timedPlayGame(handler, logger, guid)
		if gameError != nil {
			//log
			continue
		}
		results <- game
	}
}

func createWorkers(quantity int, handler *handlers.GameHandler, logger *zap.Logger, jobs <-chan int, results chan<- *game.Data) {
	for w := 1; w <= quantity; w++ {
		go worker(handler, jobs, results, logger)
	}
}

func queJobs(quantity int, jobs chan<- int) {
	for j := 1; j <= quantity; j++ {
		jobs <- j
	}
}
