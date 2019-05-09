package engine

type ShipType struct {
	name   string
	length int
}

func newShipType(name string, length int) *ShipType {
	shipType := new(ShipType)
	shipType.name = name
	shipType.length = length
	return shipType
}

type Ship struct {
	size int
}

func newShip(size int) *Ship {
	ship := new(Ship)
	ship.size = size
	return ship
}
