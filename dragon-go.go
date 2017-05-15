package main

import (
	"fmt"

	"time"

	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
	"github.com/grayMou5e/dragon-go/weather"
)

func main() {
	amountOfGames := 5000

	jobs := make(chan int, amountOfGames)
	results := make(chan *game.Data, amountOfGames)

	handler := handlers.NewHandler()
	var hndlr handlers.GameHandler
	hndlr = handler

	for w := 1; w <= 100; w++ {
		go worker(&hndlr, jobs, results)
	}

	startTime := time.Now()
	for j := 1; j <= amountOfGames; j++ {
		jobs <- j
	}
	close(jobs)

	wins := 0
	for a := 1; a <= amountOfGames; a++ {
		data := <-results
		if data.Result.Victory {
			wins++
		}
	}
	close(results)
	elapsed := time.Since(startTime)

	fmt.Println(fmt.Sprintf("Won - %d Lost - %d", wins, amountOfGames-wins))
	fmt.Printf("Time elapsed %s", elapsed)
}

func worker(handler *handlers.GameHandler, jobs <-chan int, results chan<- *game.Data) {

	for _ = range jobs {
		game := playGame(*handler)
		results <- game
	}
}

func playGame(handler handlers.GameHandler) *game.Data {

	game := startGame(handler)
	getWeather(handler, game)
	addDragon(game)

	result := handler.FightAgainstTheKnight(&game.Dragon, game.GameID)
	game.Result = *result
	game.Result.Summarize()

	return game
}

func startGame(handler handlers.GameHandler) *game.Data {
	gameData := *handler.GetGame()

	return &gameData
}

func getWeather(handler handlers.GameHandler, gameData *game.Data) {
	weatherData := handler.GetWeather(gameData.GameID)
	weather.AddType(weatherData)
	gameData.Weather = *weatherData
}

func addDragon(gameData *game.Data) {
	gameData.Dragon = *dragon.CreateDragon(gameData.Knight.Attack,
		gameData.Knight.Armor,
		gameData.Knight.Agility,
		gameData.Knight.Endurance,
		gameData.Weather.Type)
}
