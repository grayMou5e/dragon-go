package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	"encoding/json"

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

	bodyBytes := getAPIBodyBytes(server.URL, httpHandler)

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

	game := handler.GetGame()

	assert.Equal(expectedGame.GameID, game.GameID)
	assert.Equal(expectedGame.Knight, game.Knight)
}
