package handlers

import (
	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/result"
	"github.com/grayMou5e/dragon-go/weather"
)

//GameHandler interface for game handlers
type GameHandler interface {
	GetGame() (gameData *game.Data)
	GetWeather(gameID int) (weatherData *weather.Data)
	FightAgainstTheKnight(dragonData *dragon.Data, gameID int) (result *result.Data)
}
