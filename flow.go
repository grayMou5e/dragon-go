package main

import (
	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
)

//playGame inits game flow
func playGame(handler *handlers.GameHandler) *game.Data {
	game := startGame(handler)
	setWeather(handler, game)
	addDragon(game)

	game.Result = *(*handler).FightAgainstTheKnight(&game.Dragon, game.GameID)
	game.Result.Summarize()

	return game
}

//startGame receives game from 3rd party
func startGame(handler *handlers.GameHandler) *game.Data {
	gameData := *(*handler).GetGame()

	return &gameData
}

//getWeather receives weather information from 3rd party
func setWeather(handler *handlers.GameHandler, gameData *game.Data) {
	weatherData := (*handler).GetWeather(gameData.GameID)
	weatherData.AddType()
	gameData.Weather = *weatherData
}

//addDragon generates dragon by using knight data & assigns it to game data
func addDragon(gameData *game.Data) {
	gameData.Dragon = *dragon.CreateDragon(gameData.Knight.Attack,
		gameData.Knight.Armor,
		gameData.Knight.Agility,
		gameData.Knight.Endurance,
		gameData.Weather.Type)
}
