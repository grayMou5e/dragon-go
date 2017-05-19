package main

import (
	"bufio"
	"fmt"
	"os"

	"time"

	"strconv"

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
