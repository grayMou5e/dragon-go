package main

import (
	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
	"github.com/grayMou5e/dragon-go/weather"
)

//playGame inits game flow
func playGame(handler handlers.GameHandler) *game.Data {
	game := startGame(handler)
	getWeather(handler, game)
	addDragon(game)

	result := handler.FightAgainstTheKnight(&game.Dragon, game.GameID)
	game.Result = *result
	game.Result.Summarize()

	return game
}

//startGame receives game from 3rd party
func startGame(handler handlers.GameHandler) *game.Data {
	gameData := *handler.GetGame()

	return &gameData
}

//getWeather receives weather information from 3rd party
func getWeather(handler handlers.GameHandler, gameData *game.Data) {
	weatherData := handler.GetWeather(gameData.GameID)
	weather.AddType(weatherData)
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
