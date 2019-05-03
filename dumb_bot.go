package battleship_go

type DumbBot struct {
	board *Board
}

func (bot *DumbBot) placeShips() {
	pos := newPosition(0, 0)
	for i := 0; i < 1; i++ {
		pos.x += i
		bot.board.placePiece(pos, newShip(2), false)
	}
}

func newDumbBot(game *Game) *DumbBot {
	bot := new(DumbBot)
	bot.board = newBoard(game.boardSize)
	return bot
}
