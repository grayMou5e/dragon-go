package main

import (
	"fmt"

	"time"

	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
)

func main() {
	amountOfGames := 100

	jobs := make(chan int, amountOfGames)
	results := make(chan *game.Data, amountOfGames)

	handler := handlers.NewHandler()
	var hndlr handlers.GameHandler
	hndlr = handler

	createWorkers(100, &hndlr, jobs, results)

	startTime := time.Now()
	queJobs(amountOfGames, jobs)
	close(jobs)

	wins := processPlayedGames(amountOfGames, results)
	close(results)
	elapsed := time.Since(startTime)

	fmt.Println(fmt.Sprintf("Won - %d Lost - %d", wins, amountOfGames-wins))
	fmt.Printf("Time elapsed %s", elapsed)
}

func processPlayedGames(quantity int, results <-chan *game.Data) (gamesWon int) {
	for a := 1; a <= quantity; a++ {
		data := <-results
		if data.Result.Victory {
			gamesWon++
		}
	}

	return gamesWon
}
