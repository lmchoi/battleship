package battleship_go

var testCases = []struct {
	description string
	grid        *Position
	ship        *Ship
	horizontal  bool
	grids       []*Position
}{
	{
		"place ship of size 1 horizontally",
		newPosition(0, 0),
		newShip(1),
		true,
		[]*Position{newPosition(0, 0)},
	},
	{
		"place ship of size 2 horizontally",
		newPosition(0, 0),
		newShip(2),
		true,
		[]*Position{newPosition(0, 0), newPosition(1, 0)},
	},
	{
		"place ship of size 2 vertically",
		newPosition(0, 0),
		newShip(2),
		false,
		[]*Position{newPosition(0, 0), newPosition(0, 1)},
	},
}

var testCasesShouldNotReplaceShipAlreadyInPosition = []struct {
	description string
	pos1        *Position
	pos2        *Position
	ship1       *Ship
	ship2       *Ship
	horizontal  bool
	otherPos    *Position // to be nil
}{
	{
		"both ships start at the same position, horizontal",
		newPosition(0, 0),
		newPosition(0, 0),
		newShip(1),
		newShip(2),
		true,
		newPosition(1, 0),
	},
	{
		"new ship at tail, horizontal",
		newPosition(1, 0),
		newPosition(0, 0),
		newShip(1),
		newShip(2),
		true,
		newPosition(0, 0),
	},
	{
		"both ships start at the same position, vertical",
		newPosition(0, 0),
		newPosition(0, 0),
		newShip(1),
		newShip(2),
		false,
		newPosition(0, 1),
	},
	{
		"new ship at tail, vertical",
		newPosition(0, 1),
		newPosition(0, 0),
		newShip(1),
		newShip(2),
		false,
		newPosition(0, 0),
	},
}
