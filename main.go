package main

import (
	"bufio"
	"fmt"
	"os"

	"time"

	"strconv"

	"github.com/grayMou5e/dragon-go/config"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
)

func main() {
	amountOfGames := getAmountOfGames()

	jobs := make(chan int, amountOfGames)
	results := make(chan *game.Data, amountOfGames)

	handler := handlers.NewHandler()
	var hndlr handlers.GameHandler
	hndlr = handler

	logger, _ := config.NewLogger()
	defer logger.Sync()

	createWorkers(100, &hndlr, logger, jobs, results)

	startTime := time.Now()
	queJobs(amountOfGames, jobs)
	close(jobs)

	wins := processPlayedGames(amountOfGames, results)
	close(results)
	elapsed := time.Since(startTime)

	fmt.Printf("\nWon - %d Lost - %d", wins, amountOfGames-wins)
	fmt.Printf("\nTime elapsed %s\n", elapsed)
}

func getAmountOfGames() (amountOfGames int) {
	enterMessage := "Enter amount of games: "
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(enterMessage)
	for scanner.Scan() {
		var inputError error
		amountOfGames, inputError = strconv.Atoi(scanner.Text())

		if inputError == nil {
			break
		}

		fmt.Println(enterMessage)
	}

	return amountOfGames
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
