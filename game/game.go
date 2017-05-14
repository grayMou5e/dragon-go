package game

import (
	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/knight"
	"github.com/grayMou5e/dragon-go/weather"
)

//Data struct is dedicated for holding up the information about current game
type Data struct {
	GameID  int           `json:"gameId"`
	Knight  knight.Knight `json:"knight"`
	Dragon  dragon.Dragon
	Weather weather.Weather
}
