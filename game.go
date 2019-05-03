package battleship_go

const defaultBoardSize = 10 // 10 x 10

type Game struct {
	bot       *DumbBot
	boardSize int
	ships     []Ship
}

func newGame() *Game {
	game := new(Game)
	game.boardSize = defaultBoardSize
	game.ships = make([]Ship, 10)
	for i := 0; i < len(game.ships); i++ {
		game.ships[i] = *newShip(2)
	}
	return game
}

func (game *Game) addBot() {
	game.bot = newDumbBot(game)
}

func (game *Game) start() {
	game.bot.placeShips()
}
