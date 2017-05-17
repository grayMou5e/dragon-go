package handlers

import (
	"net/http"

	"io/ioutil"

	"encoding/json"
	"encoding/xml"

	"fmt"

	"bytes"

	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/result"
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
func (GameHandler *GameHTTPHandler) GetWeather(gameID int) (weatherData *weather.Data) {
	bodyBytes := getAPIBodyBytes(fmt.Sprintf("%sweather/api/report/%d", baseRestAPIURL, gameID), GameHandler)
	xml.Unmarshal(*bodyBytes, &weatherData)

	return weatherData
}

//FightAgainstTheKnight Method for starting fight against the knight
func (GameHandler *GameHTTPHandler) FightAgainstTheKnight(dragonData *dragon.Data, gameID int) (result *result.Data) {
	var b []byte
	if dragonData.Scared {
		b = []byte("{\"dragon\":null}")
	} else {
		var err error
		b, err = json.Marshal(struct {
			Dragon dragon.Data `json:"dragon"`
		}{*dragonData})
		if err != nil {
			panic(err)
		}
	}

	req, err2 := http.NewRequest("PUT", fmt.Sprintf("%sapi/game/%d/solution", baseRestAPIURL, gameID), bytes.NewBuffer(b))
	if err2 != nil {
		panic(err2)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err3 := GameHandler.httpClient.Do(req)
	req.Close = true
	if err3 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		//do something
	}

	bodyBytes, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		panic(errr)
	}
	json.Unmarshal(bodyBytes, &result)

	return result
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
