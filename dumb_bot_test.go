package battleship_go

import "testing"

func TestDumbBot(t *testing.T) {
	game := newGame()
	bot := newDumbBot(game)
	println("Before: ")
	bot.board.print()
	bot.placeShips()
	println("After: ")
	bot.board.print()
	if bot.board.get(newPosition(0, 0)) == nil {
		t.Fatal("there should be a ship at the corner")
	}
}
