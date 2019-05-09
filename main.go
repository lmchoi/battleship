package main

import (
	"doink/battleship/engine"
)

func main() {
	// select number of players / bots
	// select rules (should have default rules) -> create rules

	// create game using rules + config -> for now create in the constructor; remember to extract
	game := engine.NewGame()

	// game start -> prompt player to place ships -> player interface? human vs bot
	game.Start()

}
