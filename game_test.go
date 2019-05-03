package battleship_go

import "testing"

func TestPlaceShip(t *testing.T) {
	for _, tc := range testCases {
		expected := tc.ship
		game := newGame()
		game.placePiece(tc.grid, tc.ship, tc.horizontal)
		for _, expectedGrids := range tc.grids {
			if observed := game.board.get(expectedGrids); observed != expected {
				t.Fatalf("%v : %v, want %v", tc.description, observed, expected)
			}
		}
	}
}

func TestCanPlaceShipWithDifferentStartingPosition(t *testing.T) {
	for _, tc := range testCasesShouldNotReplaceShipAlreadyInPosition {
		game := newGame()
		game.placePiece(tc.pos1, tc.ship1, tc.horizontal)
		game.placePiece(tc.pos2, tc.ship2, tc.horizontal)

		expected := tc.ship1
		if observed := game.board.get(tc.pos1); observed != expected {
			t.Fatalf("%v : %v, want %v", tc.description, observed, expected)
		}

		if observed := game.board.get(tc.otherPos); observed != nil {
			t.Fatalf("ship2 should not be placed at all : was %v", observed)
		}
	}
}

func BenchmarkPlaceShip(b *testing.B) {
	//for i := 0; i < b.N; i++ {
	//PlaceShip()
	//}
}
