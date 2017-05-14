package main

import (
	"fmt"

	"github.com/grayMou5e/dragon-go/game"
	"github.com/grayMou5e/dragon-go/handlers"
)

func main() {
	playGame()
}

func playGame() {
	game := startGame()
	fmt.Println(game.Knight)

}

func startGame() *game.Data {
	handler := handlers.NewHandler()

	var t handlers.GameHandler
	t = handler

	gameData := t.GetGame()

	t.GetWeather(gameData.GameID)

	return gameData
}
