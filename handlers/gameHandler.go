package handlers

import (
	"net/http"

	"io/ioutil"

	"encoding/json"
	"encoding/xml"

	"fmt"

	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/weather"
)

const baseRestAPIURL = "http://www.dragonsofmugloar.com/"

//NewHandler creates new Game Handler
func NewHandler() *GameHTTPHandler {
	httpClient := http.Client{}

	return &GameHTTPHandler{
		httpClient: httpClient,
	}
}

//GameHTTPHandler handles game http requests
type GameHTTPHandler struct {
	httpClient http.Client
}

//GetGame method for receiving game data
func (GameHandler *GameHTTPHandler) GetGame() (gameData *game.Data) {
	bodyBytes := getAPIBodyBytes(baseRestAPIURL+"api/game", GameHandler)

	json.Unmarshal(*bodyBytes, &gameData)

	return gameData
}

//GetWeather receives weather from api
func (GameHandler *GameHTTPHandler) GetWeather(gameID int) (weatherData *weather.Weather) {
	bodyBytes := getAPIBodyBytes(fmt.Sprintf("%sweather/api/report/%d", baseRestAPIURL, gameID), GameHandler)
	xml.Unmarshal(*bodyBytes, &weatherData)

	return weatherData
}

func getAPIBodyBytes(url string, gameHandler *GameHTTPHandler) *[]byte {
	resp, err := gameHandler.httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		//isue
	}

	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err2)
	}

	return &bodyBytes
}
