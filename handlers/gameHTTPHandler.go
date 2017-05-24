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
func (h *GameHTTPHandler) GetGame() (gameData *game.Data, err error) {
	bodyBytes, err := getAPIBodyBytes(h.baseRestAPIURL+"/api/game", h)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(bodyBytes, &gameData)

	return gameData, nil
}

//GetWeather receives weather from api
func (h *GameHTTPHandler) GetWeather(gameID int) (weatherData *weather.Data, err error) {
	bodyBytes, err := getAPIBodyBytes(fmt.Sprintf("%s/weather/api/report/%d", h.baseRestAPIURL, gameID), h)
	if err != nil {
		return nil, err
	}

	xml.Unmarshal(bodyBytes, &weatherData)

	return weatherData, nil
}

//FightAgainstTheKnight Method for starting fight against the knight
func (h *GameHTTPHandler) FightAgainstTheKnight(dragonData *dragon.Data, gameID int) (resultData *result.Data, err error) {
	b, preparationErr := prepareDragonForSending(*dragonData)
	if preparationErr != nil {
		return nil, preparationErr
	}

	req, err2 := http.NewRequest("PUT", fmt.Sprintf("%s/api/game/%d/solution", h.baseRestAPIURL, gameID), bytes.NewBuffer(b))
	if err2 != nil {
		return nil, err2
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err3 := h.httpClient.Do(req)
	req.Close = true
	if err3 != nil {
		return nil, err2
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server is giving - %d", resp.StatusCode)
	}

	bodyBytes, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		return nil, (errr)
	}

	unmarshalErr := json.Unmarshal(bodyBytes, &resultData)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return resultData, nil
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

func getAPIBodyBytes(url string, gameHandler *GameHTTPHandler) ([]byte, error) {
	resp, err := gameHandler.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Returned %d instead of 200", resp.StatusCode)
	}

	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	return bodyBytes, nil
}
