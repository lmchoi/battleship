package battleship_go

const defaultBoardSize = 10 // 10 x 10

type Game struct {
	board *Board
}

func (game *Game) placePiece(grid *Position, ship *Ship, horizontal bool) {
	if game.canPlace(grid, ship, horizontal) {
		tmpGrid := newPosition(grid.x, grid.y)
		if horizontal {
			for i := 0; i < ship.size; i++ {
				tmpGrid.x += i
				game.board.set(tmpGrid, ship)
			}
		} else {
			for i := 0; i < ship.size; i++ {
				tmpGrid.y += i
				game.board.set(tmpGrid, ship)
			}
		}
	}
}

func (game *Game) canPlace(startingPos *Position, ship *Ship, horizontal bool) bool {
	pos := newPosition(startingPos.x, startingPos.y)
	if horizontal {
		for i := 0; i < ship.size; i++ {
			pos.x += i
			if game.board.get(pos) != nil {
				return false
			}
		}
	} else {
		for i := 0; i < ship.size; i++ {
			pos.y += i
			if game.board.get(pos) != nil {
				return false
			}
		}
	}

	return true
}

func newGame() *Game {
	game := new(Game)
	game.board = newBoard(defaultBoardSize)
	return game
}
