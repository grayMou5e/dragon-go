package main

import (
	"fmt"

	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
	"github.com/grayMou5e/dragon-go/weather"
)

func main() {
	handler := handlers.NewHandler()
	var hndlr handlers.GameHandler
	hndlr = handler

	playGame(hndlr)
}

func playGame(handler handlers.GameHandler) *game.Data {

	game := startGame(handler)
	getWeather(handler, game)
	addDragon(game)
	fmt.Println(game)
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
	gameData.Dragon = *dragon.CreateDragon(gameData.Knight.Attack, gameData.Knight.Armor, gameData.Knight.Agility, gameData.Knight.Endurance)
}
