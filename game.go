package battleship_go

const defaultBoardSize = 10 // 10 x 10

type Game struct {
	bot       *DumbBot
	boardSize int
}

func newGame() *Game {
	game := new(Game)
	game.boardSize = defaultBoardSize
	return game
}

func (game *Game) addBot() {
	game.bot = newDumbBot(game)
}

func (game *Game) start() {
	game.bot.placeShips()
}
