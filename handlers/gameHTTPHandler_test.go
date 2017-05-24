package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	"encoding/json"

	"github.com/grayMou5e/dragon-go/dragon"
	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/knight"
	"github.com/stretchr/testify/assert"
)

func Test_getAPIBodyBytes(t *testing.T) {
	assert := assert.New(t)
	expectedResponse := "{Hello}"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedResponse))
	}))
	defer server.Close()

	httpHandler := NewHandler()

	bodyBytes, err := getAPIBodyBytes(server.URL, httpHandler)

	assert.Nil(err)
	assert.NotNil(bodyBytes, "Returned response body bytes shouldnt be nil")

	assert.Equal(expectedResponse, string(bodyBytes), fmt.Sprintf("Expected [%s] but received [%s]", expectedResponse, bodyBytes))
}

func Test_GetGame(t *testing.T) {
	assert := assert.New(t)
	expectedGame := game.Data{GameID: 10, Knight: knight.Data{Attack: 1, Agility: 1, Endurance: 1, Armor: 1}}
	response, _ := json.Marshal(expectedGame)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/game" {
			w.Write(response)
		}
	}))
	defer server.Close()

	handler := NewHandler()
	handler.baseRestAPIURL = server.URL

	game, err := handler.GetGame()

	assert.Nil(err)
	assert.Equal(expectedGame.GameID, game.GameID)
	assert.Equal(expectedGame.Knight, game.Knight)
}

func Test_GetWeather(t *testing.T) {
	assert := assert.New(t)
	gameID := 1
	code := "NMR"
	message := "Another day of everyday normal regular weather, business as usual, unless it’s going to be like the time of the Great Paprika Mayonnaise Incident of 2014, that was some pretty nasty stuff."
	xml := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?><report><code>%s</code><message>%s</message></report>", code, message)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == fmt.Sprintf("/weather/api/report/%d", gameID) {
			w.Write([]byte(xml))
		}
	}))
	defer server.Close()

	handler := NewHandler()
	handler.baseRestAPIURL = server.URL

	weather, err := handler.GetWeather(gameID)

	assert.Nil(err)
	assert.Equal(message, weather.Message)
}

func Test_PrepareDragonForSending(t *testing.T) {
	assert := assert.New(t)
	dragonData := []struct {
		data           dragon.Data
		expectedResult []byte
	}{
		{dragon.Data{Scared: true}, []byte("{\"dragon\":null}")},
		{dragon.Data{ClawSharpness: 1, FireBreath: 1, ScaleThickness: 1, WingStrength: 1},
			[]byte("{\"dragon\":{\"scaleThickness\":1,\"clawSharpness\":1,\"wingStrength\":1,\"fireBreath\":1}}")}}

	for _, data := range dragonData {
		body, err := prepareDragonForSending(data.data)

		assert.Nil(err)
		assert.Equal(string(data.expectedResult), string(body))
	}
}

func Test_FightAgainstTheKnight(t *testing.T) {
	assert := assert.New(t)
	dragonData := dragon.Data{ClawSharpness: 1}
	gameID := 1
	status := "Defeat"
	message := "Sorry mate you lost"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == fmt.Sprintf("/api/game/%d/solution", gameID) && r.Method == "PUT" {
			w.Write([]byte(fmt.Sprintf("{\"status\":\"%s\",\"message\":\"%s\"}", status, message)))
		}
	}))
	defer server.Close()

	handler := NewHandler()
	handler.baseRestAPIURL = server.URL

	rez, err := handler.FightAgainstTheKnight(&dragonData, gameID)

	assert.Nil(err)
	assert.Equal(status, rez.Status)
	assert.Equal(message, rez.Message)
}

func Test_FightAgainstTheKnight_Return500(t *testing.T) {
	assert := assert.New(t)
	dragonData := dragon.Data{ClawSharpness: 1}
	gameID := 1
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == fmt.Sprintf("/api/game/%d/solution", gameID) && r.Method == "PUT" {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer server.Close()

	handler := NewHandler()
	handler.baseRestAPIURL = server.URL

	_, err := handler.FightAgainstTheKnight(&dragonData, gameID)

	assert.Error(err)
}
