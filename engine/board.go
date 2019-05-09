package engine

type Board struct {
	size  int
	board map[int]*Ship // grid id, ship id
}

func newBoard(size int) *Board {
	board := new(Board)
	board.size = size
	board.board = make(map[int]*Ship)
	return board
}

func (board *Board) placePiece(startingPos *Position, ship *Ship, horizontal bool) bool {
	if horizontal {
		if (startingPos.x + ship.size) > board.size {
			return false
		}
	} else {
		if (startingPos.y + ship.size) > board.size {
			return false
		}
	}

	if board.canPlace(startingPos, ship, horizontal) {
		pos := newPosition(startingPos.x, startingPos.y)
		if horizontal {
			for i := 0; i < ship.size; i++ {
				board.set(pos, ship)
				pos.x++
			}
		} else {
			for i := 0; i < ship.size; i++ {
				board.set(pos, ship)
				pos.y++
			}
		}
		return true
	}

	return false
}

func (board *Board) canPlace(startingPos *Position, ship *Ship, horizontal bool) bool {
	pos := newPosition(startingPos.x, startingPos.y)
	if horizontal {
		for i := 0; i < ship.size; i++ {
			if board.get(pos) != nil {
				return false
			}
			pos.x++
		}
	} else {
		for i := 0; i < ship.size; i++ {
			if board.get(pos) != nil {
				return false
			}
			pos.y++
		}
	}

	return true
}

func (board *Board) set(grid *Position, ship *Ship) bool {
	if board.get(grid) == nil {
		board.board[board.convertToGridId(grid.x, grid.y)] = ship
		return true
	}
	return false
}

func (board *Board) get(grid *Position) *Ship {
	return board.board[board.convertToGridId(grid.x, grid.y)]
}

func (board *Board) print() {
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			ship := board.board[board.convertToGridId(j, i)]
			if ship != nil {
				print(ship.size)
			} else {
				print(".")
			}
		}
		println()
	}
}

func (board *Board) convertToGridId(x int, y int) int {
	return x + y*board.size
}
