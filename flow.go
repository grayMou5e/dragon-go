package main

import (
	"time"

	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
	"github.com/grayMou5e/dragon-go/utilities"
	uuid "github.com/nu7hatch/gouuid"
	"go.uber.org/zap"
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

	fightErr := fightAgainstTheKnight(handler, game)
	if fightErr != nil {
		return nil, fightErr
	}

	return game, nil
}

func timedPlayGame(handler *handlers.GameHandler, logger *zap.Logger, correlationID *uuid.UUID) (*game.Data, error) {
	game, gameErr := timedStartGame(handler, logger, correlationID)
	if gameErr != nil {
		return nil, gameErr
	}

	timedSetWeather(handler, game, logger, correlationID)
	timedAddDragon(game, logger, correlationID)

	fightErr := timedFightAgainstTheKnight(handler, game, logger, correlationID)
	if fightErr != nil {
		return nil, fightErr
	}

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

//timedStartGame receives game from 3rd party & logs how long did the call took
func timedStartGame(handler *handlers.GameHandler, logger *zap.Logger, correlationID *uuid.UUID) (*game.Data, error) {
	startTime := time.Now()
	gameData, err := startGame(handler)
	elapsedTime := time.Since(startTime)
	if err != nil {
		utilities.LogError(err, "Get game", elapsedTime, correlationID, logger)
		return nil, err
	}

	utilities.LogInfo("Get game", elapsedTime, correlationID, logger)
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

//getWeather receives weather information from 3rd party & logs how long did the call took
func timedSetWeather(handler *handlers.GameHandler, gameData *game.Data, logger *zap.Logger, correlationID *uuid.UUID) error {
	startTime := time.Now()
	err := setWeather(handler, gameData)
	elapsedTime := time.Since(startTime)
	if err != nil {
		if err != nil {
			utilities.LogError(err, "Get weather", elapsedTime, correlationID, logger)
			return err
		}
	}

	utilities.LogInfo("Get weather", elapsedTime, correlationID, logger)
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

//fightAgainstTheKnight fights agains the knight using 3rd party provider
func fightAgainstTheKnight(handler *handlers.GameHandler, game *game.Data) error {
	gameResult, fightError := (*handler).FightAgainstTheKnight(&game.Dragon, game.GameID)
	if fightError != nil {
		return fightError
	}
	game.Result = *gameResult
	game.Result.Summarize()

	return nil
}

//fightAgainstTheKnight fights agains the knight using 3rd party provider & logs how long did it took
func timedFightAgainstTheKnight(handler *handlers.GameHandler, game *game.Data, logger *zap.Logger, correlationID *uuid.UUID) error {
	startTime := time.Now()
	fightError := fightAgainstTheKnight(handler, game)
	elapsedTime := time.Since(startTime)
	if fightError != nil {
		utilities.LogError(fightError, "Fight", elapsedTime, correlationID, logger)
		return fightError
	}

	utilities.LogInfo("Fight", elapsedTime, correlationID, logger)
	return nil
}

//timedAddDragon creates dragon by current knight & logs elapsed time
func timedAddDragon(gameData *game.Data, logger *zap.Logger, correlationID *uuid.UUID) {
	startTime := time.Now()
	addDragon(gameData)
	elapsedTime := time.Since(startTime)

	utilities.LogInfo("Create dragon", elapsedTime, correlationID, logger)
}
