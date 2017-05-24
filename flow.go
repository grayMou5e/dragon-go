package main

import (
	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
)

//playGame inits game flow
func playGame(handler *handlers.GameHandler) (*game.Data, error) {
	game, gameErr := startGame(handler)
	if gameErr != nil {
		return nil, gameErr
	}

	weatherErr := setWeather(handler, game)
	if weatherErr != nil {
		//log
	}

	addDragon(game)

	gameResult, fightError := (*handler).FightAgainstTheKnight(&game.Dragon, game.GameID)
	if fightError != nil {
		return nil, fightError
	}
	game.Result = *gameResult
	game.Result.Summarize()

	return game, nil
}

//startGame receives game from 3rd party
func startGame(handler *handlers.GameHandler) (*game.Data, error) {
	gameData, err := (*handler).GetGame()
	if err != nil {
		return nil, err
	}

	return gameData, nil
}

//getWeather receives weather information from 3rd party
func setWeather(handler *handlers.GameHandler, gameData *game.Data) error {
	weatherData, err := (*handler).GetWeather(gameData.GameID)
	if err != nil {
		return err
	}

	weatherData.AddType()
	gameData.Weather = *weatherData

	return nil
}

//addDragon generates dragon by using knight data & assigns it to game data
func addDragon(gameData *game.Data) {
	gameData.Dragon = *dragon.CreateDragon(gameData.Knight.Attack,
		gameData.Knight.Armor,
		gameData.Knight.Agility,
		gameData.Knight.Endurance,
		gameData.Weather.Type)
}
