package battleship_go

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
			ship := board.board[board.convertToGridId(i, j)]
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
