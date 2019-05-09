package engine

const defaultBoardSize = 10 // 10 x 10

type Game struct {
	bot        *DumbBot
	boardSize  int
	shipCounts map[ShipType]int
	ships      []Ship
}

func NewGame() *Game {
	game := new(Game)
	game.boardSize = defaultBoardSize
	game.shipCounts = map[ShipType]int{
		*newShipType("Carrier", 5):    1,
		*newShipType("Battleship", 4): 2,
		*newShipType("Cruiser", 3):    3,
		*newShipType("Submarine", 3):  4,
		*newShipType("Destroyer", 2):  5,
	}

	// total ship count is 15 for the default set
	game.ships = make([]Ship, 15)
	var id = 0
	for ship, count := range game.shipCounts {
		for i := 0; i < count; i++ {
			game.ships[id] = *newShip(ship.length)
			id++
		}
	}

	// TODO  part of setup?
	game.addBot()
	return game
}

func (game *Game) addBot() {
	game.bot = newDumbBot(game)
}

func (game *Game) Start() {
	game.bot.placeShips()

	game.ShowPlayersBoard()

	// turn start -> prompt player to guess
	// update status
	// show results -> end game?
	// next round
	//

	// show winner

	//game.addBot()
}

func (game *Game) ShowPlayersBoard() {
	// TODO for now show bot's board
	game.bot.board.print()
}
