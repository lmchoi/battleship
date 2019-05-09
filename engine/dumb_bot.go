package engine

import (
	"math/rand"
	"time"
)

type DumbBot struct {
	board *Board
	ships []Ship
}

func (bot *DumbBot) placeShips() {
	pos := newPosition(0, 0)
	for i := 0; i < len(bot.ships); {
		pos.x = rand.Intn(8)
		pos.y = rand.Intn(8)
		var horizontal = rand.Int()%2 == 0
		if bot.board.placePiece(pos, &bot.ships[i], horizontal) {
			i++
		}
	}
}

func newDumbBot(game *Game) *DumbBot {
	bot := new(DumbBot)
	bot.board = newBoard(game.boardSize)
	bot.ships = game.ships
	rand.Seed(time.Now().UnixNano())
	return bot
}
