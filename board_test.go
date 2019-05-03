package battleship_go

import "testing"

func TestBoard(t *testing.T) {
	var testCases = []struct {
		description    string
		grid           *Position
		ship           *Ship
		expectedGridId int
	}{
		{
			"set grid to be occupied by a ship at (0, 0)",
			newPosition(0, 0),
			newShip(1),
			0,
		},
		{
			"set grid to be occupied by a ship at (1. 0)",
			newPosition(1, 0),
			newShip(1),
			1,
		},
		{
			"set grid to be occupied by a ship at (0, 1)",
			newPosition(0, 1),
			newShip(1),
			defaultBoardSize,
		},
		{
			"set grid to be occupied by a ship at (1, 1)",
			newPosition(1, 1),
			newShip(1),
			11,
		},
	}

	for _, tc := range testCases {
		board := newBoard(defaultBoardSize)
		ship := tc.ship
		expected := ship
		board.set(tc.grid, ship)
		if observed := board.get(tc.grid); observed != expected {
			t.Fatalf("%v : %v, want %v", tc.description, observed, expected)
		}
	}
}

func TestReturnFromSet(t *testing.T) {
	description := "should fail to set ship2 and return ship1 on the grid and return false"
	board := newBoard(defaultBoardSize)
	grid := newPosition(0, 0)
	ship1 := newShip(2)
	ship2 := newShip(3)
	result1 := board.set(grid, ship1)
	result2 := board.set(grid, ship2)
	expected := ship1
	if observed := board.get(grid); observed != expected {
		t.Fatalf("%v : %v, want %v", description, observed, expected)
	}

	if result1 != true {
		t.Fatalf("expected successful operation to return true, but got %v", result1)

	}

	if result2 != false {
		t.Fatalf("expected successful operation to return false, but got %v", result1)
	}
}

func BenchmarkBoard(b *testing.B) {
	//for i := 0; i < b.N; i++ {
	//PlaceShip()
	//}
}
