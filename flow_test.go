package main

import (
	"testing"

	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
	"github.com/grayMou5e/dragon-go/knight"
	"github.com/grayMou5e/dragon-go/result"
	"github.com/grayMou5e/dragon-go/weather"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
)

type Handler struct {
}

func (h *Handler) GetGame() (gameData *game.Data, err error) {
	gameData = &game.Data{GameID: 1, Knight: knight.Data{Agility: 8, Armor: 6, Attack: 5, Endurance: 1}}
	return gameData, nil
}
func (h *Handler) GetWeather(gameID int) (weatherData *weather.Data, err error) {
	return &weather.Data{Code: "nmr", Message: "The Weather Girls - It's not Raining Men"}, nil
}
func (h *Handler) FightAgainstTheKnight(dragonData *dragon.Data, gameID int) (resultData *result.Data, err error) {
	resultData = &result.Data{Status: "Victory", Message: "yeah"}
	return resultData, nil
}

func Test_startGame(t *testing.T) {
	assert := assert.New(t)
	handlerMock := Handler{}
	var handler handlers.GameHandler
	handler = &handlerMock

	gameData, _ := startGame(&handler)

	assert.Equal(1, gameData.GameID)
}

func Test_GetWeather(t *testing.T) {
	assert := assert.New(t)
	handlerMock := Handler{}
	var handler handlers.GameHandler
	handler = &handlerMock

	expectedWeather, _ := handler.GetWeather(1)
	gameData := game.Data{GameID: 1}

	setWeather(&handler, &gameData)
	assert.Equal(expectedWeather.Code, gameData.Weather.Code)
	assert.Equal(expectedWeather.Message, gameData.Weather.Message)
	assert.Equal(weather.NormalWeather, gameData.Weather.Type)
}

func Test_AddDragon_WithoutWeatherData(t *testing.T) {
	assert := assert.New(t)
	gameData := game.Data{Knight: knight.Data{Agility: 8, Armor: 6, Attack: 5, Endurance: 1}}

	addDragon(&gameData)

	assert.Equal(int8(10), gameData.Dragon.WingStrength)
	assert.Equal(int8(5), gameData.Dragon.ClawSharpness)
	assert.Equal(int8(4), gameData.Dragon.ScaleThickness)
	assert.Equal(int8(1), gameData.Dragon.FireBreath)
	assert.Equal(false, gameData.Dragon.Scared)
}

func Test_AddDragon_Scared(t *testing.T) {
	assert := assert.New(t)
	gameData := game.Data{Knight: knight.Data{Agility: 8, Armor: 6, Attack: 5, Endurance: 1},
		Weather: weather.Data{Type: weather.StormWeather}}

	addDragon(&gameData)

	assert.Equal(int8(0), gameData.Dragon.WingStrength)
	assert.Equal(int8(0), gameData.Dragon.ScaleThickness)
	assert.Equal(int8(0), gameData.Dragon.ClawSharpness)
	assert.Equal(int8(0), gameData.Dragon.FireBreath)
	assert.Equal(true, gameData.Dragon.Scared)
}

func Test_FightAgainstTheKnight(t *testing.T) {
	assert := assert.New(t)
	handlerMock := Handler{}
	var handler handlers.GameHandler
	handler = &handlerMock
	guid, _ := uuid.NewV4()

	gameData, _ := playGame(&handler, guid)

	assert.Equal(1, gameData.GameID)

	assert.Equal(int8(8), gameData.Knight.Agility)
	assert.Equal(int8(6), gameData.Knight.Armor)
	assert.Equal(int8(5), gameData.Knight.Attack)
	assert.Equal(int8(1), gameData.Knight.Endurance)

	assert.Equal(int8(10), gameData.Dragon.WingStrength)
	assert.Equal(int8(5), gameData.Dragon.ClawSharpness)
	assert.Equal(int8(4), gameData.Dragon.ScaleThickness)
	assert.Equal(int8(1), gameData.Dragon.FireBreath)
	assert.Equal(false, gameData.Dragon.Scared)

	assert.Equal("nmr", gameData.Weather.Code)
	assert.Equal("The Weather Girls - It's not Raining Men", gameData.Weather.Message)
	assert.Equal(weather.NormalWeather, gameData.Weather.Type)

	assert.Equal("yeah", gameData.Result.Message)
	assert.Equal("Victory", gameData.Result.Status)
	assert.Equal(true, gameData.Result.Victory)
}
