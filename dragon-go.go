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
	amountOfGames := 100
	handler := handlers.NewHandler()
	var hndlr handlers.GameHandler
	hndlr = handler
	wins := 0
	startTime := time.Now()
	for i := 0; i < amountOfGames; i++ {
		game := playGame(hndlr)
		if game.Result.Victory {
			wins++
		}
	}
	elapsed := time.Since(startTime)
	fmt.Println(fmt.Sprintf("Won - %d Lost - %d", wins, amountOfGames-wins))
	fmt.Printf("Time elapsed %s", elapsed)
}

func playGame(handler handlers.GameHandler) *game.Data {

	game := startGame(handler)
	getWeather(handler, game)
	addDragon(game)

	result := handler.FightAgainstTheKnight(&game.Dragon, game.GameID)
	game.Result = *result
	game.Result.Summarize()

	// fmt.Println(result)

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
