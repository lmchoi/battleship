package battleship_go

type Ship struct {
	size int
}

func newShip(size int) *Ship {
	ship := new(Ship)
	ship.size = size
	return ship
}
