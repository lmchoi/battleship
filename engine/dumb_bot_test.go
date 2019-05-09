package engine

import (
	"testing"
)

func TestDumbBot(t *testing.T) {
	game := NewGame()
	bot := newDumbBot(game)
	println("Before: ")
	bot.board.print()
	bot.placeShips()
	println("After: ")
	bot.board.print()
	//TODO
	//if bot.board.get(newPosition(0, 0)) == nil {
	//	t.Fatal("there should be a ship at the corner")
	//}
}
