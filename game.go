package battleship_go

const defaultBoardSize = 10 // 10 x 10

type Game struct {
	bot        *DumbBot
	boardSize  int
	shipCounts map[ShipType]int
	ships      []Ship
}

func newGame() *Game {
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

	return game
}

func (game *Game) addBot() {
	game.bot = newDumbBot(game)
}

func (game *Game) start() {
	game.bot.placeShips()
}
