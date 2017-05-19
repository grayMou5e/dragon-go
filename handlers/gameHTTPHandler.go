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

//NewHandler creates new Game Handler
func NewHandler() *GameHTTPHandler {
	httpClient := http.Client{}

	return &GameHTTPHandler{
		httpClient:     httpClient,
		baseRestAPIURL: "http://www.dragonsofmugloar.com",
	}
}

//GameHTTPHandler handles game http requests
type GameHTTPHandler struct {
	httpClient     http.Client
	baseRestAPIURL string
}

//GetGame method for receiving game data
func (h *GameHTTPHandler) GetGame() (gameData *game.Data) {
	bodyBytes := getAPIBodyBytes(h.baseRestAPIURL+"/api/game", h)

	json.Unmarshal(bodyBytes, &gameData)

	return gameData
}

//GetWeather receives weather from api
func (h *GameHTTPHandler) GetWeather(gameID int) (weatherData *weather.Data) {
	bodyBytes := getAPIBodyBytes(fmt.Sprintf("%s/weather/api/report/%d", h.baseRestAPIURL, gameID), h)

	xml.Unmarshal(bodyBytes, &weatherData)

	return weatherData
}

//FightAgainstTheKnight Method for starting fight against the knight
func (h *GameHTTPHandler) FightAgainstTheKnight(dragonData *dragon.Data, gameID int) (resultData *result.Data) {
	b, preparationErr := prepareDragonForSending(*dragonData)
	if preparationErr != nil {
		panic(preparationErr)
	}

	req, err2 := http.NewRequest("PUT", fmt.Sprintf("%s/api/game/%d/solution", h.baseRestAPIURL, gameID), bytes.NewBuffer(b))
	if err2 != nil {
		panic(err2)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err3 := h.httpClient.Do(req)
	req.Close = true
	if err3 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Server is giving - %d", resp.StatusCode))
	}

	bodyBytes, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		panic(errr)
	}
	json.Unmarshal(bodyBytes, &resultData)

	return resultData
}

func prepareDragonForSending(dragonData dragon.Data) (b []byte, err error) {
	if dragonData.Scared {
		return []byte("{\"dragon\":null}"), nil
	}

	b, err = json.Marshal(struct {
		Dragon dragon.Data `json:"dragon"`
	}{dragonData})

	if err != nil {
		return nil, err
	}

	return b, err
}

func getAPIBodyBytes(url string, gameHandler *GameHTTPHandler) []byte {
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

	return bodyBytes
}
