package engine

type Position struct {
	x int
	y int
}

func newPosition(x int, y int) *Position {
	pos := new(Position)
	pos.x = x
	pos.y = y
	return pos
}
